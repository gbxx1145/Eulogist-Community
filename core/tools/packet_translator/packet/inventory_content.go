package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type InventoryContent struct{}

func (pk *InventoryContent) ToNetNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventoryContent{}
	input := standard.(*standardPacket.InventoryContent)

	p.WindowID = input.WindowID
	p.Content = packet_translator.ConvertSlice(
		input.Content,
		func(from standardProtocol.ItemInstance) neteaseProtocol.ItemInstance {
			return packet_translator.ConvertToNetEaseItemInstance(from)
		},
	)

	return &p
}

func (pk *InventoryContent) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventoryContent{}
	input := netease.(*neteasePacket.InventoryContent)

	p.WindowID = input.WindowID
	p.Content = packet_translator.ConvertSlice(
		input.Content,
		func(from neteaseProtocol.ItemInstance) standardProtocol.ItemInstance {
			return packet_translator.ConvertToStandardItemInstance(from)
		},
	)

	return &p
}
