package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type CraftingData struct{}

// 将 standard 转换为 neteaseProtocol.ItemDescriptorCount
func (pk *CraftingData) ToNetEaseItemDescriptorCount(
	standard standardProtocol.ItemDescriptorCount,
) neteaseProtocol.ItemDescriptorCount {
	switch d := standard.Descriptor.(type) {
	case *standardProtocol.InvalidItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.InvalidItemDescriptor{},
			Count:      standard.Count,
		}
	case *standardProtocol.DefaultItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.DefaultItemDescriptor{
				NetworkID:     d.NetworkID,
				MetadataValue: d.MetadataValue,
			},
			Count: standard.Count,
		}
	case *standardProtocol.MoLangItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.MoLangItemDescriptor{
				Expression: d.Expression,
				Version:    d.Version,
			},
			Count: standard.Count,
		}
	case *standardProtocol.ItemTagItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.ItemTagItemDescriptor{
				Tag: d.Tag,
			},
			Count: standard.Count,
		}
	case *standardProtocol.DeferredItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.DeferredItemDescriptor{
				Name:          d.Name,
				MetadataValue: d.MetadataValue,
			},
			Count: standard.Count,
		}
	case *standardProtocol.ComplexAliasItemDescriptor:
		return neteaseProtocol.ItemDescriptorCount{
			Descriptor: &neteaseProtocol.ComplexAliasItemDescriptor{
				Name: d.Name,
			},
			Count: standard.Count,
		}
	}

	return neteaseProtocol.ItemDescriptorCount{
		Count: standard.Count,
	}
}

// 将 netease 转换为 standardProtocol.ItemDescriptorCount
func (pk *CraftingData) ToStandardItemDescriptorCount(
	netease neteaseProtocol.ItemDescriptorCount,
) standardProtocol.ItemDescriptorCount {
	switch d := netease.Descriptor.(type) {
	case *neteaseProtocol.InvalidItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.InvalidItemDescriptor{},
			Count:      netease.Count,
		}
	case *neteaseProtocol.DefaultItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.DefaultItemDescriptor{
				NetworkID:     d.NetworkID,
				MetadataValue: d.MetadataValue,
			},
			Count: netease.Count,
		}
	case *neteaseProtocol.MoLangItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.MoLangItemDescriptor{
				Expression: d.Expression,
				Version:    d.Version,
			},
			Count: netease.Count,
		}
	case *neteaseProtocol.ItemTagItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.ItemTagItemDescriptor{
				Tag: d.Tag,
			},
			Count: netease.Count,
		}
	case *neteaseProtocol.DeferredItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.DeferredItemDescriptor{
				Name:          d.Name,
				MetadataValue: d.MetadataValue,
			},
			Count: netease.Count,
		}
	case *neteaseProtocol.ComplexAliasItemDescriptor:
		return standardProtocol.ItemDescriptorCount{
			Descriptor: &standardProtocol.ComplexAliasItemDescriptor{
				Name: d.Name,
			},
			Count: netease.Count,
		}
	}

	return standardProtocol.ItemDescriptorCount{
		Count: netease.Count,
	}
}

