package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type RemoveVolumeEntity struct{}

func (pk *RemoveVolumeEntity) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.RemoveVolumeEntity{}
	input := standard.(*standardPacket.RemoveVolumeEntity)

	p.EntityRuntimeID = uint32(input.EntityRuntimeID)
	p.Dimension = input.Dimension

	return &p
}

func (pk *RemoveVolumeEntity) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.RemoveVolumeEntity{}
	input := netease.(*neteasePacket.RemoveVolumeEntity)

	p.EntityRuntimeID = uint64(input.EntityRuntimeID)
	p.Dimension = input.Dimension

	return &p
}
