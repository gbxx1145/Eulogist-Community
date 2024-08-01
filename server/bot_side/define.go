package BotSide

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	RaknetConnection "Eulogist/core/raknet"
)

type BotSide struct {
	fbClient       *fbauth.Client
	entityUniqueID int64

	getCheckNumEverPassed bool

	*RaknetConnection.RaknetConnection
}
