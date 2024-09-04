package packet_translator

import (
	"Eulogist/tools/chunk_process/chunk"
	"Eulogist/tools/netease_blocks/blocks"
	"fmt"
)

// 将国际版的方块运行时转换为网易协议下的方块运行时
func ConvertToNetEaseBlockRuntimeID(standardRuntimeID uint32) (neteaseBlockRuntimeID uint32, found bool) {
	blockName, blockStates, found := chunk.RuntimeIDToState(standardRuntimeID)
	if !found {
		return 0, found
	}

	neteaseBlockRuntimeID, found = blocks.BlockNameAndStateToRuntimeID(blockName, blockStates)
	if !found {
		return 0, found
	}

	return neteaseBlockRuntimeID, true
}

// 将网易版本的方块运行时转换为国际版下的方块运行时
func ConvertToStandardBlockRuntimeID(neteaseBlockRuntimeID uint32) (standardRuntimeID uint32, found bool) {
	blockName, blockStates, found := blocks.RuntimeIDToState(neteaseBlockRuntimeID)
	if !found {
		return 0, found
	}

	standardRuntimeID, found = chunk.StateToRuntimeID(fmt.Sprintf("minecraft:%s", blockName), blockStates)
	if !found {
		return 0, found
	}

	return standardRuntimeID, true
}
