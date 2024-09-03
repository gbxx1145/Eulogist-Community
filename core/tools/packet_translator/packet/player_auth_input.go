package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PlayerAuthInput struct{}

// 将 from 转换为 neteaseProtocol.StackRequestAction
func (pk *PlayerAuthInput) ToNetEaseStackRequestAction(
	from standardProtocol.StackRequestAction,
) neteaseProtocol.StackRequestAction {
	switch data := from.(type) {
	case *standardProtocol.TakeStackRequestAction:
		action := neteaseProtocol.TakeStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.PlaceStackRequestAction:
		action := neteaseProtocol.PlaceStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.SwapStackRequestAction:
		return &neteaseProtocol.SwapStackRequestAction{
			Source:      neteaseProtocol.StackRequestSlotInfo(data.Source),
			Destination: neteaseProtocol.StackRequestSlotInfo(data.Destination),
		}
	case *standardProtocol.DropStackRequestAction:
		return &neteaseProtocol.DropStackRequestAction{
			Count:    data.Count,
			Source:   neteaseProtocol.StackRequestSlotInfo(data.Source),
			Randomly: data.Randomly,
		}
	case *standardProtocol.DestroyStackRequestAction:
		return &neteaseProtocol.DestroyStackRequestAction{
			Count:  data.Count,
			Source: neteaseProtocol.StackRequestSlotInfo(data.Source),
		}
	case *standardProtocol.ConsumeStackRequestAction:
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
		action := neteaseProtocol.PlaceInContainerStackRequestAction{}
		action.Count = data.Count
		action.Source = neteaseProtocol.StackRequestSlotInfo(data.Source)
		action.Destination = neteaseProtocol.StackRequestSlotInfo(data.Destination)
		return &action
	case *standardProtocol.TakeOutContainerStackRequestAction:
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
			Ingredients: packet_translator.ConvertSlice(
				data.Ingredients,
				func(from standardProtocol.ItemDescriptorCount) neteaseProtocol.ItemDescriptorCount {
					return packet_translator.ToNetEaseItemDescriptorCount(from)
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
			ResultItems: packet_translator.ConvertSlice(
				data.ResultItems,
				func(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
					return packet_translator.ConvertToNetEaseItemStack(from)
				},
			),
			TimesCrafted: data.TimesCrafted,
		}
	}

	panic("ToNetNetEaseStackRequestAction: Invalid stack request action enum")
}

// 将 from 转换为 neteaseProtocol.StackRequestAction
func (pk *PlayerAuthInput) ToStandardStackRequestAction(
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
			Ingredients: packet_translator.ConvertSlice(
				data.Ingredients,
				func(from neteaseProtocol.ItemDescriptorCount) standardProtocol.ItemDescriptorCount {
					return packet_translator.ToStandardItemDescriptorCount(from)
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
			ResultItems: packet_translator.ConvertSlice(
				data.ResultItems,
				func(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
					return packet_translator.ConvertToStandardItemStack(from)
				},
			),
			TimesCrafted: data.TimesCrafted,
		}
	}

	panic("ToStandardStackRequestAction: Invalid stack request action enum")
}

func (pk *PlayerAuthInput) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.PlayerAuthInput{}
	input := standard.(*standardPacket.PlayerAuthInput)

	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.Position = input.Position
	p.MoveVector = input.MoveVector
	p.HeadYaw = input.HeadYaw
	p.InputData = input.InputData
	p.InputMode = input.InputMode
	p.PlayMode = input.PlayMode
	p.GazeDirection = input.GazeDirection
	p.Tick = input.Tick
	p.Delta = input.Delta
	p.ItemInteractionData = *packet_translator.ConvertToNetEaseUseItemTransactionData(&input.ItemInteractionData)
	p.BlockActions = packet_translator.ConvertSlice(
		input.BlockActions,
		func(from standardProtocol.PlayerBlockAction) neteaseProtocol.PlayerBlockAction {
			return neteaseProtocol.PlayerBlockAction{
				Action:   from.Action,
				BlockPos: neteaseProtocol.BlockPos(from.BlockPos),
				Face:     from.Face,
			}
		},
	)
	p.AnalogueMoveVector = input.AnalogueMoveVector
	p.InteractionModel = uint32(input.InteractionModel)

	p.ItemStackRequest = neteaseProtocol.ItemStackRequest{
		RequestID: input.ItemStackRequest.RequestID,
		Actions: packet_translator.ConvertSlice(
			input.ItemStackRequest.Actions,
			pk.ToNetEaseStackRequestAction,
		),
		FilterStrings: input.ItemStackRequest.FilterStrings,
		FilterCause:   input.ItemStackRequest.FilterCause,
	}

	p.PitchRepeat = p.Pitch
	p.YawRepeat = p.Yaw
	p.IsFlying = false
	p.IsOnGround = false
	p.Unknown1 = false

	return &p
}

func (pk *PlayerAuthInput) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.PlayerAuthInput{}
	input := netease.(*neteasePacket.PlayerAuthInput)

	p.Pitch = input.Pitch
	p.Yaw = input.Yaw
	p.Position = input.Position
	p.MoveVector = input.MoveVector
	p.HeadYaw = input.HeadYaw
	p.InputData = input.InputData
	p.InputMode = input.InputMode
	p.PlayMode = input.PlayMode
	p.GazeDirection = input.GazeDirection
	p.Tick = input.Tick
	p.Delta = input.Delta
	p.ItemInteractionData = *packet_translator.ConvertToStandardUseItemTransactionData(&input.ItemInteractionData)
	p.BlockActions = packet_translator.ConvertSlice(
		input.BlockActions,
		func(from neteaseProtocol.PlayerBlockAction) standardProtocol.PlayerBlockAction {
			return standardProtocol.PlayerBlockAction{
				Action:   from.Action,
				BlockPos: standardProtocol.BlockPos(from.BlockPos),
				Face:     from.Face,
			}
		},
	)
	p.AnalogueMoveVector = input.AnalogueMoveVector
	p.InteractionModel = int32(input.InteractionModel)

	p.ItemStackRequest = standardProtocol.ItemStackRequest{
		RequestID: input.ItemStackRequest.RequestID,
		Actions: packet_translator.ConvertSlice(
			input.ItemStackRequest.Actions,
			pk.ToStandardStackRequestAction,
		),
		FilterStrings: input.ItemStackRequest.FilterStrings,
		FilterCause:   input.ItemStackRequest.FilterCause,
	}

	return &p
}
