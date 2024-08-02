package mc_server

import (
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	"fmt"
)

// 数据包过滤器过滤来自租赁服的数据包，
// 并根据实际情况由本处的桥接选择是否直接发送回应。
//
// 如果必要，将使用 writePacketToClient 向已连接的
// Minecraft 客户端发送新数据包。
//
// 返回的 shouldSendCopy 指代该数据包是否需要同步到
// Minecraft 客户端
func (m *MinecraftServer) PacketFilter(
	pk packet.Packet, writePacketToClient func(pk RaknetConnection.MinecraftPacket, useBytes bool) error,
) (shouldSendCopy bool, err error) {
	// 如果传入的数据包为空，
	// 则直接返回 true 表示需要同步到客户端
	if pk == nil {
		return true, nil
	}

	// 根据数据包的类型进行不同的处理
	switch p := pk.(type) {
	case *packet.PyRpc:
		shouldSendCopy, err = m.OnPyRpc(p)
		if err != nil {
			err = fmt.Errorf("PacketFilter: %v", err)
		}
		return shouldSendCopy, err
	case *packet.StartGame:
		m.entityUniqueID = m.HandleStartGame(p)
		m.SetShouldDecode(false)
		return true, nil
	case *packet.UpdatePlayerGameType:
		if p.PlayerUniqueID == m.entityUniqueID {
			// 如果玩家的唯一 ID 与数据包中记录的值匹配，
			// 则向客户端发送 SetPlayerGameType 数据包，
			// 并放弃当前数据包的发送，
			// 以确保 Minecraft 客户端可以正常同步游戏模式更改。
			// 否则，按原样抄送当前数据包
			err = writePacketToClient(RaknetConnection.MinecraftPacket{
				Packet: &packet.SetPlayerGameType{GameType: p.GameType},
			}, false)
			if err != nil {
				err = fmt.Errorf("PacketFilter: %v", err)
			}
		}
		// 返回是否需要同步到客户端
		return p.PlayerUniqueID != m.entityUniqueID, err
	}

	// 默认情况下，返回 true 表示需要同步到客户端
	return true, nil
}
