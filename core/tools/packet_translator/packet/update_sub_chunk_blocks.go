package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type UpdateSubChunkBlocks struct{}

func (pk *UpdateSubChunkBlocks) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.UpdateSubChunkBlocks{}
	input := standard.(*standardPacket.UpdateSubChunkBlocks)

	p.Position = neteaseProtocol.SubChunkPos(input.Position)

	p.Blocks = packet_translator.ConvertSlice(
		input.Blocks,
		func(from standardProtocol.BlockChangeEntry) neteaseProtocol.BlockChangeEntry {
			return neteaseProtocol.BlockChangeEntry{
				BlockPos:                   neteaseProtocol.BlockPos(from.BlockPos),
				BlockRuntimeID:             from.BlockRuntimeID,
				Flags:                      from.Flags,
				SyncedUpdateEntityUniqueID: from.SyncedUpdateEntityUniqueID,
				SyncedUpdateType:           from.SyncedUpdateType,
			}
		},
	)
	p.Extra = packet_translator.ConvertSlice(
		input.Extra,
		func(from standardProtocol.BlockChangeEntry) neteaseProtocol.BlockChangeEntry {
			return neteaseProtocol.BlockChangeEntry{
				BlockPos:                   neteaseProtocol.BlockPos(from.BlockPos),
				BlockRuntimeID:             from.BlockRuntimeID,
				Flags:                      from.Flags,
				SyncedUpdateEntityUniqueID: from.SyncedUpdateEntityUniqueID,
				SyncedUpdateType:           from.SyncedUpdateType,
			}
		},
	)

	return &p
}

func (pk *UpdateSubChunkBlocks) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.UpdateSubChunkBlocks{}
	input := netease.(*neteasePacket.UpdateSubChunkBlocks)

	p.Position = standardProtocol.SubChunkPos(input.Position)

	p.Blocks = packet_translator.ConvertSlice(
		input.Blocks,
		func(from neteaseProtocol.BlockChangeEntry) standardProtocol.BlockChangeEntry {
			return standardProtocol.BlockChangeEntry{
				BlockPos:                   standardProtocol.BlockPos(from.BlockPos),
				BlockRuntimeID:             from.BlockRuntimeID,
				Flags:                      from.Flags,
				SyncedUpdateEntityUniqueID: from.SyncedUpdateEntityUniqueID,
				SyncedUpdateType:           from.SyncedUpdateType,
			}
		},
	)
	p.Extra = packet_translator.ConvertSlice(
		input.Extra,
		func(from neteaseProtocol.BlockChangeEntry) standardProtocol.BlockChangeEntry {
			return standardProtocol.BlockChangeEntry{
				BlockPos:                   standardProtocol.BlockPos(from.BlockPos),
				BlockRuntimeID:             from.BlockRuntimeID,
				Flags:                      from.Flags,
				SyncedUpdateEntityUniqueID: from.SyncedUpdateEntityUniqueID,
				SyncedUpdateType:           from.SyncedUpdateType,
			}
		},
	)

	return &p
}
