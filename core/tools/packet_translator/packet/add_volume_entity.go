package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "Eulogist/core/standard/protocol"
	standardPacket "Eulogist/core/standard/protocol/packet"
)

type AddVolumeEntity struct{}

func (pk *AddVolumeEntity) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.AddVolumeEntity{}
	input := standard.(*standardPacket.AddVolumeEntity)

	p.EntityRuntimeID = uint32(input.EntityRuntimeID)
	p.EntityMetadata = input.EntityMetadata
	p.EncodingIdentifier = input.EncodingIdentifier
	p.InstanceIdentifier = input.InstanceIdentifier
	p.Bounds = [2]neteaseProtocol.BlockPos{p.Bounds[0], p.Bounds[1]}
	p.Dimension = input.Dimension
	p.EngineVersion = input.EngineVersion

	return &p
}

func (pk *AddVolumeEntity) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.AddVolumeEntity{}
	input := netease.(*neteasePacket.AddVolumeEntity)

	p.EntityRuntimeID = uint64(input.EntityRuntimeID)
	p.EntityMetadata = input.EntityMetadata
	p.EncodingIdentifier = input.EncodingIdentifier
	p.InstanceIdentifier = input.InstanceIdentifier
	p.Bounds = [2]standardProtocol.BlockPos{p.Bounds[0], p.Bounds[1]}
	p.Dimension = input.Dimension
	p.EngineVersion = input.EngineVersion

	return &p
}
