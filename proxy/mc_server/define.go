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
	playerSkin   *SkinProcess.Skin

	entityUniqueID int64

	*RaknetConnection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
