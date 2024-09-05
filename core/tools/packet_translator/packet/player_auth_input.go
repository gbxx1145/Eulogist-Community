package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type PlayerAuthInput struct{}

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
	p.ItemInteractionData = *packet_translate_struct.ConvertToNetEaseUseItemTransactionData(&input.ItemInteractionData)
	p.BlockActions = packet_translate_struct.ConvertSlice(
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
		Actions: packet_translate_struct.ConvertSlice(
			input.ItemStackRequest.Actions,
			packet_translate_struct.ToNetEaseStackRequestAction,
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
	p.ItemInteractionData = *packet_translate_struct.ConvertToStandardUseItemTransactionData(&input.ItemInteractionData)
	p.BlockActions = packet_translate_struct.ConvertSlice(
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
		Actions: packet_translate_struct.ConvertSlice(
			input.ItemStackRequest.Actions,
			packet_translate_struct.ToStandardStackRequestAction,
		),
		FilterStrings: input.ItemStackRequest.FilterStrings,
		FilterCause:   input.ItemStackRequest.FilterCause,
	}

	return &p
}
