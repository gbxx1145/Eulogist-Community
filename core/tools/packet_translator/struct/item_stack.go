package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	"Eulogist/tools/chunk_process/chunk"
	"Eulogist/tools/netease_blocks/blocks"
	"fmt"
)

// ...
func convertItemNBTMap(mapping map[string]any) map[string]any {
	convertResult := map[string]any{}

	for key, value := range mapping {
		val, ok := value.([]any)

		if key == "ench" && ok {
			convertResult[key] = convertEnchantSlice(val)
			continue
		}

		if ok {
			convertResult[key] = convertItemNBTSlice(val)
			continue
		}

		convertResult[key] = value
	}

	return convertResult
}

// ...
func convertItemNBTSlice(slice []any) []any {
	convertResult := []any{}

	for _, value := range slice {
		val, ok := value.(map[string]any)
		if ok {
			convertResult = append(convertResult, convertItemNBTMap(val))
		} else {
			convertResult = append(convertResult, value)
		}
	}

	return convertResult
}

// ...
func convertEnchantSlice(slice []any) []any {
	convertResult := []any{}

	for _, value := range slice {
		val, ok := value.(map[string]any)
		if !ok {
			return slice
		}

		val["modEnchant"] = ""
		convertResult = append(convertResult, val)
	}

	return convertResult
}

// 将 from 转换为 neteaseProtocol.ItemStack
func ConvertToNetEaseItemStack(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
	blockName, blockStates, foundA := chunk.RuntimeIDToState(uint32(from.BlockRuntimeID))
	neteaseBlockRuntimeID, foundB := blocks.BlockNameAndStateToRuntimeID(blockName, blockStates)
	if foundA && foundB {
		from.BlockRuntimeID = int32(neteaseBlockRuntimeID)
	}

	itemStack := neteaseProtocol.ItemStack{
		ItemType:       neteaseProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}

	itemStack.NBTData = convertItemNBTMap(from.NBTData)

	return itemStack
}

// 将 from 转换为 standardProtocol.ItemStack
func ConvertToStandardItemStack(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
	blockName, blockStates, foundA := blocks.RuntimeIDToState(uint32(from.BlockRuntimeID))
	standardRuntimeID, foundB := chunk.StateToRuntimeID(fmt.Sprintf("minecraft:%s", blockName), blockStates)
	if foundA && foundB {
		from.BlockRuntimeID = int32(standardRuntimeID)
	}

	return standardProtocol.ItemStack{
		ItemType:       standardProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}
}
