package mc_server

import (
	fbauth "Eulogist/core/fb_auth/mv4"
	fb_client "Eulogist/core/fb_auth/mv4/client"
	"Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/minecraft/netease/protocol/packet"
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/core/raknet/handshake"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/skin_process"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"time"

	"Eulogist/core/minecraft/netease/raknet"
)

// ConnectToServer 用于连接到租赁服号为 serverCode，
// 服务器密码为 serverPassword 的网易租赁服。
// token 指代 FB Token
func ConnectToServer(basicConfig BasicConfig) (*MinecraftServer, error) {
	// 准备
	var mcServer MinecraftServer
	// 初始化
	mcServer.fbClient = fb_client.CreateClient(basicConfig.AuthServer)
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
		err = skin_process.GetSkinFromAuthResponse(authResponse, mcServer.GetPlayerSkin())
		if err != nil {
			return nil, fmt.Errorf("ConnectToServer: %v", err)
		}
	}
	// 初始化相关数据
	mcServer.SetNeteaseUID(mcServer.fbClient.ClientInfo.Uid)
	mcServer.SetOutfitInfo(authResponse.OutfitInfo)
	// 连接到服务器
	connection, err := raknet.DialContext(ctx, authResponse.RentalServerIP)
	if err != nil {
		return nil, fmt.Errorf("ConnectToServer: %v", err)
	}
	// 设置数据
	mcServer.authResponse = authResponse
	mcServer.Raknet = raknet_connection.NewNetEaseRaknetWrapper()
	mcServer.SetConnection(connection, clientkey)
	go mcServer.ProcessIncomingPackets()
	// 返回值
	return &mcServer, nil
}

/*
FinishHandshake 用于赞颂者完成
与网易租赁服的基本数据包交换。

在与网易租赁服建立 Raknet 连接后，
由赞颂者发送第一个数据包，
用于向服务器请求网络信息设置。

随后，得到来自网易服务器的回应，
并由赞颂者完成基础登录序列，
然后，最终完成与网易租赁服的握手。

此函数应当只被调用一次
*/
func (m *MinecraftServer) FinishHandshake() error {
	// 准备
	var err error
	// 向网易租赁服请求网络设置，
	// 这是赞颂者登录到网易租赁服的第一个数据包
	m.WriteSinglePacket(
		raknet_wrapper.MinecraftPacket[packet.Packet]{
			Packet: &packet.RequestNetworkSettings{ClientProtocol: protocol.CurrentProtocol},
		},
	)
	// 处理来自 bot 端的登录相关数据包
	for {
		for _, pk := range m.ReadPackets() {
			// 处理初始连接数据包
			switch p := pk.Packet.(type) {
			case *packet.NetworkSettings:
				m.identityData, m.clientData, err = handshake.HandleNetworkSettings(m.Raknet, p, m.authResponse, m.playerSkin)
				if err != nil {
					return fmt.Errorf("FinishHandshake: %v", err)
				}
			case *packet.ServerToClientHandshake:
				err = handshake.HandleServerToClientHandshake(m.Raknet, p)
				if err != nil {
					return fmt.Errorf("FinishHandshake: %v", err)
				}
				// 连接已完成初始化，
				// 于是我们返回值
				return nil
			}
		}
		// 检查连接状态
		select {
		case <-m.Context.Done():
			return fmt.Errorf("FinishHandshake: NetEase Minecraft Rental Server closed their connection to eulogist")
		default:
		}
	}
}
