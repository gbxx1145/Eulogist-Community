package ModPC

import (
	"Eulogist/minecraft/protocol/packet"
	"fmt"

	"github.com/pterm/pterm"
)

// 在 127.0.0.1 运行一个代理服务器以等待
// Mod PC 连接，
// 并指定该服务器开放的端口为 runningPort
func RunServer(runningPort uint16) *Server {
	var downInitConnect bool
	server := new(Server)
	// prepare
	err := server.CreateListener(fmt.Sprintf("127.0.0.1:%d", runningPort))
	if err != nil {
		panic(fmt.Sprintf("RunServer: %v", err))
	}
	pterm.Success.Printf("Server is successful to turn on, now waiting Mod PC to connect.\nServer IP Address: 127.0.0.1:%d\n", runningPort)
	// open server
	go func() {
		err = server.WaitConnect()
		if err != nil {
			panic(fmt.Sprintf("RunServer: %v", err))
		}
		server.ProcessIncomingPackets()
	}()
	<-server.connected
	close(server.connected)
	// wait connect and process packets
	for {
		pk := server.ReadPacket()
		// read packet
		switch p := pk.Packet.(type) {
		case *packet.RequestNetworkSettings:
			err = server.HandleRequestNetworkSettings(p)
			if err != nil {
				panic(fmt.Sprintf("RunServer: %v", err))
			}
		case *packet.Login:
			err = server.HandleLogin(p)
			if err != nil {
				panic(fmt.Sprintf("RunServer: %v", err))
			}
		case *packet.ClientToServerHandshake:
			downInitConnect = true
		}
		// handle init connection packets
		select {
		case <-server.GetContext().Done():
			return nil
		default:
		}
		// check connection states
		if downInitConnect {
			return server
		}
		// return
	}
	// process login related packets from mod pc
}
