package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type PlayerSkin struct{}

func (pk *PlayerSkin) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.PlayerSkin{}
	input := standard.(*standardPacket.PlayerSkin)

	p.UUID = input.UUID
	p.NewSkinName = input.NewSkinName
	p.OldSkinName = input.OldSkinName
	p.Skin = ConvertToNetEaseSkin(input.Skin)

	return &p
}

func (pk *PlayerSkin) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.PlayerSkin{}
	input := netease.(*neteasePacket.PlayerSkin)

	p.UUID = input.UUID
	p.NewSkinName = input.NewSkinName
	p.OldSkinName = input.OldSkinName
	p.Skin = ConvertToStandardSkin(input.Skin)

	return &p
}
