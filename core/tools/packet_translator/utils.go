package packet_translator

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
)

// 将切片 from([]From) 转换为 []To。
// converter 是用于转换切片内的单个元素的函数
func ConvertSlice[From any, To any](
	from []From,
	converter func(from From) To,
) []To {
	to := make([]To, len(from))
	for index, value := range from {
		to[index] = converter(value)
	}
	return to
}

// 将 from 转换为 neteaseProtocol.ItemStack
func ConvertToNetEaseItemStack(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
	return neteaseProtocol.ItemStack{
		ItemType:       neteaseProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}
}

// 将 from 转换为 standardProtocol.ItemStack
func ConvertToStandardItemStack(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
	return standardProtocol.ItemStack{
		ItemType:       standardProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}
}

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

// 将 from 转换为 neteaseProtocol.InventoryAction
func ConvertToNetEaseInventoryAction(from standardProtocol.InventoryAction) neteaseProtocol.InventoryAction {
	return neteaseProtocol.InventoryAction{
		SourceType:    from.SourceType,
		WindowID:      from.WindowID,
		SourceFlags:   from.SourceFlags,
		InventorySlot: from.InventorySlot,
		OldItem:       ConvertToNetEaseItemInstance(from.OldItem),
		NewItem:       ConvertToNetEaseItemInstance(from.NewItem),
	}
}

// 将 from 转换为 standardProtocol.InventoryAction
func ConvertToStandardInventoryAction(from neteaseProtocol.InventoryAction) standardProtocol.InventoryAction {
	return standardProtocol.InventoryAction{
		SourceType:    from.SourceType,
		WindowID:      from.WindowID,
		SourceFlags:   from.SourceFlags,
		InventorySlot: from.InventorySlot,
		OldItem:       ConvertToStandardItemInstance(from.OldItem),
		NewItem:       ConvertToStandardItemInstance(from.NewItem),
	}
}
