package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
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
