package mc_server

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
func ConnectToServer(serverCode string, serverPassword string, token string, authServer string) (*MinecraftServer, error) {
	var downInitConnect bool
	var mcServer MinecraftServer
	mcServer.fbClient = fbauth.CreateClient(&fbauth.ClientOptions{AuthServer: authServer})
	authenticator := fbauth.NewAccessWrapper(mcServer.fbClient, serverCode, serverPassword, token, "", "")
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
	mcServer.Raknet = RaknetConnection.NewRaknet()
	mcServer.SetConnection(connection, clientkey)
	go mcServer.ProcessIncomingPackets()
	// set connection
	err = mcServer.WritePacket(
		RaknetConnection.MinecraftPacket{
			Packet: &packet.RequestNetworkSettings{ClientProtocol: protocol.CurrentProtocol},
		}, false,
	)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// request network settings
	for {
		pk := mcServer.ReadPacket()
		switch p := pk.Packet.(type) {
		case *packet.NetworkSettings:
			err = mcServer.HandleNetworkSettings(p, authResponse)
			if err != nil {
				return nil, fmt.Errorf("ConnectToServer: %v", err)
			}
		case *packet.ServerToClientHandshake:
			err = mcServer.HandleServerToClientHandshake(p)
			if err != nil {
				return nil, fmt.Errorf("ConnectToServer: %v", err)
			}
			downInitConnect = true
		}
		// handle init connection packets
		select {
		case <-mcServer.GetContext().Done():
			return nil, fmt.Errorf("ConnectToServer: NetEase Minecraft Rental Server closed their connection to eulogist")
		default:
		}
		// check connection states
		if downInitConnect {
			return &mcServer, nil
		}
		// return
	}
	// process login related packets from bot side
}
