package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type InventoryTransaction struct{}

func (pk *InventoryTransaction) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventoryTransaction{}
	input := standard.(*standardPacket.InventoryTransaction)

	p.LegacyRequestID = input.LegacyRequestID

	p.LegacySetItemSlots = packet_translate_struct.ConvertSlice(
		input.LegacySetItemSlots,
		func(from standardProtocol.LegacySetItemSlot) neteaseProtocol.LegacySetItemSlot {
			return neteaseProtocol.LegacySetItemSlot(from)
		},
	)
	p.Actions = packet_translate_struct.ConvertSlice(
		input.Actions,
		packet_translate_struct.ConvertToNetEaseInventoryAction,
	)

	switch data := input.TransactionData.(type) {
	case *standardProtocol.NormalTransactionData:
		p.TransactionData = &neteaseProtocol.NormalTransactionData{}
	case *standardProtocol.MismatchTransactionData:
		p.TransactionData = &neteaseProtocol.MismatchTransactionData{}
	case *standardProtocol.UseItemTransactionData:
		p.TransactionData = packet_translate_struct.ConvertToNetEaseUseItemTransactionData(data)
	case *standardProtocol.UseItemOnEntityTransactionData:
		p.TransactionData = &neteaseProtocol.UseItemOnEntityTransactionData{
			TargetEntityRuntimeID: data.TargetEntityRuntimeID,
			ActionType:            data.ActionType,
			HotBarSlot:            data.HotBarSlot,
			HeldItem:              packet_translate_struct.ConvertToNetEaseItemInstance(data.HeldItem),
			Position:              data.Position,
			ClickedPosition:       data.ClickedPosition,
		}
	case *standardProtocol.ReleaseItemTransactionData:
		p.TransactionData = &neteaseProtocol.ReleaseItemTransactionData{
			ActionType:   data.ActionType,
			HotBarSlot:   data.HotBarSlot,
			HeldItem:     packet_translate_struct.ConvertToNetEaseItemInstance(data.HeldItem),
			HeadPosition: data.HeadPosition,
		}
	}

	return &p
}

func (pk *InventoryTransaction) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventoryTransaction{}
	input := netease.(*neteasePacket.InventoryTransaction)

	p.LegacyRequestID = input.LegacyRequestID

	p.LegacySetItemSlots = packet_translate_struct.ConvertSlice(
		input.LegacySetItemSlots,
		func(from neteaseProtocol.LegacySetItemSlot) standardProtocol.LegacySetItemSlot {
			return standardProtocol.LegacySetItemSlot(from)
		},
	)
	p.Actions = packet_translate_struct.ConvertSlice(
		input.Actions,
		packet_translate_struct.ConvertToStandardInventoryAction,
	)

	switch data := input.TransactionData.(type) {
	case *neteaseProtocol.NormalTransactionData:
		p.TransactionData = &standardProtocol.NormalTransactionData{}
	case *neteaseProtocol.MismatchTransactionData:
		p.TransactionData = &standardProtocol.MismatchTransactionData{}
	case *neteaseProtocol.UseItemTransactionData:
		p.TransactionData = packet_translate_struct.ConvertToStandardUseItemTransactionData(data)
	case *neteaseProtocol.UseItemOnEntityTransactionData:
		p.TransactionData = &standardProtocol.UseItemOnEntityTransactionData{
			TargetEntityRuntimeID: data.TargetEntityRuntimeID,
			ActionType:            data.ActionType,
			HotBarSlot:            data.HotBarSlot,
			HeldItem:              packet_translate_struct.ConvertToStandardItemInstance(data.HeldItem),
			Position:              data.Position,
			ClickedPosition:       data.ClickedPosition,
		}
	case *neteaseProtocol.ReleaseItemTransactionData:
		p.TransactionData = &standardProtocol.ReleaseItemTransactionData{
			ActionType:   data.ActionType,
			HotBarSlot:   data.HotBarSlot,
			HeldItem:     packet_translate_struct.ConvertToStandardItemInstance(data.HeldItem),
			HeadPosition: data.HeadPosition,
		}
	}

	return &p
}
