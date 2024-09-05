package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type ChangeMobProperty struct{}

func (pk *ChangeMobProperty) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.ChangeMobProperty{}
	input := standard.(*standardPacket.ChangeMobProperty)

	p.EntityUniqueID = int64(input.EntityUniqueID)
	p.Property = input.Property
	p.BoolValue = input.BoolValue
	p.StringValue = input.StringValue
	p.IntValue = input.IntValue
	p.FloatValue = input.FloatValue

	return &p
}

func (pk *ChangeMobProperty) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.ChangeMobProperty{}
	input := netease.(*neteasePacket.ChangeMobProperty)

	p.EntityUniqueID = uint64(input.EntityUniqueID)
	p.Property = input.Property
	p.BoolValue = input.BoolValue
	p.StringValue = input.StringValue
	p.IntValue = input.IntValue
	p.FloatValue = input.FloatValue

	return &p
}
