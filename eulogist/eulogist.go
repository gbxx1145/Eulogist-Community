package Eulogist

import (
	RaknetConnection "Eulogist/core/raknet"
	BotSide "Eulogist/server/bot_side"
	ModPC "Eulogist/server/mod_pc"
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
	var modPC *ModPC.Server
	var botSide *BotSide.BotSide
	var modPCWasConnected chan struct{}

	{
		config, err = ReadEulogistConfig()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		// read config
		if config.LaunchType == LaunchTypeNormal {
			if !FileExist(config.NEMCPath) {
				return fmt.Errorf("Eulogist: Mod PC not found, maybe you did not download or the the path are incorrect")
			}
			// check mod pc is download
			neteaseConfigPath, err = GenerateNetEaseConfig(config)
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			// generate netease config
			modPC, modPCWasConnected, err = ModPC.RunServer(config.ServerIP, config.ServerPort)
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			// run server
			command := exec.Command(config.NEMCPath)
			command.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf("%#v config=%#v", config.NEMCPath, neteaseConfigPath)}
			go command.Run()
			pterm.Success.Println("Success to turn on Mod PC and Eulogist, now waiting Mod PC to connect.")
			// launch mod pc
		} else {
			modPC, modPCWasConnected, err = ModPC.RunServer(config.ServerIP, config.ServerPort)
			if err != nil {
				return fmt.Errorf("Eulogist: %v", err)
			}
			pterm.Success.Printf("Eulogist is successful to turn on, now waiting Mod PC to connect.\nEulogist Server IP Address: %s:%d\n", config.ServerIP, config.ServerPort)
		}
		// run eulogist
	}

	{
		if config.LaunchType == LaunchTypeNormal {
			timer := time.NewTimer(time.Second * 120)
			defer timer.Stop()
			select {
			case <-timer.C:
				return fmt.Errorf("Eulogist: Failed to create connection with Mod PC")
			case <-modPCWasConnected:
				close(modPCWasConnected)
			}
		} else {
			<-modPCWasConnected
			close(modPCWasConnected)
		}
		pterm.Success.Println("Success to create connection with Mod PC, now we try to create handshake with it.")
		// waiting mod pc to connect
		err = modPC.WaitClientHandshakeDown()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with Mod PC, now we try to communicate with auth server.")
		// finish mod pc handshake
	}

	{
		botSide, err = BotSide.ConnectToServer(config.RentalServerCode, config.RentalServerPassword, config.FBToken, LookUpAuthServerAddress(config.AuthServerType))
		if err != nil {
			return fmt.Errorf("UnfoldEulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with NetEase Minecraft Rental Server, and then you will login to it.")
	}
	// create connection with bot side

	waitGroup.Add(2)
	// set wait group

	go func() {
		defer func() {
			botSide.CloseConnection()
			modPC.CloseConnection()
			waitGroup.Add(-1)
		}()
		// ...
		for {
			pk := botSide.ReadPacket()
			if err != nil {
				return
			}
			shouldSendCopy, err := botSide.PacketFilter(pk.Packet)
			if err != nil {
				pterm.Warning.Printf("UnfoldEulogist: %v\n", err)
				continue
			}
			// filte the packets
			if shouldSendCopy {
				err = modPC.WritePacket(RaknetConnection.MinecraftPacket{Bytes: pk.Bytes}, true)
				if err != nil {
					return
				}
			}
			// send a copy to mod pc
			if shieldID := botSide.GetShieldID(); shieldID != 0 {
				modPC.SetShieldID(shieldID)
			}
			// sync shield id
			if !botSide.GetShouldDecode() {
				modPC.SetShouldDecode(false)
			}
			// sync should decode states
			select {
			case <-botSide.GetContext().Done():
				return
			case <-modPC.GetContext().Done():
				return
			default:
			}
			// check connection states
		}
	}()
	// bot side <-> eulogist

	go func() {
		defer func() {
			modPC.CloseConnection()
			botSide.CloseConnection()
			waitGroup.Add(-1)
		}()
		// ...
		for {
			pk := modPC.ReadPacket()
			// read packet from mod pc
			err = botSide.WritePacket(RaknetConnection.MinecraftPacket{Bytes: pk.Bytes}, true)
			if err != nil {
				return
			}
			// sync packet to bot side
			select {
			case <-modPC.GetContext().Done():
				return
			case <-botSide.GetContext().Done():
				return
			default:
			}
			// check connection states
		}
	}()
	// mod pc <-> eulogist

	waitGroup.Wait()
	pterm.Info.Println("Server Down. Now all connection was closed.")
	return nil
}
