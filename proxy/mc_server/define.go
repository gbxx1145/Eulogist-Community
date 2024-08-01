package mc_server

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	RaknetConnection "Eulogist/core/raknet"
)

type MCServer struct {
	fbClient       *fbauth.Client
	entityUniqueID int64

	getCheckNumEverPassed bool

	*RaknetConnection.RaknetConnection
}
