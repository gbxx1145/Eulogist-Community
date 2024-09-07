package chunk_process

// CAN NOT USE, SO WE DO NO TRANSLATE.

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
*/
