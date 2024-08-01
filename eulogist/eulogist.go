package Eulogist

import (
	RaknetConnection "Eulogist/core/raknet"
	Client "Eulogist/proxy/mc_client"
	Server "Eulogist/proxy/mc_server"
	"fmt"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/pterm/pterm"
)

// 展开 “赞颂者”
func Eulogist() error {
	var err error
	var config *EulogistConfig
	var neteaseConfigPath string
	var waitGroup sync.WaitGroup
	var client *Client.MinecraftClient
	var server *Server.MinecraftServer
	var ClientWasConnected chan struct{}

	{
		config, err = ReadEulogistConfig()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		// read config
		if config.LaunchType == LaunchTypeNormal {
			if !FileExist(config.NEMCPath) {
				return fmt.Errorf("Eulogist: Client not found, maybe you did not download or the the path is incorrect")
			}
			// check Minecraft is download
			client, ClientWasConnected, err = Client.RunServer()
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			// run server
			neteaseConfigPath, err = GenerateNetEaseConfig(config, client.GetServerIP(), client.GetServerPort())
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			// generate netease config
			command := exec.Command(config.NEMCPath)
			command.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf("%#v config=%#v", config.NEMCPath, neteaseConfigPath)}
			go command.Run()
			pterm.Success.Println("Eulogist is ready! Now we are going to start Minecarft Client.\nThen, the Minecraft Client will connect to Eulogist automatically.")
			// launch Minecraft
		} else {
			client, ClientWasConnected, err = Client.RunServer()
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			pterm.Success.Printf(
				"Eulogist is ready! Please connect to Eulogist manually.\nEulogist server address: %s:%d\n",
				client.GetServerIP(), client.GetServerPort(),
			)
		}
		// run eulogist
	}

	{
		if config.LaunchType == LaunchTypeNormal {
			timer := time.NewTimer(time.Second * 120)
			defer timer.Stop()
			select {
			case <-timer.C:
				return fmt.Errorf("Eulogist: Failed to create connection with Minecraft")
			case <-ClientWasConnected:
				close(ClientWasConnected)
			}
		} else {
			<-ClientWasConnected
			close(ClientWasConnected)
		}
		pterm.Success.Println("Success to create connection with Minecraft Client, now we try to create handshake with it.")
		// waiting Minecraft to connect
		err = client.WaitClientHandshakeDown()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with Minecraft Client, now we try to communicate with auth server.")
		// finish Minecraft handshake
	}

	{
		server, err = Server.ConnectToServer(config.RentalServerCode, config.RentalServerPassword, config.FBToken, LookUpAuthServerAddress(config.FBToken))
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with NetEase Minecraft Rental Server, and then you will login to it.")
	}
	// create connection with bot side

	waitGroup.Add(2)
	// set wait group

	go func() {
		defer func() {
			server.CloseConnection()
			client.CloseConnection()
			waitGroup.Add(-1)
		}()
		// ...
		for {
			pk := server.ReadPacket()
			if err != nil {
				return
			}
			shouldSendCopy, err := server.PacketFilter(pk.Packet)
			if err != nil {
				pterm.Warning.Printf("Eulogist: %v\n", err)
				continue
			}
			// filte the packets
			if shouldSendCopy {
				err = client.WritePacket(RaknetConnection.MinecraftPacket{Bytes: pk.Bytes}, true)
				if err != nil {
					return
				}
			}
			// send a copy to Minecraft
			if shieldID := server.GetShieldID(); shieldID != 0 {
				client.SetShieldID(shieldID)
			}
			// sync shield id
			if !server.GetShouldDecode() {
				client.SetShouldDecode(false)
			}
			// sync should decode states
			select {
			case <-server.GetContext().Done():
				return
			case <-client.GetContext().Done():
				return
			default:
			}
			// check connection states
		}
	}()
	// bot side <-> eulogist

	go func() {
		defer func() {
			client.CloseConnection()
			server.CloseConnection()
			waitGroup.Add(-1)
		}()
		// ...
		for {
			pk := client.ReadPacket()
			// read packet from Minecraft
			err = server.WritePacket(RaknetConnection.MinecraftPacket{Bytes: pk.Bytes}, true)
			if err != nil {
				return
			}
			// sync packet to bot side
			select {
			case <-client.GetContext().Done():
				return
			case <-server.GetContext().Done():
				return
			default:
			}
			// check connection states
		}
	}()
	// Minecraft <-> eulogist

	waitGroup.Wait()
	pterm.Info.Println("Server Down. Now all connection was closed.")
	return nil
}
