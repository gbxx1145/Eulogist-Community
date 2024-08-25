package client

import (
	"Eulogist/minecraft/protocol/packet"
	"crypto/ecdsa"
)

// ...
func (s *Server) HandleRequestNetworkSettings(pk *packet.RequestNetworkSettings) error {
	// Incomplete implementation
	return nil
}

// ...
func (s *Server) HandleLogin(pk *packet.Login) error {
	// Incomplete implementation
	return nil
}

// ...
func (s *Server) EnableEncryption(clientPublicKey *ecdsa.PublicKey) error {
	// Incomplete implementation
	return nil
}
