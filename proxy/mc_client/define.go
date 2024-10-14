package mc_client

import (
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/proxy/persistence_data"
	"net"

	"Eulogist/core/minecraft/raknet"
)

type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}

	Address         *net.UDPAddr
	PersistenceData *persistence_data.PersistenceData

	*raknet_connection.Raknet
}
