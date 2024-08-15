package mc_client

import (
	"Eulogist/core/minecraft/protocol/login"
	RaknetConnection "Eulogist/core/raknet"
	SkinProcess "Eulogist/core/tools/skin_process"
	"net"

	"github.com/sandertv/go-raknet"
)

type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}
	address   *net.UDPAddr

	identityData *login.IdentityData
	clientData   *login.ClientData
	playerSkin   *SkinProcess.Skin

	entityUniqueID  int64
	entityRuntimeID uint64

	*RaknetConnection.Raknet
}