// 将 standard 转换为 neteaseProtocol.Recipe
func (pk *CraftingData) ToNetEaseRecipe(
	standard standardProtocol.Recipe,
) neteaseProtocol.Recipe {
	switch data := standard.(type) {
	case *standardProtocol.ShapelessRecipe:
		return &neteaseProtocol.ShapelessRecipe{
			RecipeID: data.RecipeID,
			Input: packet_translator.ConvertSlice(
				data.Input,
				pk.ToNetEaseItemDescriptorCount,
			),
		}
	case *standardProtocol.ShapedRecipe:
		return &neteaseProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: packet_translator.ConvertSlice(
				data.Input,
				pk.ToNetEaseItemDescriptorCount,
			),
			Output: packet_translator.ConvertSlice(
				data.Output,
				func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
					return packet_translator.ConvertToNetEaseItemStack(from)
				},
			),
		}
	case *standardProtocol.FurnaceRecipe:
		return &neteaseProtocol.FurnaceRecipe{
			InputType: neteaseProtocol.ItemType(data.InputType),
			Output:    packet_translator.ConvertToNetEaseItemStack(data.Output),
			Block:     data.Block,
		}
	case *standardProtocol.FurnaceDataRecipe:
		return &neteaseProtocol.FurnaceDataRecipe{
			FurnaceRecipe: neteaseProtocol.FurnaceRecipe{
				InputType: neteaseProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    packet_translator.ConvertToNetEaseItemStack(data.FurnaceRecipe.Output),
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
				Input: packet_translator.ConvertSlice(
					data.ShapelessRecipe.Input,
					pk.ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapelessChemistryRecipe:
		return &neteaseProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: neteaseProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translator.ConvertSlice(
					data.ShapelessRecipe.Input,
					pk.ToNetEaseItemDescriptorCount,
				),
			},
		}
	case *standardProtocol.ShapedChemistryRecipe:
		return &neteaseProtocol.ShapedChemistryRecipe{
			ShapedRecipe: neteaseProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: packet_translator.ConvertSlice(
					data.ShapedRecipe.Input,
					pk.ToNetEaseItemDescriptorCount,
				),
				Output: packet_translator.ConvertSlice(
					data.ShapedRecipe.Output,
					func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
						return packet_translator.ConvertToNetEaseItemStack(from)
					},
				),
			},
		}
	case *standardProtocol.SmithingTransformRecipe:
		return &neteaseProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        pk.ToNetEaseItemDescriptorCount(data.Template),
			Base:            pk.ToNetEaseItemDescriptorCount(data.Base),
			Addition:        pk.ToNetEaseItemDescriptorCount(data.Addition),
			Result:          packet_translator.ConvertToNetEaseItemStack(data.Result),
			Block:           data.Block,
		}
	case *standardProtocol.SmithingTrimRecipe:
		return &neteaseProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        pk.ToNetEaseItemDescriptorCount(data.Template),
			Base:            pk.ToNetEaseItemDescriptorCount(data.Base),
			Addition:        pk.ToNetEaseItemDescriptorCount(data.Addition),
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
			Input: packet_translator.ConvertSlice(
				data.Input,
				pk.ToStandardItemDescriptorCount,
			),
		}
	case *neteaseProtocol.ShapedRecipe:
		return &standardProtocol.ShapedRecipe{
			RecipeID: data.RecipeID,
			Width:    data.Width,
			Height:   data.Height,
			Input: packet_translator.ConvertSlice(
				data.Input,
				pk.ToStandardItemDescriptorCount,
			),
			Output: packet_translator.ConvertSlice(
				data.Output,
				func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
					return packet_translator.ConvertToStandardItemStack(from)
				},
			),
		}
	case *neteaseProtocol.FurnaceRecipe:
		return &standardProtocol.FurnaceRecipe{
			InputType: standardProtocol.ItemType(data.InputType),
			Output:    packet_translator.ConvertToStandardItemStack(data.Output),
			Block:     data.Block,
		}
	case *neteaseProtocol.FurnaceDataRecipe:
		return &standardProtocol.FurnaceDataRecipe{
			FurnaceRecipe: standardProtocol.FurnaceRecipe{
				InputType: standardProtocol.ItemType(data.FurnaceRecipe.InputType),
				Output:    packet_translator.ConvertToStandardItemStack(data.FurnaceRecipe.Output),
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
				Input: packet_translator.ConvertSlice(
					data.ShapelessRecipe.Input,
					pk.ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapelessChemistryRecipe:
		return &standardProtocol.ShapelessChemistryRecipe{
			ShapelessRecipe: standardProtocol.ShapelessRecipe{
				RecipeID: data.ShapelessRecipe.RecipeID,
				Input: packet_translator.ConvertSlice(
					data.ShapelessRecipe.Input,
					pk.ToStandardItemDescriptorCount,
				),
			},
		}
	case *neteaseProtocol.ShapedChemistryRecipe:
		return &standardProtocol.ShapedChemistryRecipe{
			ShapedRecipe: standardProtocol.ShapedRecipe{
				RecipeID: data.ShapedRecipe.RecipeID,
				Width:    data.ShapedRecipe.Width,
				Height:   data.ShapedRecipe.Height,
				Input: packet_translator.ConvertSlice(
					data.ShapedRecipe.Input,
					pk.ToStandardItemDescriptorCount,
				),
				Output: packet_translator.ConvertSlice(
					data.ShapedRecipe.Output,
					func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
						return packet_translator.ConvertToStandardItemStack(from)
					},
				),
			},
		}
	case *neteaseProtocol.SmithingTransformRecipe:
		return &standardProtocol.SmithingTransformRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        pk.ToStandardItemDescriptorCount(data.Template),
			Base:            pk.ToStandardItemDescriptorCount(data.Base),
			Addition:        pk.ToStandardItemDescriptorCount(data.Addition),
			Result:          packet_translator.ConvertToStandardItemStack(data.Result),
			Block:           data.Block,
		}
	case *neteaseProtocol.SmithingTrimRecipe:
		return &standardProtocol.SmithingTrimRecipe{
			RecipeNetworkID: data.RecipeNetworkID,
			RecipeID:        data.RecipeID,
			Template:        pk.ToStandardItemDescriptorCount(data.Template),
			Base:            pk.ToStandardItemDescriptorCount(data.Base),
			Addition:        pk.ToStandardItemDescriptorCount(data.Addition),
			Block:           data.Block,
		}
	}

	panic("ToStandardRecipe: Invalid recipe enum")
}

