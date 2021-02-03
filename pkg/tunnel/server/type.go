package server

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/armon/go-metrics"
	"github.com/inconshreveable/go-vhost"
	"github.com/yelinaung/go-haikunator"
	"go.uber.org/zap"
)

type ClientInfo struct {
	Hostname string
	Address  string
}

type PublicAddress struct {
	Address string
	net.Listener

	free func()
	once sync.Once
}

func (c *PublicAddress) Close() error {
	var err error

	c.once.Do(func() {
		err = c.Listener.Close()
		c.free()
	})

	return err
}

func NewAddresses(logger *zap.Logger, hostname string, muxer *vhost.VhostMuxer) *Addresses {
	metrics.SetGauge([]string{"addresses"}, 0)

	return &Addresses{
		logger:     logger,
		hostname:   hostname,
		httpMux:    muxer,
		addresses:  make(map[string]struct{}),
		haikunator: haikunator.New(time.Now().UnixNano()),
	}
}

type Addresses struct {
	hostname string
	httpMux  *vhost.VhostMuxer
	logger   *zap.Logger

	addresses map[string]struct{}
	lock      sync.RWMutex

	haikunator haikunator.Name
}

func (c *Addresses) free(address string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.addresses, address)

	metrics.SetGauge([]string{"addresses"}, float32(len(c.addresses)))
	c.logger.Debug("public address freed", zap.String("address", address))
}

func (c *Addresses) put(address string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.addresses[address]; ok {
		return errors.New("address already exists")
	}

	c.addresses[address] = struct{}{}

	metrics.SetGauge([]string{"addresses"}, float32(len(c.addresses)))
	c.logger.Debug("public address created", zap.String("address", address))
	return nil
}

func (c *Addresses) ClaimAddress(addressType string, address string) (*PublicAddress, error) {
	switch addressType {
	case "tcp":
		_, port, err := net.SplitHostPort(address)
		if err != nil {
			return nil, err
		}

		listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			return nil, err
		}

		if err := c.put(address); err != nil {
			listener.Close()
			return nil, err
		}

		return &PublicAddress{
			Address:  address,
			Listener: listener,
			free: func() {
				c.free(address)
			},
		}, nil
	case "tls":
		vhost, err := c.httpMux.Listen(address)
		if err != nil {
			return nil, err
		}

		if err := c.put(address); err != nil {
			vhost.Close()
			return nil, err
		}

		return &PublicAddress{
			Address:  address,
			Listener: vhost,
			free: func() {
				c.free(address)
			},
		}, nil
	default:
		url, err := url.Parse(address)
		if err != nil {
			return nil, err
		}

		hostname := url.Host
		address := fmt.Sprintf("https://%s", hostname)

		vhost, err := c.httpMux.Listen(hostname)
		if err != nil {
			return nil, err
		}
		c.logger.Debug("mux created for http address claim", zap.String("hostname", hostname), zap.String("address", address))

		if err := c.put(address); err != nil {
			vhost.Close()
			return nil, err
		}

		return &PublicAddress{
			Address:  address,
			Listener: vhost,
			free: func() {
				c.free(address)
			},
		}, nil
	}

}

func (c *Addresses) NewAddress(addressType string) (*PublicAddress, error) {
	switch addressType {
	case "tcp":
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			return nil, err
		}

		_, port, err := net.SplitHostPort(listener.Addr().String())
		if err != nil {
			listener.Close()
			return nil, err
		}

		address := fmt.Sprintf("tcp.%s:%s", c.hostname, port)

		if err := c.put(address); err != nil {
			listener.Close()
			return nil, err
		}

		return &PublicAddress{
			Address:  address,
			Listener: listener,
			free: func() {
				c.free(address)
			},
		}, nil
	default:
		id := c.haikunator.Haikunate()

		hostname := fmt.Sprintf("%s.%s", id, c.hostname)
		address := fmt.Sprintf("https://%s", hostname)

		vhost, err := c.httpMux.Listen(hostname)
		if err != nil {
			return nil, err
		}

		if err := c.put(address); err != nil {
			vhost.Close()
			return nil, err
		}
		c.logger.Debug("mux created for new http address", zap.String("hostname", hostname), zap.String("address", address))

		return &PublicAddress{
			Address:  address,
			Listener: vhost,
			free: func() {
				c.free(address)
			},
		}, nil
	}
}

type TunnelEndpoint interface {
	Id() string
	Serve() error
}
