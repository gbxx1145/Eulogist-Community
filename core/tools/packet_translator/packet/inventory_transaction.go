package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type InventoryTransaction struct{}

func (pk *InventoryTransaction) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventoryTransaction{}
	input := standard.(*standardPacket.InventoryTransaction)

	p.LegacyRequestID = input.LegacyRequestID

	p.LegacySetItemSlots = ConvertSlice(
		input.LegacySetItemSlots,
		func(from standardProtocol.LegacySetItemSlot) neteaseProtocol.LegacySetItemSlot {
			return neteaseProtocol.LegacySetItemSlot(from)
		},
	)
	p.Actions = ConvertSlice(
		input.Actions,
		ConvertToNetEaseInventoryAction,
	)

	switch data := input.TransactionData.(type) {
	case *standardProtocol.NormalTransactionData:
		p.TransactionData = &neteaseProtocol.NormalTransactionData{}
	case *standardProtocol.MismatchTransactionData:
		p.TransactionData = &neteaseProtocol.MismatchTransactionData{}
	case *standardProtocol.UseItemTransactionData:
		p.TransactionData = ConvertToNetEaseUseItemTransactionData(data)
	case *standardProtocol.UseItemOnEntityTransactionData:
		p.TransactionData = &neteaseProtocol.UseItemOnEntityTransactionData{
			TargetEntityRuntimeID: data.TargetEntityRuntimeID,
			ActionType:            data.ActionType,
			HotBarSlot:            data.HotBarSlot,
			HeldItem:              ConvertToNetEaseItemInstance(data.HeldItem),
			Position:              data.Position,
			ClickedPosition:       data.ClickedPosition,
		}
	case *standardProtocol.ReleaseItemTransactionData:
		p.TransactionData = &neteaseProtocol.ReleaseItemTransactionData{
			ActionType:   data.ActionType,
			HotBarSlot:   data.HotBarSlot,
			HeldItem:     ConvertToNetEaseItemInstance(data.HeldItem),
			HeadPosition: data.HeadPosition,
		}
	}

	return &p
}

func (pk *InventoryTransaction) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventoryTransaction{}
	input := netease.(*neteasePacket.InventoryTransaction)

	p.LegacyRequestID = input.LegacyRequestID

	p.LegacySetItemSlots = ConvertSlice(
		input.LegacySetItemSlots,
		func(from neteaseProtocol.LegacySetItemSlot) standardProtocol.LegacySetItemSlot {
			return standardProtocol.LegacySetItemSlot(from)
		},
	)
	p.Actions = ConvertSlice(
		input.Actions,
		ConvertToStandardInventoryAction,
	)

	switch data := input.TransactionData.(type) {
	case *neteaseProtocol.NormalTransactionData:
		p.TransactionData = &standardProtocol.NormalTransactionData{}
	case *neteaseProtocol.MismatchTransactionData:
		p.TransactionData = &standardProtocol.MismatchTransactionData{}
	case *neteaseProtocol.UseItemTransactionData:
		p.TransactionData = ConvertToStandardUseItemTransactionData(data)
	case *neteaseProtocol.UseItemOnEntityTransactionData:
		p.TransactionData = &standardProtocol.UseItemOnEntityTransactionData{
			TargetEntityRuntimeID: data.TargetEntityRuntimeID,
			ActionType:            data.ActionType,
			HotBarSlot:            data.HotBarSlot,
			HeldItem:              ConvertToStandardItemInstance(data.HeldItem),
			Position:              data.Position,
			ClickedPosition:       data.ClickedPosition,
		}
	case *neteaseProtocol.ReleaseItemTransactionData:
		p.TransactionData = &standardProtocol.ReleaseItemTransactionData{
			ActionType:   data.ActionType,
			HotBarSlot:   data.HotBarSlot,
			HeldItem:     ConvertToStandardItemInstance(data.HeldItem),
			HeadPosition: data.HeadPosition,
		}
	}

	return &p
}
