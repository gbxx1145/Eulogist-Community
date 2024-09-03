package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type AddActor struct{}

func (pk *AddActor) ToNetNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
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

	p.Attributes = packet_translator.ConvertSlice(
		input.Attributes,
		func(from standardProtocol.AttributeValue) neteaseProtocol.AttributeValue {
			return neteaseProtocol.AttributeValue(from)
		},
	)
	p.EntityProperties = neteaseProtocol.EntityProperties{
		IntegerProperties: packet_translator.ConvertSlice(
			input.EntityProperties.IntegerProperties,
			func(from standardProtocol.IntegerEntityProperty) neteaseProtocol.IntegerEntityProperty {
				return neteaseProtocol.IntegerEntityProperty(from)
			},
		),
		FloatProperties: packet_translator.ConvertSlice(
			input.EntityProperties.FloatProperties,
			func(from standardProtocol.FloatEntityProperty) neteaseProtocol.FloatEntityProperty {
				return neteaseProtocol.FloatEntityProperty(from)
			},
		),
	}
	p.EntityLinks = packet_translator.ConvertSlice(
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

	p.Attributes = packet_translator.ConvertSlice(
		input.Attributes,
		func(from neteaseProtocol.AttributeValue) standardProtocol.AttributeValue {
			return standardProtocol.AttributeValue(from)
		},
	)
	p.EntityProperties = standardProtocol.EntityProperties{
		IntegerProperties: packet_translator.ConvertSlice(
			input.EntityProperties.IntegerProperties,
			func(from neteaseProtocol.IntegerEntityProperty) standardProtocol.IntegerEntityProperty {
				return standardProtocol.IntegerEntityProperty(from)
			},
		),
		FloatProperties: packet_translator.ConvertSlice(
			input.EntityProperties.FloatProperties,
			func(from neteaseProtocol.FloatEntityProperty) standardProtocol.FloatEntityProperty {
				return standardProtocol.FloatEntityProperty(from)
			},
		),
	}
	p.EntityLinks = packet_translator.ConvertSlice(
		input.EntityLinks,
		func(from neteaseProtocol.EntityLink) standardProtocol.EntityLink {
			return standardProtocol.EntityLink(from)
		},
	)

	return &p
}
