# Tunl

With Tunl you can expose services on your localhost to the public via a fast and secure tunnel.

## Expose local http process

```
$ tunl http localhost:3000
https://red-fox.tunl.es/ -> http://localhost:3000
```

## Share a local directory

```
$ tunl files /the/directory
https://red-fox.tunl.es/ -> /the/directory
```

## Expose a local MySQL

```
$ tunl tcp localhost:3306
tunl.es:89311 -> localhost:3306
```

## Expose a service running in a docker container

```
$ tunl docker <container-name> <port>
tunl.es:48221 -> <container-name>:<port>
```

# Installation

Download the [binary release](https://github.com/pjvds/tunl/releases/latest) and put it somewhere in your `$PATH`.

## Basic 

**curl**
```
# install it to ./bin/tunl
curl -fsSL https://get.tunl.es | sh
```

```
# install it to /sbin/tunl
curl -fsSL https://get.tunl.es | sudo sh -s -- -b /sbin
```

**wget**

```
# install it to ./bin/tunl
wget -qO- https://get.tunl.es | sh
```

```
# install it to /sbin/tunl
wget -qO- https://get.tunl.es | sudo sh -s -- -b /sbin
```

## Arch Linux

```
yay tunl
```

## Ubuntu

```
sudo snap install tunl --edge
```

## Docker

```
docker run --rm --network=host -v $PWD:$PWD tunl/tunl
```

**warning**: this maps the current working directory and means files outside the current working directory cannot be shared
