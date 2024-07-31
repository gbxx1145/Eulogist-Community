package main

import (
	BotSide "Eulogist/bot_side"
	ModPC "Eulogist/mod_pc"
	RaknetConnection "Eulogist/raknet_connection"
	"fmt"
	"sync"

	"github.com/pterm/pterm"
)

// 展开 赞颂者。
//
// serverCode 和 serverPassword
// 分别代表要赞颂的租赁服编号及其密码。
//
// token 指代 FB Token
func UnfoldEulogist(serverCode string, serverPassword string, token string, authServer string) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	modPC, err := ModPC.RunServer(19132)
	if err != nil {
		return fmt.Errorf("UnfoldEulogist: %v", err)
	}
	pterm.Success.Println("Success to create connection with Mod PC, now we try to communicate with auth server.")

	botSide, err := BotSide.ConnectToServer(serverCode, serverPassword, token, authServer)
	if err != nil {
		return fmt.Errorf("UnfoldEulogist: %v", err)
	}
	pterm.Success.Println("Success to create handshake with NetEase Minecraft Rental Server, and then you will login to it.")

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
	pterm.Info.Printfln("Server Down. Now all connection was closed.")
	return nil
}
