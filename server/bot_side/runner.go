package BotSide

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/sandertv/go-raknet"
)

// 连接到租赁服号为 serverCode，
// 服务器密码为 serverPassword 的网易租赁服。
// token 指代 FB Token
func ConnectToServer(serverCode string, serverPassword string, token string, authServer string) (*BotSide, error) {
	var downInitConnect bool
	var botSide BotSide
	botSide.fbClient = fbauth.CreateClient(&fbauth.ClientOptions{AuthServer: authServer})
	authenticator := fbauth.NewAccessWrapper(botSide.fbClient, serverCode, serverPassword, token, "", "")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	// prepare
	clientkey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	armoured_key, _ := x509.MarshalPKIXPublicKey(&clientkey.PublicKey)
	authResponse, err := authenticator.GetAccess(ctx, armoured_key)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// generate key and send request to auth server
	connection, err := raknet.DialContext(ctx, authResponse.RentalServerIP)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// connect to server
	botSide.RaknetConnection = RaknetConnection.NewRaknetConnection()
	botSide.SetConnection(connection, clientkey)
	go botSide.ProcessIncomingPackets()
	// set connection
	err = botSide.WritePacket(
		RaknetConnection.MinecraftPacket{
			Packet: &packet.RequestNetworkSettings{ClientProtocol: protocol.CurrentProtocol},
		}, false,
	)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// request network settings
	for {
		pk := botSide.ReadPacket()
		switch p := pk.Packet.(type) {
		case *packet.NetworkSettings:
			err = botSide.HandleNetworkSettings(p, authResponse)
			if err != nil {
				return nil, fmt.Errorf("ConnectToServer: %v", err)
			}
		case *packet.ServerToClientHandshake:
			err = botSide.HandleServerToClientHandshake(p)
			if err != nil {
				return nil, fmt.Errorf("ConnectToServer: %v", err)
			}
			downInitConnect = true
		}
		// handle init connection packets
		select {
		case <-botSide.GetContext().Done():
			return nil, fmt.Errorf("ConnectToServer: NetEase Minecraft Rental Server closed their connection to eulogist")
		default:
		}
		// check connection states
		if downInitConnect {
			return &botSide, nil
		}
		// return
	}
	// process login related packets from bot side
}
