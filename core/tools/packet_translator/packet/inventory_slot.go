package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type InventorySlot struct{}

func (pk *InventorySlot) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventorySlot{}
	input := standard.(*standardPacket.InventorySlot)

	p.WindowID = input.WindowID
	p.Slot = input.Slot
	p.NewItem = ConvertToNetEaseItemInstance(input.NewItem)

	return &p
}

func (pk *InventorySlot) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventorySlot{}
	input := netease.(*neteasePacket.InventorySlot)

	p.WindowID = input.WindowID
	p.Slot = input.Slot
	p.NewItem = ConvertToStandardItemInstance(input.NewItem)

	return &p
}
