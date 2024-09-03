package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PlayerList struct{}

func (pk *PlayerList) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.PlayerList{}
	input := standard.(*standardPacket.PlayerList)

	p.ActionType = input.ActionType

	p.Entries = packet_translator.ConvertSlice(
		input.Entries,
		func(from standardProtocol.PlayerListEntry) neteaseProtocol.PlayerListEntry {
			return neteaseProtocol.PlayerListEntry{
				UUID:           from.UUID,
				EntityUniqueID: from.EntityUniqueID,
				Username:       from.Username,
				XUID:           from.XUID,
				PlatformChatID: from.PlatformChatID,
				BuildPlatform:  from.BuildPlatform,
				Skin:           packet_translator.ConvertToNetEaseSkin(from.Skin),
			}
		},
	)

	p.Unknown1 = make([]neteaseProtocol.NeteaseUnknownPlayerListEntry, 0)
	p.Unknown2 = make([]neteaseProtocol.NeteaseUnknownPlayerListEntry, 0)
	p.Unknown3 = make([]string, 0)
	p.Unknown4 = make([]string, 0)
	p.GrowthLevels = make([]uint32, 0)

	return &p
}

func (pk *PlayerList) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.PlayerList{}
	input := netease.(*neteasePacket.PlayerList)

	p.ActionType = input.ActionType

	p.Entries = packet_translator.ConvertSlice(
		input.Entries,
		func(from neteaseProtocol.PlayerListEntry) standardProtocol.PlayerListEntry {
			return standardProtocol.PlayerListEntry{
				UUID:           from.UUID,
				EntityUniqueID: from.EntityUniqueID,
				Username:       from.Username,
				XUID:           from.XUID,
				PlatformChatID: from.PlatformChatID,
				BuildPlatform:  from.BuildPlatform,
				Skin:           packet_translator.ConvertToStandardSkin(from.Skin),
			}
		},
	)

	return &p
}
