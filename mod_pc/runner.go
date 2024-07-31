package ModPC

import (
	"Eulogist/minecraft/protocol/packet"
	"fmt"

	"github.com/pterm/pterm"
)

// 在 127.0.0.1:19132 运行一个代理服务器，
// 用于等待 ModPC 连接
func RunServer() *Server {
	var downInitConnect bool
	server := new(Server)
	// prepare
	err := server.CreateListener("127.0.0.1:19132")
	if err != nil {
		panic(fmt.Sprintf("RunServer: %v", err))
	}
	pterm.Success.Println("Server is successful to turn on, now waiting Mod PC to connect.\nServer IP Address: 127.0.0.1:19132")
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
