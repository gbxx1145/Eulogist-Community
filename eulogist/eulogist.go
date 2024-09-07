package Eulogist

import (
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"
	Client "Eulogist/proxy/mc_client"
	Server "Eulogist/proxy/mc_server"
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/pterm/pterm"
)

// Eulogist 函数是整个“赞颂者”程序的入口点
func Eulogist() error {
	var err error
	var config *EulogistConfig
	var waitGroup sync.WaitGroup
	var client *Client.MinecraftClient
	var server *Server.MinecraftServer
	var clientWasConnected chan struct{}

	// 读取配置文件
	{
		config, err = ReadEulogistConfig()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
	}

	// 使赞颂者连接到网易租赁服
	{
		pterm.Info.Println("Now we try to communicate with Auth Server.")

		server, err = Server.ConnectToServer(
			Server.BasicConfig{
				ServerCode:     config.RentalServerCode,
				ServerPassword: config.RentalServerPassword,
				Token:          config.FBToken,
				AuthServer:     LookUpAuthServerAddress(config.FBToken),
			},
		)
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		defer server.CloseConnection()

		pterm.Success.Println("Success to create connection with NetEase Minecraft Bedrock Rental Server, now we try to create handshake with it.")

		err = server.FinishHandshake()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}

		pterm.Success.Println("Success to create handshake with NetEase Minecraft Bedrock Rental Server.")
	}

	// 召唤——赞颂者
	{
		// 启动赞颂者
		client, clientWasConnected, err = Client.RunServer()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		defer client.CloseConnection()
		// 打印赞颂者准备完成的信息
		pterm.Success.Printf(
			"Eulogist is ready! Please connect to Eulogist manually.\nEulogist server address: %s:%d\n",
			client.GetServerIP(), client.GetServerPort(),
		)
	}

	// 等待 Minecraft 客户端与赞颂者完成基本数据包交换
	{
		// 等待 Minecraft 客户端连接
		<-clientWasConnected
		close(clientWasConnected)
		pterm.Success.Println("Success to create connection with Minecraft Client, now we try to create handshake with it.")
		// 等待 Minecraft 客户端完成握手
		err = client.WaitClientHandshakeDown()
		if err != nil {
			return fmt.Errorf("Eulogist: %v", err)
		}
		pterm.Success.Println("Success to create handshake with Minecraft Client, and then you will login to NetEase Minecraft Bedrock Rental Server.")
	}

	// 同步相关数据，
	// 并设置等待队列
	client.SetNeteaseUID(server.GetNeteaseUID())
	client.SetPlayerSkin(server.GetPlayerSkin())
	client.SetOutfitInfo(server.GetOutfitInfo())
	server.SetStandardBedrockIdentity(client.GetStandardBedrockIdentity())
	waitGroup.Add(2)

	// 处理网易租赁服到赞颂者的数据包
	go func() {
		// 关闭已建立的所有连接
		defer func() {
			waitGroup.Add(-1)
			server.CloseConnection()
			client.CloseConnection()
		}()
		// 显示程序崩溃错误信息
		defer func() {
			r := recover()
			if r != nil {
				pterm.Error.Printf(
					"Eulogist/GoFunc/RentalServerToEulogist: err = %v\n\n[Stack Info]\n%s\n",
					r, string(debug.Stack()),
				)
				fmt.Println()
			}
		}()
		// 数据包抄送
		for {
			// 初始化一个函数，
			// 用于同步数据到 Minecraft 客户端
			syncFunc := func() error {
				if shieldID := server.ShieldID.Load(); shieldID != 0 {
					client.ShieldID.Store(shieldID)
				}
				if entityUniqueID := server.GetEntityUniqueID(); entityUniqueID != 0 {
					client.SetEntityUniqueID(entityUniqueID)
				}
				if entityRuntimeID := server.GetEntityRuntimeID(); entityRuntimeID != 0 {
					client.SetEntityRuntimeID(entityRuntimeID)
				}
				if serverSkin := server.GetServerSkin(); serverSkin != nil {
					standardServerSkin := packet_translate_struct.ConvertToStandardSkin(*serverSkin)
					client.SetServerSkin(&standardServerSkin)
				}
				return nil
			}
			// 读取、过滤数据包，
			// 然后抄送其到 Minecraft 客户端
			errResults, syncError := server.FiltePacketsAndSendCopy(server.ReadPackets(), client.WritePackets, syncFunc)
			if syncError != nil {
				pterm.Warning.Printf("Eulogist: Failed to sync data when process packets from server, and the error log is %v", syncError)
			}
			for _, err = range errResults {
				if err != nil {
					pterm.Warning.Printf("Eulogist: Process packets from server error: %v\n", err)
				}
			}
			// 检查连接状态
			select {
			case <-server.Context.Done():
				return
			case <-client.Context.Done():
				return
			default:
			}
		}
	}()

	// 处理 Minecraft 客户端到赞颂者的数据包
	go func() {
		// 关闭已建立的所有连接
		defer func() {
			waitGroup.Add(-1)
			client.CloseConnection()
			server.CloseConnection()
		}()
		// 显示程序崩溃错误信息
		defer func() {
			r := recover()
			if r != nil {
				pterm.Error.Printf(
					"Eulogist/GoFunc/MinecraftClientToEulogist: err = %v\n\n[Stack Info]\n%s\n",
					r, string(debug.Stack()),
				)
				fmt.Println()
			}
		}()
		// 数据包抄送
		for {
			// 初始化一个函数，
			// 用于同步数据到网易租赁服
			syncFunc := func() error {
				return nil
			}
			// 读取、过滤数据包，
			// 然后抄送其到网易租赁服
			errResults, syncError := client.FiltePacketsAndSendCopy(client.ReadPackets(), server.WritePackets, syncFunc)
			if syncError != nil {
				pterm.Warning.Printf("Eulogist: Failed to sync data when process packets from client, and the error log is %v", syncError)
			}
			for _, err = range errResults {
				if err != nil {
					pterm.Warning.Printf("Eulogist: Process packets from client error: %v\n", err)
				}
			}
			// 检查连接状态
			select {
			case <-client.Context.Done():
				return
			case <-server.Context.Done():
				return
			default:
			}
		}
	}()

	// 等待所有 goroutine 完成
	waitGroup.Wait()
	pterm.Info.Println("Server Down. Now all connection was closed.")
	return nil
}
