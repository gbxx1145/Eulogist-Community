package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type Animate struct{}

func (pk *Animate) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.Animate{}
	input := standard.(*standardPacket.Animate)

	p.ActionType = input.ActionType
	p.EntityRuntimeID = input.EntityRuntimeID
	p.BoatRowingTime = input.BoatRowingTime

	p.AttackerEntityUniqueID = 0

	return &p
}

func (pk *Animate) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.Animate{}
	input := netease.(*neteasePacket.Animate)

	p.ActionType = input.ActionType
	p.EntityRuntimeID = input.EntityRuntimeID
	p.BoatRowingTime = input.BoatRowingTime

	return &p
}
