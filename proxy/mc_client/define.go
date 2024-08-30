package mc_client

import (
	"Eulogist/core/minecraft/protocol/login"
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/core/tools/skin_process"
	"net"

	"Eulogist/core/minecraft/raknet"
)

type MinecraftClient struct {
	listener  *raknet.Listener
	connected chan struct{}
	address   *net.UDPAddr

	identityData *login.IdentityData
	clientData   *login.ClientData

	neteaseUID string
	playerSkin *skin_process.Skin
	outfitInfo map[string]*int

	entityUniqueID  int64
	entityRuntimeID uint64

	*raknet_connection.Raknet
}
