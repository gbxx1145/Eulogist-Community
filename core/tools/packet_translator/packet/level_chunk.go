package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/tools/chunk_process"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type LevelChunk struct{}

func (pk *LevelChunk) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.LevelChunk{}
	input := standard.(*standardPacket.LevelChunk)

	p.Position = neteaseProtocol.ChunkPos(input.Position)
	p.HighestSubChunk = input.HighestSubChunk
	p.SubChunkCount = input.SubChunkCount
	p.CacheEnabled = input.CacheEnabled
	p.BlobHashes = input.BlobHashes
	p.RawPayload = input.RawPayload

	return &p
}

func (pk *LevelChunk) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.LevelChunk{}
	input := netease.(*neteasePacket.LevelChunk)

	chunk_process.DecodeNetEaseLevelChunk(input)

	p.Position = standardProtocol.ChunkPos(input.Position)
	p.HighestSubChunk = input.HighestSubChunk
	p.SubChunkCount = input.SubChunkCount
	p.CacheEnabled = input.CacheEnabled
	p.BlobHashes = input.BlobHashes
	p.RawPayload = input.RawPayload

	return &p
}
