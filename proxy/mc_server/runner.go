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

// ConnectToServer 用于连接到租赁服号为 serverCode，
// 服务器密码为 serverPassword 的网易租赁服。
// token 指代 FB Token
func ConnectToServer(serverCode string, serverPassword string, token string, authServer string) (*MinecraftServer, error) {
	// 准备
	var downInitConnect bool
	var mcServer MinecraftServer
	mcServer.fbClient = fbauth.CreateClient(&fbauth.ClientOptions{AuthServer: authServer})
	authenticator := fbauth.NewAccessWrapper(mcServer.fbClient, serverCode, serverPassword, token, "", "")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// 生成密钥并发送请求到认证服务器
	clientkey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	armoured_key, _ := x509.MarshalPKIXPublicKey(&clientkey.PublicKey)
	authResponse, err := authenticator.GetAccess(ctx, armoured_key)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 连接到服务器
	connection, err := raknet.DialContext(ctx, authResponse.RentalServerIP)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 设置底层 Raknet 连接
	mcServer.Raknet = RaknetConnection.NewRaknet()
	mcServer.SetConnection(connection, clientkey)
	go mcServer.ProcessIncomingPackets()
	// 向网易租赁服请求网络设置，
	// 这是赞颂者登录到网易租赁服的第一个数据包
	err = mcServer.WritePacket(
		RaknetConnection.MinecraftPacket{
			Packet: &packet.RequestNetworkSettings{ClientProtocol: protocol.CurrentProtocol},
		}, false,
	)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 处理来自 bot 端的登录相关数据包
	for {
		// 读取数据包
		pk := mcServer.ReadPacket()
		// 处理初始连接数据包
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
		// 检查连接状态
		select {
		case <-mcServer.GetContext().Done():
			return nil, fmt.Errorf("ConnectToServer: NetEase Minecraft Rental Server closed their connection to eulogist")
		default:
		}
		// 返回 MinecraftServer 实例
		if downInitConnect {
			return &mcServer, nil
		}
	}
}
