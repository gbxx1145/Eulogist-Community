package mc_server

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	SkinProcess "Eulogist/core/tools/skin_process"
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
func ConnectToServer(basicConfig BasicConfig) (*MinecraftServer, error) {
	// 准备
	var mcServer MinecraftServer
	// 初始化
	mcServer.fbClient = fbauth.CreateClient(&fbauth.ClientOptions{AuthServer: basicConfig.AuthServer})
	authenticator := fbauth.NewAccessWrapper(
		mcServer.fbClient, basicConfig.ServerCode, basicConfig.ServerPassword, basicConfig.Token, "", "",
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// 向验证服务器请求信息
	clientkey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	armoured_key, _ := x509.MarshalPKIXPublicKey(&clientkey.PublicKey)
	authResponse, err := authenticator.GetAccess(ctx, armoured_key)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 初始化皮肤信息
	if len(authResponse.SkinInfo.SkinDownloadURL) > 0 {
		mcServer.InitPlayerSkin()
		err = SkinProcess.GetSkinFromAuthResponse(authResponse, mcServer.GetPlayerSkin())
		if err != nil {
			return nil, fmt.Errorf("ConnectToServer: %v", err)
		}
	}
	// 连接到服务器
	connection, err := raknet.DialContext(ctx, authResponse.RentalServerIP)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 设置数据
	mcServer.authResponse = authResponse
	mcServer.Raknet = RaknetConnection.NewRaknet()
	mcServer.SetConnection(connection, clientkey)
	go mcServer.ProcessIncomingPackets()
	// 返回值
	return &mcServer, nil
}

// WaitClientHandshakeDown 等待赞颂者
// 完成与网易租赁服的基本数据包交换。
// 此函数应当只被调用一次
func (m *MinecraftServer) WaitClientHandshakeDown() error {
	// 准备
	var downInitConnect bool
	var err error
	// 向网易租赁服请求网络设置，
	// 这是赞颂者登录到网易租赁服的第一个数据包
	m.WriteSinglePacket(
		RaknetConnection.MinecraftPacket{
			Packet: &packet.RequestNetworkSettings{ClientProtocol: protocol.CurrentProtocol},
		},
	)
	// 处理来自 bot 端的登录相关数据包
	for {
		for _, pk := range m.ReadPackets() {
			// 处理初始连接数据包
			switch p := pk.Packet.(type) {
			case *packet.NetworkSettings:
				m.identityData, m.clientData, err = m.HandleNetworkSettings(p, m.authResponse, m.playerSkin)
				if err != nil {
					return fmt.Errorf("ConnectToServer: %v", err)
				}
			case *packet.ServerToClientHandshake:
				err = m.HandleServerToClientHandshake(p)
				if err != nil {
					return fmt.Errorf("ConnectToServer: %v", err)
				}
				downInitConnect = true
			}
			// 检查连接状态
			select {
			case <-m.GetContext().Done():
				return fmt.Errorf("ConnectToServer: NetEase Minecraft Rental Server closed their connection to eulogist")
			default:
			}
			// 连接已完成初始化，
			// 于是我们返回值
			if downInitConnect {
				return nil
			}
		}
	}
}
