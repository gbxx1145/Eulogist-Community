package chunk_process

import (
	"Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/minecraft/netease/protocol/block_actors"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	"Eulogist/tools/chunk_process/chunk"
	"Eulogist/tools/chunk_process/cube"
	"Eulogist/tools/netease_blocks/blocks"
	"bytes"
	"fmt"

	"Eulogist/core/minecraft/standard/nbt"

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

// DecodeNetEaseSubChunk 解析区块 SubChunk。
//
// 解析过程中将会自动将方块的运行时 ID 与国际版对齐，
// 然后，再将其中包含的方块实体数据翻译为正常形式，
// 最终，将翻译结果回写到其自身
func DecodeNetEaseSubChunk(pk *neteasePacket.SubChunk) {
	for index, value := range pk.SubChunkEntries {
		multipleBlockNBT := make([]map[string]any, 0)

		buf := bytes.NewBuffer(value.RawPayload)
		chunkGet, idx, err := chunk.DecodeSubChunkPublic(blocks.AIR_RUNTIMEID, buf, LookUpCubeRange(pk.Dimension))
		if err != nil {
			continue
		}

		func() {
			defer func() {
				recover()
			}()

			for len(buf.Bytes()) > 0 {
				var blockNBT map[string]any
				_ = nbt.NewDecoderWithEncoding(buf, nbt.NetworkLittleEndian).Decode(&blockNBT)
				blockNBT, _ = TranslateNBT(blockNBT)
				multipleBlockNBT = append(multipleBlockNBT, blockNBT)
			}
		}()

		subChunk := chunk.EncodeSubChunk(chunkGet, chunk.NetworkEncoding, int(idx))
		blockEntityBuf := bytes.NewBuffer(nil)
		for _, v := range multipleBlockNBT {
			_ = nbt.NewEncoderWithEncoding(blockEntityBuf, nbt.NetworkLittleEndian).Encode(v)
		}

		pk.SubChunkEntries[index].RawPayload = append(subChunk, blockEntityBuf.Bytes()...)
	}
}

/*
DecodeNetEaseLevelChunk 解析区块 LevelChunk。

解析过程中将会自动将方块的运行时 ID 与国际版对齐，
但不会翻译方块实体数据，因为它们本身就是非平铺的。
最后，翻译结果会被回写到数据包 pk 自身。

对于租赁服，由于租赁服启用了子区块请求系统，
因此这个数据包似乎只用于告知生物的群系数据。
对于其他的部分，例如区块内包含的方块，
亦或区块内包含的方块实体，
它们存在于 SubChunk 数据包携带的负载。

因此，该函数主要为未实现子区块请求系统的 网络游戏 服务。
但此处不保证完整性，因为方块的运行时在这种情况下不能完全对齐。
究其原因，这可能是因为网络游戏引入了一些自定义方块，
这导致方块调色板及运行时 ID 表发生偏移或完全重排
*/
func DecodeNetEaseLevelChunk(pk *neteasePacket.LevelChunk) {
	multipleBlockNBT := make([]map[string]any, 0)
	buf := bytes.NewBuffer(pk.RawPayload)

	// 对于租赁服而言，下方的 NetworkDecode 将永远返回错误，
	// 因为租赁服启用了子区块请求系统，
	// 这导致 pk.RawPayload 将只携带区块的生物群系数据
	chunkGet, err := chunk.NetworkDecode(blocks.AIR_RUNTIMEID, buf, int(pk.SubChunkCount), LookUpCubeRange(0))
	if err != nil {
		// 无法保证 DecodeBiomesPublic 可以永远不返回错误，
		// 这是一个难以处理的问题
		chunkGet, err = chunk.DecodeBiomesPublic(blocks.AIR_RUNTIMEID, buf, LookUpCubeRange(0))
		if err != nil {
			return
		}
		// 似乎为 pk.RawPayload 填写空值也能使 Minecraft 正常工作，
		// 目前尚不清楚该字段的变化会给实际的地形表现产生何种影响
		pk.RawPayload = chunk.EncodeBiomes(chunkGet, chunk.NetworkEncoding)
		return
	}

	// Length of 1 byte for the border block count.
	borderBlockCount, _ := buf.ReadByte()

	func() {
		defer func() {
			recover()
		}()

		for len(buf.Bytes()) > 0 {
			var blockNBT map[string]any
			_ = nbt.NewDecoderWithEncoding(buf, nbt.NetworkLittleEndian).Decode(&blockNBT)

			// 经过观察，布吉岛的方块实体数据不是 __tag NBT。
			// 因此，也许对于网络游戏而言，
			// 方块实体数据无需再进行翻译
			{
				// blockNBT, _ = TranslateNBT(blockNBT)
			}

			multipleBlockNBT = append(multipleBlockNBT, blockNBT)
		}
	}()

	data := chunk.Encode(chunkGet, chunk.NetworkEncoding)
	chunkBuf := bytes.NewBuffer(nil)
	for _, s := range data.SubChunks {
		_, _ = chunkBuf.Write(s)
	}
	_, _ = chunkBuf.Write(data.Biomes)

	// Length of 1 byte for the border block count.
	chunkBuf.WriteByte(borderBlockCount)

	for _, v := range multipleBlockNBT {
		_ = nbt.NewEncoderWithEncoding(chunkBuf, nbt.NetworkLittleEndian).Encode(v)
	}

	pk.RawPayload = chunkBuf.Bytes()
	pk.SubChunkCount = uint32(len(data.SubChunks))
}
