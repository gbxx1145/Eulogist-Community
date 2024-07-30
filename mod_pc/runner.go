package ModPC

import (
	"Eulogist/minecraft/protocol/packet"
	"fmt"
)

func RunServer() {
	server := new(Server)

	err := server.CreateListener("127.0.0.1:19132")
	if err != nil {
		panic(fmt.Sprintf("RunServer: %v", err))
	}
	fmt.Println("STARTED")

	go func() {
		err = server.WaitConnect()
		if err != nil {
			panic(fmt.Sprintf("RunServer: %v", err))
		}
		server.ProcessIncomingPackets()
	}()

	<-server.connected

	for {
		pk := server.ReadPacket()

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
		}

		select {
		case <-server.closed:
			fmt.Println("CLOSED")
			return
		default:
		}
	}
}
