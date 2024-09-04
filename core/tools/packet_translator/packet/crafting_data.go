package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "Eulogist/core/standard/protocol"
	standardPacket "Eulogist/core/standard/protocol/packet"
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
			Input: ConvertSlice(
				data.Input,
				ToNetEaseItemDescriptorCount,
			),
		}
	case *standardProtocol.ShapedRecipe:
		return &neteaseProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: ConvertSlice(
				data.Input,
				ToNetEaseItemDescriptorCount,
			),
			Output: ConvertSlice(
				data.Output,
				func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
					return ConvertToNetEaseItemStack(from)
				},
			),
		}
	case *standardProtocol.FurnaceRecipe:
		return &neteaseProtocol.FurnaceRecipe{
			InputType: neteaseProtocol.ItemType(data.InputType),
			Output:    ConvertToNetEaseItemStack(data.Output),
			Block:     data.Block,
		}
	case *standardProtocol.FurnaceDataRecipe:
		return &neteaseProtocol.FurnaceDataRecipe{
			FurnaceRecipe: neteaseProtocol.FurnaceRecipe{
				InputType: neteaseProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    ConvertToNetEaseItemStack(data.FurnaceRecipe.Output),
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
				Input: ConvertSlice(
					data.ShapelessRecipe.Input,
					ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapelessChemistryRecipe:
		return &neteaseProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: neteaseProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: ConvertSlice(
					data.ShapelessRecipe.Input,
					ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapedChemistryRecipe:
		return &neteaseProtocol.ShapedChemistryRecipe{
			ShapedRecipe: neteaseProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: ConvertSlice(
					data.ShapedRecipe.Input,
					ToNetEaseItemDescriptorCount,
				),
				Output: ConvertSlice(
					data.ShapedRecipe.Output,
					func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
						return ConvertToNetEaseItemStack(from)
					},
				),
			},
		}
	case *standardProtocol.SmithingTransformRecipe:
		return &neteaseProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        ToNetEaseItemDescriptorCount(data.Template),
			Base:            ToNetEaseItemDescriptorCount(data.Base),
			Addition:        ToNetEaseItemDescriptorCount(data.Addition),
			Result:          ConvertToNetEaseItemStack(data.Result),
			Block:           data.Block,
		}
	case *standardProtocol.SmithingTrimRecipe:
		return &neteaseProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        ToNetEaseItemDescriptorCount(data.Template),
			Base:            ToNetEaseItemDescriptorCount(data.Base),
			Addition:        ToNetEaseItemDescriptorCount(data.Addition),
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
			Input: ConvertSlice(
				data.Input,
				ToStandardItemDescriptorCount,
			),
		}
	case *neteaseProtocol.ShapedRecipe:
		return &standardProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: ConvertSlice(
				data.Input,
				ToStandardItemDescriptorCount,
			),
			Output: ConvertSlice(
				data.Output,
				func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
					return ConvertToStandardItemStack(from)
				},
			),
		}
	case *neteaseProtocol.FurnaceRecipe:
		return &standardProtocol.FurnaceRecipe{
			InputType: standardProtocol.ItemType(data.InputType),
			Output:    ConvertToStandardItemStack(data.Output),
			Block:     data.Block,
		}
	case *neteaseProtocol.FurnaceDataRecipe:
		return &standardProtocol.FurnaceDataRecipe{
			FurnaceRecipe: standardProtocol.FurnaceRecipe{
				InputType: standardProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    ConvertToStandardItemStack(data.FurnaceRecipe.Output),
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
				Input: ConvertSlice(
					data.ShapelessRecipe.Input,
					ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapelessChemistryRecipe:
		return &standardProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: standardProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: ConvertSlice(
					data.ShapelessRecipe.Input,
					ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapedChemistryRecipe:
		return &standardProtocol.ShapedChemistryRecipe{
			ShapedRecipe: standardProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: ConvertSlice(
					data.ShapedRecipe.Input,
					ToStandardItemDescriptorCount,
				),
				Output: ConvertSlice(
					data.ShapedRecipe.Output,
					func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
						return ConvertToStandardItemStack(from)
					},
				),
			},
		}
	case *neteaseProtocol.SmithingTransformRecipe:
		return &standardProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        ToStandardItemDescriptorCount(data.Template),
			Base:            ToStandardItemDescriptorCount(data.Base),
			Addition:        ToStandardItemDescriptorCount(data.Addition),
			Result:          ConvertToStandardItemStack(data.Result),
			Block:           data.Block,
		}
	case *neteaseProtocol.SmithingTrimRecipe:
		return &standardProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        ToStandardItemDescriptorCount(data.Template),
			Base:            ToStandardItemDescriptorCount(data.Base),
			Addition:        ToStandardItemDescriptorCount(data.Addition),
			Block:           data.Block,
		}
	}

	panic("ToStandardRecipe: Invalid recipe enum")
}

func (pk *CraftingData) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.CraftingData{}
	input := standard.(*standardPacket.CraftingData)

	p.ClearRecipes = input.ClearRecipes

	p.Recipes = ConvertSlice(
		input.Recipes,
		pk.ToNetEaseRecipe,
	)
	p.PotionRecipes = ConvertSlice(
		input.PotionRecipes,
		func(from standardProtocol.PotionRecipe) neteaseProtocol.PotionRecipe {
			return neteaseProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from standardProtocol.PotionContainerChangeRecipe) neteaseProtocol.PotionContainerChangeRecipe {
			return neteaseProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = ConvertSlice(
		input.MaterialReducers,
		func(from standardProtocol.MaterialReducer) neteaseProtocol.MaterialReducer {
			return neteaseProtocol.MaterialReducer{
				InputItem: neteaseProtocol.ItemType(from.InputItem),
				Outputs: ConvertSlice(
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

	p.Recipes = ConvertSlice(
		input.Recipes,
		pk.ToStandardRecipe,
	)
	p.PotionRecipes = ConvertSlice(
		input.PotionRecipes,
		func(from neteaseProtocol.PotionRecipe) standardProtocol.PotionRecipe {
			return standardProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from neteaseProtocol.PotionContainerChangeRecipe) standardProtocol.PotionContainerChangeRecipe {
			return standardProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = ConvertSlice(
		input.MaterialReducers,
		func(from neteaseProtocol.MaterialReducer) standardProtocol.MaterialReducer {
			return standardProtocol.MaterialReducer{
				InputItem: standardProtocol.ItemType(from.InputItem),
				Outputs: ConvertSlice(
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
