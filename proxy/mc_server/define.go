package mc_server

import (
	fb_client "Eulogist/core/fb_auth/mv4/client"
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/proxy/persistence_data"
)

type MinecraftServer struct {
	fbClient              *fb_client.Client
	authResponse          *fb_client.AuthResponse
	getCheckNumEverPassed bool

	PersistenceData *persistence_data.PersistenceData
	*raknet_connection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
