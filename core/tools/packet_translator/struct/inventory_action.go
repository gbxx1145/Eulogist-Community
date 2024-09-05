package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
)

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
