package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type InventorySlot struct{}

func (pk *InventorySlot) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventorySlot{}
	input := standard.(*standardPacket.InventorySlot)

	p.WindowID = input.WindowID
	p.Slot = input.Slot
	p.NewItem = packet_translator.ConvertToNetEaseItemInstance(input.NewItem)

	return &p
}

func (pk *InventorySlot) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventorySlot{}
	input := netease.(*neteasePacket.InventorySlot)

	p.WindowID = input.WindowID
	p.Slot = input.Slot
	p.NewItem = packet_translator.ConvertToStandardItemInstance(input.NewItem)

	return &p
}
