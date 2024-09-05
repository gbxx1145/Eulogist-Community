package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"
	"Eulogist/tools/chunk_process"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type SubChunk struct{}

func (pk *SubChunk) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.SubChunk{}
	input := standard.(*standardPacket.SubChunk)

	p.CacheEnabled = input.CacheEnabled
	p.Dimension = input.Dimension
	p.Position = neteaseProtocol.SubChunkPos(input.Position)
	p.SubChunkEntries = packet_translate_struct.ConvertSlice(
		input.SubChunkEntries,
		func(from standardProtocol.SubChunkEntry) neteaseProtocol.SubChunkEntry {
			return neteaseProtocol.SubChunkEntry{
				Offset:        neteaseProtocol.SubChunkOffset(from.Offset),
				Result:        from.Result,
				RawPayload:    from.RawPayload,
				HeightMapType: from.HeightMapType,
				HeightMapData: packet_translate_struct.ConvertSlice(
					from.HeightMapData,
					func(from int8) uint8 {
						return uint8(from)
					},
				),
				BlobHash: from.BlobHash,
			}
		},
	)

	return &p
}

func (pk *SubChunk) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.SubChunk{}
	input := netease.(*neteasePacket.SubChunk)

	chunk_process.DecodeNetEaseSubChunk(input)

	p.CacheEnabled = input.CacheEnabled
	p.Dimension = input.Dimension
	p.Position = standardProtocol.SubChunkPos(input.Position)
	p.SubChunkEntries = packet_translate_struct.ConvertSlice(
		input.SubChunkEntries,
		func(from neteaseProtocol.SubChunkEntry) standardProtocol.SubChunkEntry {
			return standardProtocol.SubChunkEntry{
				Offset:        standardProtocol.SubChunkOffset(from.Offset),
				Result:        from.Result,
				RawPayload:    from.RawPayload,
				HeightMapType: from.HeightMapType,
				HeightMapData: packet_translate_struct.ConvertSlice(
					from.HeightMapData,
					func(from uint8) int8 {
						return int8(from)
					},
				),
				BlobHash: from.BlobHash,
			}
		},
	)

	return &p
}
