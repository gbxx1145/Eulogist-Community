package ModPC

import (
	"Eulogist/core/minecraft/protocol/packet"
	"fmt"
)

// 在 serverIP 对应的 IP 上运行一个代理服务器以等待
// Mod PC 连接，并指定该服务器开放的端口为 serverPort。
// 当 Mod PC 连接时，管道 connected 将收到数据
func RunServer(serverIP string, serverPort int) (server *Server, connected chan struct{}, err error) {
	server = new(Server)
	// prepare
	err = server.CreateListener(fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		return nil, nil, fmt.Errorf("RunServer: %v", err)
	}
	// open server
	go func() {
		err = server.WaitConnect()
		if err != nil {
			panic(fmt.Sprintf("RunServer: %v", err))
		}
		server.ProcessIncomingPackets()
	}()
	// wait connect and start listening
	return server, server.connected, nil
	// return
}

// 等待 Mod PC 完成与 赞颂者 的基本数据包交换。
// 此函数应当只被调用一次
func (s *Server) WaitClientHandshakeDown() error {
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
			return fmt.Errorf("WaitClientHandshakeDown: Mod PC closed their connection to eulogist")
		default:
		}
		// check connection states
		if downInitConnect {
			return nil
		}
		// return
	}
	// process login related packets from mod pc
}
