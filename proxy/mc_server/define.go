package mc_server

import (
	fbauth "Eulogist/core/fb_auth/mv4"
	"Eulogist/core/minecraft/protocol/login"
	RaknetConnection "Eulogist/core/raknet"
	SkinProcess "Eulogist/core/tools/skin_process"
)

type MinecraftServer struct {
	fbClient              *fbauth.Client
	authResponse          fbauth.AuthResponse
	getCheckNumEverPassed bool

	identityData *login.IdentityData
	clientData   *login.ClientData

	neteaseUID string
	playerSkin *SkinProcess.Skin
	outfitInfo map[string]*int

	entityUniqueID  int64
	entityRuntimeID uint64

	*RaknetConnection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
