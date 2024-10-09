package packet

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type AgentAction struct{}

func (pk *AgentAction) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	input := standard.(*standardPacket.AgentAction)
	p := neteasePacket.AgentAction(*input)

	return &p
}

func (pk *AgentAction) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	input := netease.(*neteasePacket.AgentAction)
	p := standardPacket.AgentAction(*input)

	return &p
}
