package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type UpdateBlockSynced struct{}

func (pk *UpdateBlockSynced) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.UpdateBlockSynced{}
	input := standard.(*standardPacket.UpdateBlockSynced)

	p.Position = neteaseProtocol.BlockPos(input.Position)
	p.NewBlockRuntimeID = input.NewBlockRuntimeID
	p.Flags = input.Flags
	p.Layer = input.Layer
	p.EntityUniqueID = uint64(input.EntityUniqueID)
	p.TransitionType = input.TransitionType

	return &p
}

func (pk *UpdateBlockSynced) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.UpdateBlockSynced{}
	input := netease.(*neteasePacket.UpdateBlockSynced)

	p.Position = standardProtocol.BlockPos(input.Position)
	p.NewBlockRuntimeID = input.NewBlockRuntimeID
	p.Flags = input.Flags
	p.Layer = input.Layer
	p.EntityUniqueID = int64(input.EntityUniqueID)
	p.TransitionType = input.TransitionType

	return &p
}
