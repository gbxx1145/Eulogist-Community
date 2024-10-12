package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type AddItemActor struct{}

func (pk *AddItemActor) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.AddItemActor{}
	input := standard.(*standardPacket.AddItemActor)

	p.EntityUniqueID = input.EntityUniqueID
	p.EntityRuntimeID = input.EntityRuntimeID
	p.Item = packet_translate_struct.ConvertToNetEaseItemInstance(input.Item)
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.FromFishing = input.FromFishing

	p.EntityMetadata = make(map[uint32]any)
	for key, value := range input.EntityMetadata {
		if v, ok := value.(standardProtocol.BlockPos); ok {
			p.EntityMetadata[key] = neteaseProtocol.BlockPos(v)
		} else {
			p.EntityMetadata[key] = value
		}
	}

	return &p
}

func (pk *AddItemActor) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.AddItemActor{}
	input := netease.(*neteasePacket.AddItemActor)

	p.EntityUniqueID = input.EntityUniqueID
	p.EntityRuntimeID = input.EntityRuntimeID
	p.Item = packet_translate_struct.ConvertToStandardItemInstance(input.Item)
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.FromFishing = input.FromFishing

	p.EntityMetadata = make(map[uint32]any)
	for key, value := range input.EntityMetadata {
		if v, ok := value.(neteaseProtocol.BlockPos); ok {
			p.EntityMetadata[key] = standardProtocol.BlockPos(v)
		} else {
			p.EntityMetadata[key] = value
		}
	}

	return &p
}
