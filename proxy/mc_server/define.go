package mc_server

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	RaknetConnection "Eulogist/core/raknet"
)

type MinecraftServer struct {
	fbClient              *fbauth.Client
	authResponse          fbauth.AuthResponse
	getCheckNumEverPassed bool

	entityUniqueID int64
	playerSkin     *RaknetConnection.Skin

	*RaknetConnection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
