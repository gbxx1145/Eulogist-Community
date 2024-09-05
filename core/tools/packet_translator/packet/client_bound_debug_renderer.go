package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type ClientBoundDebugRenderer struct{}

func (pk *ClientBoundDebugRenderer) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	input := standard.(*standardPacket.ClientBoundDebugRenderer)
	p := neteasePacket.ClientBoundDebugRenderer(*input)

	return &p
}

func (pk *ClientBoundDebugRenderer) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	input := netease.(*neteasePacket.ClientBoundDebugRenderer)
	p := standardPacket.ClientBoundDebugRenderer(*input)

	return &p
}
