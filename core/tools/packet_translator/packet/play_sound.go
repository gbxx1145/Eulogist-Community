package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type PlaySound struct{}

func (pk *PlaySound) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	input := standard.(*standardPacket.PlaySound)
	p := neteasePacket.PlaySound(*input)

	return &p
}

func (pk *PlaySound) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	input := netease.(*neteasePacket.PlaySound)
	p := standardPacket.PlaySound(*input)

	return &p
}
