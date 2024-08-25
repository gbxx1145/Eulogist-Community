package client

import (
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

		switch pk.Packet.(type) {
		// DO SOMETHING;
		// Incomplete implementation
		}

		select {
		case <-server.closed:
			fmt.Println("CLOSED")
			return
		default:
		}
	}
}
