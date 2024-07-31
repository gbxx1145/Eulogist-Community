package BotSide

import (
	fbauth "Eulogist/fb_auth/pv4"
	"Eulogist/minecraft"
	RaknetConnection "Eulogist/raknet_connection"
)

type BotSide struct {
	fbClient              *fbauth.Client
	getCheckNumEverPassed bool

	gameData  minecraft.GameData
	connected chan struct{}

	*RaknetConnection.RaknetConnection
}
