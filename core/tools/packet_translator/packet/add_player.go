package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type AddPlayer struct{}

func (pk *AddPlayer) ToNetNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.AddPlayer{}
	input := standard.(*standardPacket.AddPlayer)

	p.UUID = input.UUID
	p.Username = input.Username
	p.EntityRuntimeID = input.EntityRuntimeID
	p.PlatformChatID = input.PlatformChatID
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.HeadYaw = input.HeadYaw
	p.GameType = input.GameType
	p.EntityMetadata = input.EntityMetadata
	p.DeviceID = input.DeviceID
	p.BuildPlatform = input.BuildPlatform
	p.HeldItem = packet_translator.ConvertToNetEaseItemInstance(input.HeldItem)

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
	p.AbilityData = neteaseProtocol.AbilityData{
		EntityUniqueID:     input.AbilityData.EntityUniqueID,
		PlayerPermissions:  input.AbilityData.PlayerPermissions,
		CommandPermissions: input.AbilityData.CommandPermissions,
		Layers: packet_translator.ConvertSlice(
			input.AbilityData.Layers,
			func(from standardProtocol.AbilityLayer) neteaseProtocol.AbilityLayer {
				return neteaseProtocol.AbilityLayer(from)
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
	p.Unknown3 = false
	p.Unknown4 = false

	return &p
}

func (pk *AddPlayer) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.AddPlayer{}
	input := netease.(*neteasePacket.AddPlayer)

	p.UUID = input.UUID
	p.Username = input.Username
	p.EntityRuntimeID = input.EntityRuntimeID
	p.PlatformChatID = input.PlatformChatID
	p.Position = input.Position
	p.Velocity = input.Velocity
	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.HeadYaw = input.HeadYaw
	p.GameType = input.GameType
	p.EntityMetadata = input.EntityMetadata
	p.DeviceID = input.DeviceID
	p.BuildPlatform = input.BuildPlatform
	p.HeldItem = packet_translator.ConvertToStandardItemInstance(input.HeldItem)

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
	p.AbilityData = standardProtocol.AbilityData{
		EntityUniqueID:     input.AbilityData.EntityUniqueID,
		PlayerPermissions:  input.AbilityData.PlayerPermissions,
		CommandPermissions: input.AbilityData.CommandPermissions,
		Layers: packet_translator.ConvertSlice(
			input.AbilityData.Layers,
			func(from neteaseProtocol.AbilityLayer) standardProtocol.AbilityLayer {
				return standardProtocol.AbilityLayer(from)
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
