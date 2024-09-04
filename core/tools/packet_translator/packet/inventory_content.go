package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "Eulogist/core/standard/protocol"
	standardPacket "Eulogist/core/standard/protocol/packet"
)

type InventoryContent struct{}

func (pk *InventoryContent) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.InventoryContent{}
	input := standard.(*standardPacket.InventoryContent)

	p.WindowID = input.WindowID
	p.Content = ConvertSlice(
		input.Content,
		func(from standardProtocol.ItemInstance) neteaseProtocol.ItemInstance {
			return ConvertToNetEaseItemInstance(from)
		},
	)

	return &p
}

func (pk *InventoryContent) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.InventoryContent{}
	input := netease.(*neteasePacket.InventoryContent)

	p.WindowID = input.WindowID
	p.Content = ConvertSlice(
		input.Content,
		func(from neteaseProtocol.ItemInstance) standardProtocol.ItemInstance {
			return ConvertToStandardItemInstance(from)
		},
	)

	return &p
}
