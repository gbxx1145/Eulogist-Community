package mc_client

import (
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	"fmt"
)

// 数据包过滤器过滤来自 Minecraft 客户端的数据包，
// 并根据实际情况由本处的桥接选择是否需要修改。
//
// 如果必要，将使用 writePacketToServer
// 向租赁服发送新数据包。
//
// 返回的 shouldSendCopy
// 指代该数据包是否需要同步到网易租赁服
func (m *MinecraftClient) PacketFilter(
	pk packet.Packet, writePacketToServer func(pk RaknetConnection.MinecraftPacket, useBytes bool) error,
) (shouldSendCopy bool, err error) {
	// 如果传入的数据包为空，
	// 则直接返回 true
	// 表示需要同步到网易租赁服
	if pk == nil {
		return true, nil
	}

	// 根据数据包的类型进行不同的处理
	switch p := pk.(type) {
	case *packet.PyRpc:
		shouldSendCopy, err = m.OnPyRpc(p, writePacketToServer)
		if err != nil {
			err = fmt.Errorf("PacketFilter: %v", err)
		}
		return shouldSendCopy, err
	}

	// 默认情况下，返回 true
	// 表示需要同步到网易租赁服
	return true, nil
}
