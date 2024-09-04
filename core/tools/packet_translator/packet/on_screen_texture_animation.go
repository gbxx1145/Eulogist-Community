package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type OnScreenTextureAnimation struct{}

func (pk *OnScreenTextureAnimation) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.OnScreenTextureAnimation{}
	input := standard.(*standardPacket.OnScreenTextureAnimation)

	p.AnimationType = uint32(input.AnimationType)

	return &p
}

func (pk *OnScreenTextureAnimation) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.OnScreenTextureAnimation{}
	input := netease.(*neteasePacket.OnScreenTextureAnimation)

	p.AnimationType = int32(input.AnimationType)

	return &p
}
