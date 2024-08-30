package mc_server

import (
	fbauth "Eulogist/core/fb_auth/mv4"
	"Eulogist/core/minecraft/protocol/login"
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/core/tools/skin_process"
)

type MinecraftServer struct {
	fbClient              *fbauth.Client
	authResponse          fbauth.AuthResponse
	getCheckNumEverPassed bool

	identityData *login.IdentityData
	clientData   *login.ClientData

	neteaseUID string
	playerSkin *skin_process.Skin
	outfitInfo map[string]*int

	entityUniqueID  int64
	entityRuntimeID uint64

	*raknet_connection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
