package mc_client

import (
	"Eulogist/core/fb_auth/py_rpc"
	"Eulogist/core/minecraft/protocol/packet"
	RaknetConnection "Eulogist/core/raknet"
	"fmt"
)

// OnPyRpc 处理数据包 PyRpc。
//
// 如果必要，将使用 writePacketToServer
// 向网易租赁服发送新数据包。
func (m *MinecraftClient) OnPyRpc(
	p *packet.PyRpc,
	writePacketToServer func(pk RaknetConnection.MinecraftPacket, useBytes bool) error,
) (shouldSendCopy bool, err error) {
	// 解码 PyRpc
	if p.Value == nil {
		return true, nil
	}
	content, err := py_rpc.Unmarshal(p.Value)
	if err != nil {
		return true, fmt.Errorf("OnPyRpc: %v", err)
	}
	// 根据内容类型处理不同的 PyRpc
	switch c := content.(type) {
	case *py_rpc.SyncUsingMod:
		err = c.FromGo([]any{
			[]any{},
			m.playerSkin.SkinUUID,
			m.playerSkin.SkinItemID,
			true,
			map[string]any{},
		})
		if err != nil {
			return false, fmt.Errorf("OnPyRpc: %v", err)
		}
		err = writePacketToServer(
			RaknetConnection.MinecraftPacket{
				Packet: &packet.PyRpc{
					Value:         py_rpc.Marshal(c),
					OperationType: packet.PyRpcOperationTypeSend,
				},
			}, false,
		)
		if err != nil {
			return false, fmt.Errorf("OnPyRpc: %v", err)
		}
	default:
		// 对于其他种类的 PyRpc 数据包，
		// 返回 true 表示需要将数据包抄送至
		// 网易租赁服
		return true, nil
	}
	// 返回值
	return false, nil
}
