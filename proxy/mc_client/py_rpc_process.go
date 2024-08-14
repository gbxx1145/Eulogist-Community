package mc_client

import (
	"Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/py_rpc"
	"fmt"
)

// OnPyRpc 处理数据包 PyRpc
func (m *MinecraftClient) OnPyRpc(p *packet.PyRpc) (doNotSendCopy bool, err error) {
	// 解码 PyRpc
	if p.Value == nil {
		return false, nil
	}
	content, err := py_rpc.Unmarshal(p.Value)
	if err != nil {
		return false, fmt.Errorf("OnPyRpc: %v", err)
	}
	// 根据内容类型处理不同的 PyRpc
	switch content.(type) {
	case *py_rpc.SyncUsingMod:
	default:
		// 对于其他种类的 PyRpc 数据包，
		// 返回 false 表示需要将数据包抄
		// 送至网易租赁服
		return false, nil
	}
	// 返回值
	return true, nil
}
