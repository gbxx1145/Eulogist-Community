package world

// Block is a block that may be placed or found in a world. In addition, the block may also be added to an
// inventory: It is also an item.
// Every Block implementation must be able to be hashed as key in a map.
type Block interface {
	// EncodeBlock encodes the block to a string ID such as 'minecraft:grass' and properties associated
	// with the block.
	EncodeBlock() (string, map[string]any)
	// Hash returns a unique identifier of the block including the block states. This function is used internally to
	// convert a block to a single integer which can be used in map lookups. The hash produced therefore does not need
	// to match anything in the game, but it must be unique among all registered blocks.
	// The tool in `/cmd/blockhash` may be used to automatically generate block hashes of blocks in a package.
	Hash() uint64
	// Model returns the BlockModel of the Block.
}
