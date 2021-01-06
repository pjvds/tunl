package client

type TokenStore interface {
	TokenRetriever
	TokenRefresher
}

type TokenRetriever interface {
	GetToken() string
}

type TokenRefresher interface {
	SetToken(token string)
}
