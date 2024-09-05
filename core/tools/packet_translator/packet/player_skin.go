package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type PlayerSkin struct{}

func (pk *PlayerSkin) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.PlayerSkin{}
	input := standard.(*standardPacket.PlayerSkin)

	p.UUID = input.UUID
	p.NewSkinName = input.NewSkinName
	p.OldSkinName = input.OldSkinName
	p.Skin = packet_translate_struct.ConvertToNetEaseSkin(input.Skin)

	return &p
}

func (pk *PlayerSkin) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.PlayerSkin{}
	input := netease.(*neteasePacket.PlayerSkin)

	p.UUID = input.UUID
	p.NewSkinName = input.NewSkinName
	p.OldSkinName = input.OldSkinName
	p.Skin = packet_translate_struct.ConvertToStandardSkin(input.Skin)

	return &p
}
