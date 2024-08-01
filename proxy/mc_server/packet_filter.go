package mc_server

import (
	"Eulogist/core/minecraft/protocol/packet"
	"fmt"
)

// 数据包过滤器过滤来自租赁服的数据包，
// 并根据实际情况由本处的桥接选择是否直接发送回应。
//
// shouldSendCopy 指代该数据包是否需要同步到 Client
func (m *MinecraftServer) PacketFilter(pk packet.Packet) (shouldSendCopy bool, err error) {
	if pk == nil {
		return true, nil
	}

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
	default:
		return true, nil
	}
}
