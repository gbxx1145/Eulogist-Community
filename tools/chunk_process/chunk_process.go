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

	"Eulogist/core/standard/nbt"

	"github.com/mitchellh/mapstructure"
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
		chunkGet, idx, err := chunk.DecodeSubChunkPublic(blocks.AIR_RUNTIMEID, buf, LookUpCubeRange(pk.Dimension))
		if err != nil {
			continue
		}

		for len(buf.Bytes()) > 0 {
			var blockNBT map[string]any
			_ = nbt.NewDecoderWithEncoding(buf, nbt.NetworkLittleEndian).Decode(&blockNBT)
			blockNBT, _ = TranslateNBT(blockNBT)
			multipleBlockNBT = append(multipleBlockNBT, blockNBT)
		}

		if len(buf.Bytes()) > 0 {
			fmt.Println(buf.Bytes())
		}

		subChunk := chunk.EncodeSubChunk(chunkGet, chunk.NetworkEncoding, int(idx))
		blockEntityBuf := bytes.NewBuffer(nil)
		for _, v := range multipleBlockNBT {
			_ = nbt.NewEncoderWithEncoding(blockEntityBuf, nbt.NetworkLittleEndian).Encode(v)
		}

		pk.SubChunkEntries[index].RawPayload = append(subChunk, blockEntityBuf.Bytes()...)
	}
}
