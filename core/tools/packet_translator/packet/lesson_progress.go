package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type LessonProgress struct{}

func (pk *LessonProgress) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.LessonProgress{}
	input := standard.(*standardPacket.LessonProgress)

	p.Identifier = input.Identifier
	p.Action = int32(input.Action)
	p.Score = input.Score

	return &p
}

func (pk *LessonProgress) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.LessonProgress{}
	input := netease.(*neteasePacket.LessonProgress)

	p.Identifier = input.Identifier
	p.Action = uint8(input.Action)
	p.Score = input.Score

	return &p
}
