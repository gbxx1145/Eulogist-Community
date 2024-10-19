package persistence_data

import (
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"

	"github.com/go-gl/mathgl/mgl32"
)

// 用户当前所处维度的数据
type BotDimensionData struct {
	Dimension  int32                // 用户当前所处的维度 ID
	Position   mgl32.Vec3           // 用户在抵达目标维度后的位置
	Respawn    bool                 // 当此维度交换是否因为重生而发生
	DataCache  DimensionChangeCache // 缓存数据，用于二级维度交换
	ChangeDown bool                 // 用户是否已完成当此维度交换
}

// 缓存数据，用于二级维度交换
type DimensionChangeCache struct {
	LevelChunk   []standardPacket.LevelChunk   // 缓存的区块数据
	AddActor     []standardPacket.AddActor     // 缓存的实体数据
	SetActorData []standardPacket.SetActorData // 缓存的实体设置数据
	AddItemActor []standardPacket.AddItemActor // 缓存的掉落物数据
}

/*
此函数用于解决非原版维度的相关问题。

from 指代维度更改前玩家所在的维度，
而 to 指代维度更改后玩家所在的维度。

玩家将被传送到一个中间人维度，
即便该维度不是玩家真正所处的维度，
然后再被传送到正确的维度。

此函数的作用便是根据 from 和 to 选出该中间人维度
*/
func (b *BotDimensionData) GetTransferDimensionID(from int32, to int32) (dimension int32) {
	// 此时源维度为原版维度，但目标维度不是，
	// 则中间人维度应当是末地或下界，
	// 然后，再被传送到类似于主世界的维度
	if from <= standardPacket.DimensionEnd && to > standardPacket.DimensionEnd {
		if from == standardPacket.DimensionEnd {
			return standardPacket.DimensionNether
		}
		return standardPacket.DimensionEnd
	}
	// 此时源维度不是原版维度，但目标维度是，
	// 则中间人维度应当是末地或下界，
	// 然后，再被传送到为原版维度的目标维度
	if from > standardPacket.DimensionEnd && to <= standardPacket.DimensionEnd {
		if to == standardPacket.DimensionEnd {
			return standardPacket.DimensionNether
		}
		return standardPacket.DimensionEnd
	}
	// 此时源维度和目标维度都不是原版维度，
	// 则中间人维度不妨选择末地，
	// 然后，玩家会被传送到类似于主世界的维度
	return standardPacket.DimensionEnd
}
