package mc_client

import (
	"Eulogist/core/minecraft/protocol/packet"
	"fmt"
	"net"
)

// 在 serverIP 对应的 IP 上运行一个代理服务器以等待
// Minecraft 连接，并指定该服务器开放的端口为 serverPort。
// 当 Minecraft 连接时，管道 connected 将收到数据
func RunServer() (client *MCClient, connected chan struct{}, err error) {
	client = new(MCClient)
	// prepare
	err = client.CreateListener()
	if err != nil {
		return nil, nil, fmt.Errorf("RunServer: %v", err)
	}
	// open client
	go func() {
		err = client.WaitConnect()
		if err != nil {
			panic(fmt.Sprintf("RunServer: %v", err))
		}
		client.ProcessIncomingPackets()
	}()
	// wait connect and start listening
	addr, ok := client.listener.Addr().(*net.UDPAddr)
	if !ok {
		return nil, nil, fmt.Errorf("RunServer: %v", "failed to get address")
	}
	client.IP = addr.IP.String()
	client.Port = addr.Port
	// obtain addr and port
	return client, client.connected, nil
	// return
}

// 等待 Minecraft 完成与 赞颂者 的基本数据包交换。
// 此函数应当只被调用一次
func (s *MCClient) WaitClientHandshakeDown() error {
	var downInitConnect bool
	// prepare
	for {
		pk := s.ReadPacket()
		// read packet
		switch p := pk.Packet.(type) {
		case *packet.RequestNetworkSettings:
			err := s.HandleRequestNetworkSettings(p)
			if err != nil {
				panic(fmt.Sprintf("WaitClientHandshakeDown: %v", err))
			}
		case *packet.Login:
			err := s.HandleLogin(p)
			if err != nil {
				panic(fmt.Sprintf("WaitClientHandshakeDown: %v", err))
			}
		case *packet.ClientToServerHandshake:
			downInitConnect = true
		}
		// handle init connection packets
		select {
		case <-s.GetContext().Done():
			return fmt.Errorf("WaitClientHandshakeDown: Minecraft closed its connection to eulogist")
		default:
		}
		// check connection states
		if downInitConnect {
			return nil
		}
		// return
	}
	// process login related packets from Minecraft
}
