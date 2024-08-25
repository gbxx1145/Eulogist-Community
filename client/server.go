package client

import (
	"Eulogist/minecraft/protocol/packet"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/sandertv/go-raknet"
)

// ...
func (s *Server) CreateListener(address string) error {
	listener, err := raknet.Listen(address)
	if err != nil {
		return fmt.Errorf("CreateListener: %v", err)
	}
	s.listener = listener
	s.connected = make(chan struct{}, 1)
	s.closed = make(chan struct{}, 1)
	return nil
}

// ...
func (s *Server) WaitConnect() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return fmt.Errorf("WaitConnect: %v", err)
	}
	// accept connection
	s.connection = conn
	s.encoder = packet.NewEncoder(conn)
	s.decoder = packet.NewDecoder(conn)
	s.packets = make(chan MinecraftPacket, 1024)
	s.serverKey, _ = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	_, _ = rand.Read(s.salt)
	s.connected <- struct{}{}
	// set value
	return nil
	// return
}

// ...
func (s *Server) ProcessIncomingPackets() {
	// Incomplete implementation
}

// ...
func (s *Server) ReadPacket() MinecraftPacket {
	// Incomplete implementation
	return MinecraftPacket{}
}

// ...
func (s *Server) WritePacket(pk MinecraftPacket, useBytes bool) error {
	// Incomplete implementation
	return nil
}
