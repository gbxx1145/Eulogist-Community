package client

import (
	"Eulogist/minecraft/protocol/packet"
	"crypto/ecdsa"
	"net"

	"github.com/sandertv/go-raknet"
)

type Server struct {
	listener *raknet.Listener

	connection net.Conn
	connected  chan struct{}
	closed     chan struct{}

	encoder   *packet.Encoder
	decoder   *packet.Decoder
	serverKey *ecdsa.PrivateKey
	salt      []byte

	packets chan (MinecraftPacket)
}

type MinecraftPacket struct {
	Packet packet.Packet
	Bytes  []byte
}
