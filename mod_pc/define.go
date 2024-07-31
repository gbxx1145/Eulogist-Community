package ModPC

import (
	RaknetConnection "Eulogist/raknet_connection"

	"github.com/sandertv/go-raknet"
)

type Server struct {
	listener  *raknet.Listener
	connected chan struct{}

	*RaknetConnection.RaknetConnection
}
