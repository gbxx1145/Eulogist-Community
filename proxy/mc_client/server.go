package mc_client

import (
	RaknetConnection "Eulogist/core/raknet"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net"

	"github.com/sandertv/go-raknet"
)

// 在 127.0.0.1 上以 Raknet 协议侦听 Minecraft 客户端的连接，
// 这意味着您成功创建了一个 Minecraft 数据包代理服务器。
// 稍后，您可以通过 s.GetServerAddress 来取得服务器地址
func (m *MinecraftClient) CreateListener() error {
	listener, err := raknet.Listen("127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("CreateListener: %v", err)
	}
	// get listener
	address, ok := listener.Addr().(*net.UDPAddr)
	if !ok {
		return fmt.Errorf("CreateListener: Failed to get address for listener")
	}
	// get listen address
	m.listener = listener
	m.address = address
	m.connected = make(chan struct{}, 1)
	m.Raknet = RaknetConnection.NewRaknet()
	// set value
	return nil
	// return
}

// 等待 Minecraft 客户端连接到服务器
func (m *MinecraftClient) WaitConnect() error {
	conn, err := m.listener.Accept()
	if err != nil {
		return fmt.Errorf("WaitConnect: %v", err)
	}
	// accept connection
	serverKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	m.SetConnection(conn, serverKey)
	m.connected <- struct{}{}
	// set value
	return nil
	// return
}

// ...
func (m *MinecraftClient) GetServerIP() string {
	return m.address.IP.String()
}

// ...
func (m *MinecraftClient) GetServerPort() int {
	return m.address.Port
}
