package mc_client

import (
	"Eulogist/core/minecraft/protocol/packet"
	"fmt"
)

// RunServer 在 127.0.0.1 上运行一个代理服务器以等待
// Minecraft 连接。服务器开放的端口将被自动设置，
// 您可以使用 client.GetServerPort() 来取得开放的端口。
// 当 Minecraft 连接时，管道 connected 将收到数据
func RunServer() (client *MinecraftClient, connected chan struct{}, err error) {
	// prepare
	client = new(MinecraftClient)
	// start listening
	err = client.CreateListener()
	if err != nil {
		return nil, nil, fmt.Errorf("RunServer: %v", err)
	}
	// wait Minecraft Client to connect
	go func() {
		err = client.WaitConnect()
		if err != nil {
			panic(fmt.Sprintf("RunServer: %v", err))
		}
		client.ProcessIncomingPackets()
	}()
	// return
	return client, client.connected, nil
}

// WaitClientHandshakeDown 等待 Minecraft
// 完成与 赞颂者 的基本数据包交换。
// 此函数应当只被调用一次
func (m *MinecraftClient) WaitClientHandshakeDown() error {
	// prepare
	var downInitConnect bool
	var err error
	// process login related packets from Minecraft
	for {
		// read packet
		pk := m.ReadPacket()
		// handle login related packets
		switch p := pk.Packet.(type) {
		case *packet.RequestNetworkSettings:
			err = m.HandleRequestNetworkSettings(p)
			if err != nil {
				panic(fmt.Sprintf("WaitClientHandshakeDown: %v", err))
			}
		case *packet.Login:
			m.identityData, m.clientData, err = m.HandleLogin(p)
			if err != nil {
				panic(fmt.Sprintf("WaitClientHandshakeDown: %v", err))
			}
		case *packet.ClientToServerHandshake:
			downInitConnect = true
		}
		// check connection states
		select {
		case <-m.GetContext().Done():
			return fmt.Errorf("WaitClientHandshakeDown: Minecraft closed its connection to eulogist")
		default:
		}
		// return
		if downInitConnect {
			return nil
		}
	}
}
