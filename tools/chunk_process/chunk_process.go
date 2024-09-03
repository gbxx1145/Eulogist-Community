package chunk_process

import (
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/block_actors"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/tools/chunk_process/chunk"
	"Eulogist/tools/chunk_process/cube"
	"Eulogist/tools/netease_blocks/blocks"
	"bytes"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

const (
	DimensionOverworld = iota
	DimensionNether
	DimensionEnd
)

// ...
func LookUpCubeRange(dimension int32) cube.Range {
	switch dimension {
	case DimensionOverworld:
		return cube.Range{-64, 319}
	case DimensionNether:
		return cube.Range{0, 127}
	case DimensionEnd:
		return cube.Range{0, 255}
	}

	return cube.Range{-64, 319}
}

// ...
func TranslateNBT(origNbt map[string]interface{}) (out map[string]interface{}, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("nbt translate err: %v", r)
		}
	}()

	id := origNbt["id"].(string)
	tag := ([]byte)(origNbt["__tag"].(string))

	buffer := bytes.NewBuffer(tag)
	reader := protocol.NewReader(buffer, 0, false)

	block, found := block_actors.NewPool()[id]
	if !found {
		return nil, fmt.Errorf("TranslateNBT: NBT Block %#v not supported; origNbt = %#v", id, origNbt)
	}
	block.Marshal(reader)

	mapstructure.Decode(block, &out)
	out["id"] = id
	out["x"] = origNbt["x"]
	out["y"] = origNbt["y"]
	out["z"] = origNbt["z"]

	return out, nil
}

// DecodeNetEaseSubChunk 解析子区块 pk，
// 并将其中包含的方块实体数据翻译为正常形式，
// 然后再将翻译结果回写到其自身
func DecodeNetEaseSubChunk(pk *neteasePacket.SubChunk) {
	for index, value := range pk.SubChunkEntries {
		multipleBlockNBT := make([]map[string]any, 0)

		buf := bytes.NewBuffer(value.RawPayload)
		chunkGet, _ := chunk.DecodeSubChunkPublic(blocks.AIR_RUNTIMEID, buf, LookUpCubeRange(pk.Dimension))

		for len(buf.Bytes()) > 0 {
			var blockNBT map[string]any
			_ = nbt.NewDecoderWithEncoding(buf, nbt.NetworkLittleEndian).Decode(&blockNBT)
			blockNBT, _ = TranslateNBT(blockNBT)
			multipleBlockNBT = append(multipleBlockNBT, blockNBT)
		}

		subChunk := chunk.EncodeSubChunk(chunkGet, chunk.NetworkEncoding, 0)
		blockEntityBuf := bytes.NewBuffer(nil)
		for _, v := range multipleBlockNBT {
			_ = nbt.NewEncoderWithEncoding(blockEntityBuf, nbt.NetworkLittleEndian).Encode(v)
		}

		pk.SubChunkEntries[index].RawPayload = append(subChunk, blockEntityBuf.Bytes()...)
	}
}

// DecodeNetEaseLevelChunk 解析区块 pk，
// 并将其中包含的方块实体数据翻译为正常形式，
// 然后再将翻译结果回写到其自身
func DecodeNetEaseLevelChunk(pk *neteasePacket.LevelChunk) {
	multipleBlockNBT := make([]map[string]any, 0)
	buf := bytes.NewBuffer(pk.RawPayload)
	chunkGet, _ := chunk.NetworkDecode(blocks.AIR_RUNTIMEID, buf, int(pk.SubChunkCount), LookUpCubeRange(0))

	// Length of 1 byte for the border block count.
	borderBlockCount, _ := buf.ReadByte()

	for len(buf.Bytes()) > 0 {
		var blockNBT map[string]any
		_ = nbt.NewDecoderWithEncoding(buf, nbt.NetworkLittleEndian).Decode(&blockNBT)
		blockNBT, _ = TranslateNBT(blockNBT)
		multipleBlockNBT = append(multipleBlockNBT, blockNBT)
	}

	chunkData := chunk.Encode(chunkGet, chunk.NetworkEncoding)

	chunkBuf := bytes.NewBuffer(nil)
	for _, s := range chunkData.SubChunks {
		_, _ = chunkBuf.Write(s)
	}
	_, _ = chunkBuf.Write(chunkData.Biomes)

	// Length of 1 byte for the border block count.
	chunkBuf.WriteByte(borderBlockCount)

	for _, value := range multipleBlockNBT {
		_ = nbt.NewEncoderWithEncoding(chunkBuf, nbt.NetworkLittleEndian).Encode(value)
	}

	pk.RawPayload = chunkBuf.Bytes()
}
