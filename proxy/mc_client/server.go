package mc_client

import (
	RaknetConnection "Eulogist/core/raknet"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/sandertv/go-raknet"
)

// 在指定的 IP 地址及端口号上创建 Raknet 侦听连接。
// 任何 Minecraft 客户端都可以连接到该地址
func (s *MCClient) CreateListener() error {
	listener, err := raknet.Listen("127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("CreateListener: %v", err)
	}
	s.listener = listener
	s.connected = make(chan struct{}, 1)
	s.RaknetConnection = RaknetConnection.NewRaknetConnection()
	return nil
}

// 等待 Minecraft 客户端连接到服务器
func (s *MCClient) WaitConnect() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return fmt.Errorf("WaitConnect: %v", err)
	}
	// accept connection
	serverKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	s.SetConnection(conn, serverKey)
	s.connected <- struct{}{}
	// set value
	return nil
	// return
}