func (pk *CraftingData) ToNetNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.CraftingData{}
	input := standard.(*standardPacket.CraftingData)

	p.ClearRecipes = input.ClearRecipes

	p.Recipes = packet_translator.ConvertSlice(
		input.Recipes,
		pk.ToNetEaseRecipe,
	)
	p.PotionRecipes = packet_translator.ConvertSlice(
		input.PotionRecipes,
		func(from standardProtocol.PotionRecipe) neteaseProtocol.PotionRecipe {
			return neteaseProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = packet_translator.ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from standardProtocol.PotionContainerChangeRecipe) neteaseProtocol.PotionContainerChangeRecipe {
			return neteaseProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = packet_translator.ConvertSlice(
		input.MaterialReducers,
		func(from standardProtocol.MaterialReducer) neteaseProtocol.MaterialReducer {
			return neteaseProtocol.MaterialReducer{
				InputItem: neteaseProtocol.ItemType(from.InputItem),
				Outputs: packet_translator.ConvertSlice(
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

	p.Recipes = packet_translator.ConvertSlice(
		input.Recipes,
		pk.ToStandardRecipe,
	)
	p.PotionRecipes = packet_translator.ConvertSlice(
		input.PotionRecipes,
		func(from neteaseProtocol.PotionRecipe) standardProtocol.PotionRecipe {
			return standardProtocol.PotionRecipe(from)
		},
	)
	p.PotionContainerChangeRecipes = packet_translator.ConvertSlice(
		input.PotionContainerChangeRecipes,
		func(from neteaseProtocol.PotionContainerChangeRecipe) standardProtocol.PotionContainerChangeRecipe {
			return standardProtocol.PotionContainerChangeRecipe(from)
		},
	)
	p.MaterialReducers = packet_translator.ConvertSlice(
		input.MaterialReducers,
		func(from neteaseProtocol.MaterialReducer) standardProtocol.MaterialReducer {
			return standardProtocol.MaterialReducer{
				InputItem: standardProtocol.ItemType(from.InputItem),
				Outputs: packet_translator.ConvertSlice(
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
