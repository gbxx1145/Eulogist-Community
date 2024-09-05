package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type CraftingData struct{}

// 将 standard 转换为 neteaseProtocol.Recipe
func (pk *CraftingData) ToNetEaseRecipe(
	standard standardProtocol.Recipe,
) neteaseProtocol.Recipe {
	switch data := standard.(type) {
	case *standardProtocol.ShapelessRecipe:
		return &neteaseProtocol.ShapelessRecipe{
			RecipeID: data.RecipeID,
			Input: packet_translate_struct.ConvertSlice(
				data.Input,
				packet_translate_struct.ToNetEaseItemDescriptorCount,
			),
		}
	case *standardProtocol.ShapedRecipe:
		return &neteaseProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: packet_translate_struct.ConvertSlice(
				data.Input,
				packet_translate_struct.ToNetEaseItemDescriptorCount,
			),
			Output: packet_translate_struct.ConvertSlice(
				data.Output,
				func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
					return packet_translate_struct.ConvertToNetEaseItemStack(from)
				},
			),
		}
	case *standardProtocol.FurnaceRecipe:
		return &neteaseProtocol.FurnaceRecipe{
			InputType: neteaseProtocol.ItemType(data.InputType),
			Output:    packet_translate_struct.ConvertToNetEaseItemStack(data.Output),
			Block:     data.Block,
		}
	case *standardProtocol.FurnaceDataRecipe:
		return &neteaseProtocol.FurnaceDataRecipe{
			FurnaceRecipe: neteaseProtocol.FurnaceRecipe{
				InputType: neteaseProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    packet_translate_struct.ConvertToNetEaseItemStack(data.FurnaceRecipe.Output),
				Block:     data.FurnaceRecipe.Block,
			},
		}
	case *standardProtocol.MultiRecipe:
		return &neteaseProtocol.MultiRecipe{
			UUID:            data.UUID,
			RecipeNetworkID: data.RecipeNetworkID,
		}
	case *standardProtocol.ShulkerBoxRecipe:
		return &neteaseProtocol.ShulkerBoxRecipe{
			ShapelessRecipe: neteaseProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapelessRecipe.Input,
					packet_translate_struct.ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapelessChemistryRecipe:
		return &neteaseProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: neteaseProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapelessRecipe.Input,
					packet_translate_struct.ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapedChemistryRecipe:
		return &neteaseProtocol.ShapedChemistryRecipe{
			ShapedRecipe: neteaseProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapedRecipe.Input,
					packet_translate_struct.ToNetEaseItemDescriptorCount,
				),
				Output: packet_translate_struct.ConvertSlice(
					data.ShapedRecipe.Output,
					func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
						return packet_translate_struct.ConvertToNetEaseItemStack(from)
					},
				),
			},
		}
	case *standardProtocol.SmithingTransformRecipe:
		return &neteaseProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        packet_translate_struct.ToNetEaseItemDescriptorCount(data.Template),
			Base:            packet_translate_struct.ToNetEaseItemDescriptorCount(data.Base),
			Addition:        packet_translate_struct.ToNetEaseItemDescriptorCount(data.Addition),
			Result:          packet_translate_struct.ConvertToNetEaseItemStack(data.Result),
			Block:           data.Block,
		}
	case *standardProtocol.SmithingTrimRecipe:
		return &neteaseProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        packet_translate_struct.ToNetEaseItemDescriptorCount(data.Template),
			Base:            packet_translate_struct.ToNetEaseItemDescriptorCount(data.Base),
			Addition:        packet_translate_struct.ToNetEaseItemDescriptorCount(data.Addition),
			Block:           data.Block,
		}
	}

	panic("ToNetEaseRecipe: Invalid recipe enum")
}

