package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type SetActorData struct{}

func (pk *SetActorData) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.SetActorData{}
	input := standard.(*standardPacket.SetActorData)

	p.EntityRuntimeID = input.EntityRuntimeID
	p.Tick = input.Tick

	p.EntityProperties = neteaseProtocol.EntityProperties{
		IntegerProperties: packet_translate_struct.ConvertSlice(
			input.EntityProperties.IntegerProperties,
			func(from standardProtocol.IntegerEntityProperty) neteaseProtocol.IntegerEntityProperty {
				return neteaseProtocol.IntegerEntityProperty(from)
			},
		),
		FloatProperties: packet_translate_struct.ConvertSlice(
			input.EntityProperties.FloatProperties,
			func(from standardProtocol.FloatEntityProperty) neteaseProtocol.FloatEntityProperty {
				return neteaseProtocol.FloatEntityProperty(from)
			},
		),
	}

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

func (pk *SetActorData) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.SetActorData{}
	input := netease.(*neteasePacket.SetActorData)

	p.EntityRuntimeID = input.EntityRuntimeID
	p.Tick = input.Tick

	p.EntityProperties = standardProtocol.EntityProperties{
		IntegerProperties: packet_translate_struct.ConvertSlice(
			input.EntityProperties.IntegerProperties,
			func(from neteaseProtocol.IntegerEntityProperty) standardProtocol.IntegerEntityProperty {
				return standardProtocol.IntegerEntityProperty(from)
			},
		),
		FloatProperties: packet_translate_struct.ConvertSlice(
			input.EntityProperties.FloatProperties,
			func(from neteaseProtocol.FloatEntityProperty) standardProtocol.FloatEntityProperty {
				return standardProtocol.FloatEntityProperty(from)
			},
		),
	}

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
