package mc_client

import (
	"Eulogist/core/fb_auth/py_rpc"
	"Eulogist/core/minecraft/protocol/packet"
	"fmt"
)

// OnPyRpc 处理数据包 PyRpc
func (m *MinecraftClient) OnPyRpc(p *packet.PyRpc) (shouldSendCopy bool, err error) {
	// 解码 PyRpc
	if p.Value == nil {
		return true, nil
	}
	content, err := py_rpc.Unmarshal(p.Value)
	if err != nil {
		return true, fmt.Errorf("OnPyRpc: %v", err)
	}
	// 根据内容类型处理不同的 PyRpc
	switch content.(type) {
	case *py_rpc.SyncUsingMod:
	default:
		// 对于其他种类的 PyRpc 数据包，
		// 返回 true 表示需要将数据包抄
		// 送至网易租赁服
		return true, nil
	}
	// 返回值
	return false, nil
}
