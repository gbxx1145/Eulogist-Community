package mc_client

import (
	RaknetConnection "Eulogist/core/raknet"
	"net"

	"github.com/sandertv/go-raknet"
)

type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}
	address   *net.UDPAddr

	*RaknetConnection.Raknet
}
