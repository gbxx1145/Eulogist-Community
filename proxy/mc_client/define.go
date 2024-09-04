package mc_client

import (
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/skin_process"
	"net"

	"Eulogist/core/standard/protocol/login"
	"Eulogist/core/standard/protocol/packet"

	"github.com/sandertv/go-raknet"
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

	*raknet_wrapper.Raknet[packet.Packet]
}
