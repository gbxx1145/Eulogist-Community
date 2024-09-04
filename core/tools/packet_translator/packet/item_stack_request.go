package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"

	standardProtocol "Eulogist/core/standard/protocol"
	standardPacket "Eulogist/core/standard/protocol/packet"
)

type ItemStackRequest struct{}

func (pk *ItemStackRequest) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.ItemStackRequest{}
	input := standard.(*standardPacket.ItemStackRequest)

	p.Requests = packet_translate_struct.ConvertSlice(
		input.Requests,
		func(from standardProtocol.ItemStackRequest) neteaseProtocol.ItemStackRequest {
			return neteaseProtocol.ItemStackRequest{
				RequestID: from.RequestID,
				Actions: packet_translate_struct.ConvertSlice(
					from.Actions,
					func(from standardProtocol.StackRequestAction) neteaseProtocol.StackRequestAction {
						return packet_translate_struct.ToNetEaseStackRequestAction(from)
					},
				),
				FilterStrings: from.FilterStrings,
				FilterCause:   from.FilterCause,
			}
		},
	)

	return &p
}

func (pk *ItemStackRequest) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.ItemStackRequest{}
	input := netease.(*neteasePacket.ItemStackRequest)

	p.Requests = packet_translate_struct.ConvertSlice(
		input.Requests,
		func(from neteaseProtocol.ItemStackRequest) standardProtocol.ItemStackRequest {
			return standardProtocol.ItemStackRequest{
				RequestID: from.RequestID,
				Actions: packet_translate_struct.ConvertSlice(
					from.Actions,
					func(from neteaseProtocol.StackRequestAction) standardProtocol.StackRequestAction {
						return packet_translate_struct.ToStandardStackRequestAction(from)
					},
				),
				FilterStrings: from.FilterStrings,
				FilterCause:   from.FilterCause,
			}
		},
	)

	return &p
}
