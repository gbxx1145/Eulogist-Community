package mc_client

import (
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/proxy/persistence_data"
	"net"

	"Eulogist/core/minecraft/standard/protocol/packet"

	"github.com/sandertv/go-raknet"
)

// ...
type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}

	Address         *net.UDPAddr
	PersistenceData *persistence_data.PersistenceData

	Conn *raknet_wrapper.Raknet[packet.Packet]
}
