package BotSide

import (
	"Eulogist/minecraft/protocol/packet"
	"fmt"
)

// 数据包过滤器过滤来自租赁服的数据包，
// 并根据实际情况由本处的桥接选择是否直接发送回应。
//
// shouldSendCopy 指代该数据包是否需要同步到 ModPC
func (b *BotSide) PacketFilter(pk packet.Packet) (shouldSendCopy bool, err error) {
	switch p := pk.(type) {
	case *packet.PyRpc:
		shouldSendCopy, err = b.OnPyRpc(p)
		if err != nil {
			return shouldSendCopy, fmt.Errorf("PacketFilter: %v", err)
		}
	case *packet.StartGame:
		b.gameData, err = b.HandleStartGame(p)
		if err != nil {
			return true, fmt.Errorf("PacketFilter: %v", err)
		}
		shouldSendCopy = true
	default:
		shouldSendCopy = true
	}

	return shouldSendCopy, nil
}
