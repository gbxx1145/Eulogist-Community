package mc_client

import (
	"Eulogist/core/minecraft/protocol/login"
	RaknetConnection "Eulogist/core/raknet"
	"net"

	"github.com/sandertv/go-raknet"
)

type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}
	address   *net.UDPAddr

	identityData *login.IdentityData
	clientData   *login.ClientData
	playerSkin   *RaknetConnection.Skin

	entityUniqueID int64

	*RaknetConnection.Raknet
}
