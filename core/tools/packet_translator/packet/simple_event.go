package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type SimpleEvent struct{}

func (pk *SimpleEvent) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.SimpleEvent{}
	input := standard.(*standardPacket.SimpleEvent)

	p.EventType = uint16(input.EventType)

	return &p
}

func (pk *SimpleEvent) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.SimpleEvent{}
	input := netease.(*neteasePacket.SimpleEvent)

	p.EventType = int16(input.EventType)

	return &p
}
