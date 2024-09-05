package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type AddActor struct{}

func (pk *AddActor) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.AddActor{}
	input := standard.(*standardPacket.AddActor)

	p.EntityRuntimeID = input.EntityRuntimeID
	p.EntityUniqueID = input.EntityUniqueID
	p.EntityType = input.EntityType
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.HeadYaw = input.HeadYaw
	p.BodyYaw = input.BodyYaw
	p.EntityMetadata = input.EntityMetadata

	p.Attributes = packet_translate_struct.ConvertSlice(
		input.Attributes,
		func(from standardProtocol.AttributeValue) neteaseProtocol.AttributeValue {
			return neteaseProtocol.AttributeValue(from)
		},
	)
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
	p.EntityLinks = packet_translate_struct.ConvertSlice(
		input.EntityLinks,
		func(from standardProtocol.EntityLink) neteaseProtocol.EntityLink {
			return neteaseProtocol.EntityLink(from)
		},
	)

	p.Unknown1 = ""
	p.Unknown2 = ""
	p.Unknown3 = ""
	p.Unknown4 = false
	p.Unknown5 = false
	p.Unknown6 = false

	return &p
}

func (pk *AddActor) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.AddActor{}
	input := netease.(*neteasePacket.AddActor)

	p.EntityRuntimeID = input.EntityRuntimeID
	p.EntityUniqueID = input.EntityUniqueID
	p.EntityType = input.EntityType
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.HeadYaw = input.HeadYaw
	p.BodyYaw = input.BodyYaw
	p.EntityMetadata = input.EntityMetadata

	p.Attributes = packet_translate_struct.ConvertSlice(
		input.Attributes,
		func(from neteaseProtocol.AttributeValue) standardProtocol.AttributeValue {
			return standardProtocol.AttributeValue(from)
		},
	)
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
	p.EntityLinks = packet_translate_struct.ConvertSlice(
		input.EntityLinks,
		func(from neteaseProtocol.EntityLink) standardProtocol.EntityLink {
			return standardProtocol.EntityLink(from)
		},
	)

	// Something can not resolve. --Happy2018new
	for key := range p.EntityMetadata {
		switch key {
		case 0x0, 0x3, 0x4, 0x7, 0x8, 0x9, 0xf, 0x1a, 0x26, 0x2a, 0x35, 0x36, 0x51, 0x54, 0x5c:
		default:
			delete(p.EntityMetadata, key)
		}
	}

	return &p
}