// 将 netease 转换为 standardProtocol.Recipe
func (pk *CraftingData) ToStandardRecipe(
	netease neteaseProtocol.Recipe,
) standardProtocol.Recipe {
	switch data := netease.(type) {
	case *neteaseProtocol.ShapelessRecipe:
		return &standardProtocol.ShapelessRecipe{
			RecipeID: data.RecipeID,
			Input: packet_translate_struct.ConvertSlice(
				data.Input,
				packet_translate_struct.ToStandardItemDescriptorCount,
			),
		}
	case *neteaseProtocol.ShapedRecipe:
		return &standardProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: packet_translate_struct.ConvertSlice(
				data.Input,
				packet_translate_struct.ToStandardItemDescriptorCount,
			),
			Output: packet_translate_struct.ConvertSlice(
				data.Output,
				func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
					return packet_translate_struct.ConvertToStandardItemStack(from)
				},
			),
		}
	case *neteaseProtocol.FurnaceRecipe:
		return &standardProtocol.FurnaceRecipe{
			InputType: standardProtocol.ItemType(data.InputType),
			Output:    packet_translate_struct.ConvertToStandardItemStack(data.Output),
			Block:     data.Block,
		}
	case *neteaseProtocol.FurnaceDataRecipe:
		return &standardProtocol.FurnaceDataRecipe{
			FurnaceRecipe: standardProtocol.FurnaceRecipe{
				InputType: standardProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    packet_translate_struct.ConvertToStandardItemStack(data.FurnaceRecipe.Output),
				Block:     data.FurnaceRecipe.Block,
			},
		}
	case *neteaseProtocol.MultiRecipe:
		return &standardProtocol.MultiRecipe{
			UUID:            data.UUID,
			RecipeNetworkID: data.RecipeNetworkID,
		}
	case *neteaseProtocol.ShulkerBoxRecipe:
		return &standardProtocol.ShulkerBoxRecipe{
			ShapelessRecipe: standardProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapelessRecipe.Input,
					packet_translate_struct.ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapelessChemistryRecipe:
		return &standardProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: standardProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapelessRecipe.Input,
					packet_translate_struct.ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapedChemistryRecipe:
		return &standardProtocol.ShapedChemistryRecipe{
			ShapedRecipe: standardProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: packet_translate_struct.ConvertSlice(
					data.ShapedRecipe.Input,
					packet_translate_struct.ToStandardItemDescriptorCount,
				),
				Output: packet_translate_struct.ConvertSlice(
					data.ShapedRecipe.Output,
					func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
						return packet_translate_struct.ConvertToStandardItemStack(from)
					},
				),
			},
		}
	case *neteaseProtocol.SmithingTransformRecipe:
		return &standardProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        packet_translate_struct.ToStandardItemDescriptorCount(data.Template),
			Base:            packet_translate_struct.ToStandardItemDescriptorCount(data.Base),
			Addition:        packet_translate_struct.ToStandardItemDescriptorCount(data.Addition),
			Result:          packet_translate_struct.ConvertToStandardItemStack(data.Result),
			Block:           data.Block,
		}
	case *neteaseProtocol.SmithingTrimRecipe:
		return &standardProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        packet_translate_struct.ToStandardItemDescriptorCount(data.Template),
			Base:            packet_translate_struct.ToStandardItemDescriptorCount(data.Base),
			Addition:        packet_translate_struct.ToStandardItemDescriptorCount(data.Addition),
			Block:           data.Block,
		}
	}

	panic("ToStandardRecipe: Invalid recipe enum")
}

func (pk *CraftingData) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.CraftingData{}
	input := standard.(*standardPacket.CraftingData)

	p.ClearRecipes = input.ClearRecipes

	p.Recipes = packet_translate_struct.ConvertSlice(
		input.Recipes,
		pk.ToNetEaseRecipe,
	)
	p.PotionRecipes = packet_translate_struct.ConvertSlice(
		input.PotionRecipes,
		func(from standardProtocol.PotionRecipe) neteaseProtocol.PotionRecipe {
			return neteaseProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = packet_translate_struct.ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from standardProtocol.PotionContainerChangeRecipe) neteaseProtocol.PotionContainerChangeRecipe {
			return neteaseProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = packet_translate_struct.ConvertSlice(
		input.MaterialReducers,
		func(from standardProtocol.MaterialReducer) neteaseProtocol.MaterialReducer {
			return neteaseProtocol.MaterialReducer{
				InputItem: neteaseProtocol.ItemType(from.InputItem),
				Outputs: packet_translate_struct.ConvertSlice(
					from.Outputs,
					func(from standardProtocol.MaterialReducerOutput) neteaseProtocol.MaterialReducerOutput {
						return neteaseProtocol.MaterialReducerOutput(from)
					},
				),
			}
		},
	)

	p.Unknown1 = make([]byte, 0)
	p.Unknown2 = make([]byte, 0)
	p.Unknown3 = make([]byte, 0)

	return &p
}

func (pk *CraftingData) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.CraftingData{}
	input := netease.(*neteasePacket.CraftingData)

	p.ClearRecipes = input.ClearRecipes

	p.Recipes = packet_translate_struct.ConvertSlice(
		input.Recipes,
		pk.ToStandardRecipe,
	)
	p.PotionRecipes = packet_translate_struct.ConvertSlice(
		input.PotionRecipes,
		func(from neteaseProtocol.PotionRecipe) standardProtocol.PotionRecipe {
			return standardProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = packet_translate_struct.ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from neteaseProtocol.PotionContainerChangeRecipe) standardProtocol.PotionContainerChangeRecipe {
			return standardProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = packet_translate_struct.ConvertSlice(
		input.MaterialReducers,
		func(from neteaseProtocol.MaterialReducer) standardProtocol.MaterialReducer {
			return standardProtocol.MaterialReducer{
				InputItem: standardProtocol.ItemType(from.InputItem),
				Outputs: packet_translate_struct.ConvertSlice(
					from.Outputs,
					func(from neteaseProtocol.MaterialReducerOutput) standardProtocol.MaterialReducerOutput {
						return standardProtocol.MaterialReducerOutput(from)
					},
				),
			}
		},
	)

	return &p
}
