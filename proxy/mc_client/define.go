package mc_client

import (
	RaknetConnection "Eulogist/core/raknet"

	"github.com/sandertv/go-raknet"
)

type MCClient struct {
	listener  *raknet.Listener
	connected chan struct{}

	IP   string
	Port int
	*RaknetConnection.RaknetConnection
}
