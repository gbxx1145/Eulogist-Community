package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type AddEntity struct{}

func (pk *AddEntity) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.AddEntity{}
	input := standard.(*standardPacket.AddEntity)

	p.EntityNetworkID = uint32(input.EntityNetworkID)

	return &p
}

func (pk *AddEntity) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.AddEntity{}
	input := netease.(*neteasePacket.AddEntity)

	p.EntityNetworkID = uint64(input.EntityNetworkID)

	return &p
}
