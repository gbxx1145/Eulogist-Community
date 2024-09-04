package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type RequestChunkRadius struct{}

func (pk *RequestChunkRadius) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.RequestChunkRadius{}
	input := standard.(*standardPacket.RequestChunkRadius)

	p.ChunkRadius = input.ChunkRadius
	p.MaxChunkRadius = uint8(input.MaxChunkRadius)

	return &p
}

func (pk *RequestChunkRadius) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.RequestChunkRadius{}
	input := netease.(*neteasePacket.RequestChunkRadius)

	p.ChunkRadius = input.ChunkRadius
	p.MaxChunkRadius = int32(input.MaxChunkRadius)

	return &p
}
