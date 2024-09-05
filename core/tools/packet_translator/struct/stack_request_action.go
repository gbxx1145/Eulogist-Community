package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	packet_translate_pool "Eulogist/core/tools/packet_translator/pool"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
)

// 将 from 转换为 neteaseProtocol.StackRequestAction
func ToNetEaseStackRequestAction(
	from standardProtocol.StackRequestAction,
) neteaseProtocol.StackRequestAction {
	switch data := from.(type) {
	case *standardProtocol.TakeStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		data.Destination.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Destination.ContainerID]
		action := neteaseProtocol.TakeStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.PlaceStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		data.Destination.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Destination.ContainerID]
		action := neteaseProtocol.PlaceStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.SwapStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		data.Destination.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Destination.ContainerID]
		return &neteaseProtocol.SwapStackRequestAction{
			Source:      neteaseProtocol.StackRequestSlotInfo(data.Source),
			Destination: neteaseProtocol.StackRequestSlotInfo(data.Destination),
		}
	case *standardProtocol.DropStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		return &neteaseProtocol.DropStackRequestAction{
			Count:    data.Count,
			Source:   neteaseProtocol.StackRequestSlotInfo(data.Source),
			Randomly: data.Randomly,
		}
	case *standardProtocol.DestroyStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		return &neteaseProtocol.DestroyStackRequestAction{
			Count:  data.Count,
			Source: neteaseProtocol.StackRequestSlotInfo(data.Source),
		}
	case *standardProtocol.ConsumeStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		return &neteaseProtocol.ConsumeStackRequestAction{
			DestroyStackRequestAction: neteaseProtocol.DestroyStackRequestAction{
				Count:  data.DestroyStackRequestAction.Count,
				Source: neteaseProtocol.StackRequestSlotInfo(data.DestroyStackRequestAction.Source),
			},
		}
	case *standardProtocol.CreateStackRequestAction:
		return &neteaseProtocol.CreateStackRequestAction{
			ResultsSlot: data.ResultsSlot,
		}
	case *standardProtocol.PlaceInContainerStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		data.Destination.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Destination.ContainerID]
		action := neteaseProtocol.PlaceInContainerStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.TakeOutContainerStackRequestAction:
		data.Source.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Source.ContainerID]
		data.Destination.ContainerID = packet_translate_pool.StandardContainerIDToNetEaseContainerID[data.Destination.ContainerID]
		action := neteaseProtocol.TakeOutContainerStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.LabTableCombineStackRequestAction:
		return &neteaseProtocol.LabTableCombineStackRequestAction{}
	case *standardProtocol.BeaconPaymentStackRequestAction:
		return &neteaseProtocol.BeaconPaymentStackRequestAction{
			PrimaryEffect:   data.PrimaryEffect,
			SecondaryEffect: data.SecondaryEffect,
		}
	case *standardProtocol.MineBlockStackRequestAction:
		return &neteaseProtocol.MineBlockStackRequestAction{
			HotbarSlot:          data.HotbarSlot,
			PredictedDurability: data.PredictedDurability,
			StackNetworkID:      data.StackNetworkID,
		}
	case *standardProtocol.CraftRecipeStackRequestAction:
		return &neteaseProtocol.CraftRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
		}
	case *standardProtocol.AutoCraftRecipeStackRequestAction:
		return &neteaseProtocol.AutoCraftRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
			TimesCrafted:    data.TimesCrafted,
			Ingredients: ConvertSlice(
				data.Ingredients,
				func(from standardProtocol.ItemDescriptorCount) neteaseProtocol.ItemDescriptorCount {
					return ToNetEaseItemDescriptorCount(from)
				},
			),
		}
	case *standardProtocol.CraftCreativeStackRequestAction:
		return &neteaseProtocol.CraftCreativeStackRequestAction{
			CreativeItemNetworkID: data.CreativeItemNetworkID,
		}
	case *standardProtocol.CraftRecipeOptionalStackRequestAction:
		return &neteaseProtocol.CraftRecipeOptionalStackRequestAction{
			RecipeNetworkID:   data.RecipeNetworkID,
			FilterStringIndex: data.FilterStringIndex,
		}
	case *standardProtocol.CraftGrindstoneRecipeStackRequestAction:
		return &neteaseProtocol.CraftGrindstoneRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
			Cost:            data.Cost,
		}
	case *standardProtocol.CraftLoomRecipeStackRequestAction:
		return &neteaseProtocol.CraftLoomRecipeStackRequestAction{
			Pattern: data.Pattern,
		}
	case *standardProtocol.CraftNonImplementedStackRequestAction:
		return &neteaseProtocol.CraftNonImplementedStackRequestAction{}
	case *standardProtocol.CraftResultsDeprecatedStackRequestAction:
		return &neteaseProtocol.CraftResultsDeprecatedStackRequestAction{
			ResultItems: ConvertSlice(
				data.ResultItems,
				func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
					return ConvertToNetEaseItemStack(from)
				},
			),
			TimesCrafted: data.TimesCrafted,
		}
	}

	panic("ToNetNetEaseStackRequestAction: Invalid stack request action enum")
}

