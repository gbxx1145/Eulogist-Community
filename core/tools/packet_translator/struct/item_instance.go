package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
)

// 将 from 转换为 neteaseProtocol.ItemInstance
func ConvertToNetEaseItemInstance(from standardProtocol.ItemInstance) neteaseProtocol.ItemInstance {
	return neteaseProtocol.ItemInstance{
		StackNetworkID: from.StackNetworkID,
		Stack:          ConvertToNetEaseItemStack(from.Stack),
	}
}

// 将 from 转换为 standardProtocol.ItemInstance
func ConvertToStandardItemInstance(from neteaseProtocol.ItemInstance) standardProtocol.ItemInstance {
	return standardProtocol.ItemInstance{
		StackNetworkID: from.StackNetworkID,
		Stack:          ConvertToStandardItemStack(from.Stack),
	}
}
