package mc_server

import (
	"Eulogist/core/fb_auth/py_rpc"
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	"fmt"
)

// 数据包过滤器过滤来自租赁服的多个数据包，
// 然后并将过滤后的多个数据包抄送至客户端。
//
// 如果需要，
// 将根据实际情况由本处的桥接直接发送回应。
//
// writePacketToClient 指代
// 用于向客户端抄送数据包的函数。
//
// 返回的 []error 是一个列表，
// 分别对应 packets 中每一个数据包的处理成功情况
func (m *MinecraftServer) FiltePacketsAndSendCopy(
	packets []RaknetConnection.MinecraftPacket,
	writePacketsToClient func(packets []RaknetConnection.MinecraftPacket, useBytes bool),
) []error {
	// 初始化
	sendCopy := make([]RaknetConnection.MinecraftPacket, 0)
	shouldSendCopy := make([]bool, len(packets))
	errResults := make([]error, len(packets))
	// 处理每个数据包
	for index, minecraftPacket := range packets {
		// 如果传入的数据包为空，
		// 则直接返回 true 表示需要同步到客户端
		pk := minecraftPacket.Packet
		if pk == nil {
			shouldSendCopy[index] = true
			continue
		}
		// 根据数据包的类型进行不同的处理
		switch p := pk.(type) {
		case *packet.PyRpc:
			shouldSendCopy[index], errResults[index] = m.OnPyRpc(p)
			if err := errResults[index]; err != nil {
				errResults[index] = fmt.Errorf("FiltePacketsAndSendCopy: %v", err)
			}
		case *packet.StartGame:
			// 保存 entityUniqueID
			m.entityUniqueID = m.HandleStartGame(p)
			// 发送简要身份证明
			m.WriteSinglePacket(RaknetConnection.MinecraftPacket{
				Packet: &packet.NeteaseJson{
					Data: []byte(
						fmt.Sprintf(`{"eventName":"LOGIN_UID","resid":"","uid":"%s"}`,
							m.fbClient.ClientInfo.Uid,
						),
					),
				},
			}, false)
			// 皮肤特效处理
			playerSkin := m.GetPlayerSkin()
			if playerSkin == nil {
				shouldSendCopy[index] = true
				break
			}
			m.WriteSinglePacket(RaknetConnection.MinecraftPacket{
				Packet: &packet.PyRpc{
					Value: py_rpc.Marshal(&py_rpc.SyncUsingMod{
						[]any{},
						playerSkin.SkinUUID,
						playerSkin.SkinItemID,
						true,
						map[string]any{},
					}),
					OperationType: packet.PyRpcOperationTypeSend,
				},
			}, false)
			// 设置数据包抄送状态
			shouldSendCopy[index] = true
		case *packet.UpdatePlayerGameType:
			if p.PlayerUniqueID == m.entityUniqueID {
				// 如果玩家的唯一 ID 与数据包中记录的值匹配，
				// 则向客户端发送 SetPlayerGameType 数据包，
				// 并放弃当前数据包的发送，
				// 以确保 Minecraft 客户端可以正常同步游戏模式更改。
				// 否则，按原样抄送当前数据包
				writePacketsToClient([]RaknetConnection.MinecraftPacket{
					{Packet: &packet.SetPlayerGameType{GameType: p.GameType}},
				}, false)
			}
			// 设置数据包抄送状态
			shouldSendCopy[index] = p.PlayerUniqueID != m.entityUniqueID
		default:
			// 默认情况下，
			// 我们需要将数据包同步到客户端
			shouldSendCopy[index] = true
		}
	}
	// 抄送数据包
	for index, pk := range packets {
		if !shouldSendCopy[index] {
			continue
		}
		sendCopy = append(sendCopy, pk)
	}
	writePacketsToClient(sendCopy, true)
	// 返回值
	return errResults
}