// 将 from 转换为 neteaseProtocol.StackRequestAction
func ToStandardStackRequestAction(
	from neteaseProtocol.StackRequestAction,
) standardProtocol.StackRequestAction {
	switch data := from.(type) {
	case *neteaseProtocol.TakeStackRequestAction:
		action := standardProtocol.TakeStackRequestAction{}
		action.Count = data.Count
		action.Source = standardProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = standardProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *neteaseProtocol.PlaceStackRequestAction:
		action := standardProtocol.PlaceStackRequestAction{}
		action.Count = data.Count
		action.Source = standardProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = standardProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *neteaseProtocol.SwapStackRequestAction:
		return &standardProtocol.SwapStackRequestAction{
			Source:      standardProtocol.StackRequestSlotInfo(data.Source),
			Destination: standardProtocol.StackRequestSlotInfo(data.Destination),
		}
	case *neteaseProtocol.DropStackRequestAction:
		return &standardProtocol.DropStackRequestAction{
			Count:    data.Count,
			Source:   standardProtocol.StackRequestSlotInfo(data.Source),
			Randomly: data.Randomly,
		}
	case *neteaseProtocol.DestroyStackRequestAction:
		return &standardProtocol.DestroyStackRequestAction{
			Count:  data.Count,
			Source: standardProtocol.StackRequestSlotInfo(data.Source),
		}
	case *neteaseProtocol.ConsumeStackRequestAction:
		return &standardProtocol.ConsumeStackRequestAction{
			DestroyStackRequestAction: standardProtocol.DestroyStackRequestAction{
				Count:  data.DestroyStackRequestAction.Count,
				Source: standardProtocol.StackRequestSlotInfo(data.DestroyStackRequestAction.Source),
			},
		}
	case *neteaseProtocol.CreateStackRequestAction:
		return &standardProtocol.CreateStackRequestAction{
			ResultsSlot: data.ResultsSlot,
		}
	case *neteaseProtocol.PlaceInContainerStackRequestAction:
		action := standardProtocol.PlaceInContainerStackRequestAction{}
		action.Count = data.Count
		action.Source = standardProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = standardProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *neteaseProtocol.TakeOutContainerStackRequestAction:
		action := standardProtocol.TakeOutContainerStackRequestAction{}
		action.Count = data.Count
		action.Source = standardProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = standardProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *neteaseProtocol.LabTableCombineStackRequestAction:
		return &standardProtocol.LabTableCombineStackRequestAction{}
	case *neteaseProtocol.BeaconPaymentStackRequestAction:
		return &standardProtocol.BeaconPaymentStackRequestAction{
			PrimaryEffect:   data.PrimaryEffect,
			SecondaryEffect: data.SecondaryEffect,
		}
	case *neteaseProtocol.MineBlockStackRequestAction:
		return &standardProtocol.MineBlockStackRequestAction{
			HotbarSlot:          data.HotbarSlot,
			PredictedDurability: data.PredictedDurability,
			StackNetworkID:      data.StackNetworkID,
		}
	case *neteaseProtocol.CraftRecipeStackRequestAction:
		return &standardProtocol.CraftRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
		}
	case *neteaseProtocol.AutoCraftRecipeStackRequestAction:
		return &standardProtocol.AutoCraftRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
			TimesCrafted:    data.TimesCrafted,
			Ingredients: ConvertSlice(
				data.Ingredients,
				func(from neteaseProtocol.ItemDescriptorCount) standardProtocol.ItemDescriptorCount {
					return ToStandardItemDescriptorCount(from)
				},
			),
		}
	case *neteaseProtocol.CraftCreativeStackRequestAction:
		return &standardProtocol.CraftCreativeStackRequestAction{
			CreativeItemNetworkID: data.CreativeItemNetworkID,
		}
	case *neteaseProtocol.CraftRecipeOptionalStackRequestAction:
		return &standardProtocol.CraftRecipeOptionalStackRequestAction{
			RecipeNetworkID:   data.RecipeNetworkID,
			FilterStringIndex: data.FilterStringIndex,
		}
	case *neteaseProtocol.CraftGrindstoneRecipeStackRequestAction:
		return &standardProtocol.CraftGrindstoneRecipeStackRequestAction{
			RecipeNetworkID: data.RecipeNetworkID,
			Cost:            data.Cost,
		}
	case *neteaseProtocol.CraftLoomRecipeStackRequestAction:
		return &standardProtocol.CraftLoomRecipeStackRequestAction{
			Pattern: data.Pattern,
		}
	case *neteaseProtocol.CraftNonImplementedStackRequestAction:
		return &standardProtocol.CraftNonImplementedStackRequestAction{}
	case *neteaseProtocol.CraftResultsDeprecatedStackRequestAction:
		return &standardProtocol.CraftResultsDeprecatedStackRequestAction{
			ResultItems: ConvertSlice(
				data.ResultItems,
				func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
					return ConvertToStandardItemStack(from)
				},
			),
			TimesCrafted: data.TimesCrafted,
		}
	}

	panic("ToStandardStackRequestAction: Invalid stack request action enum")
}
