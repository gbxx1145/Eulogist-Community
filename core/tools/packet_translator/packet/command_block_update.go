package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardProtocol "Eulogist/core/standard/protocol"
	standardPacket "Eulogist/core/standard/protocol/packet"
)

type CommandBlockUpdate struct{}

func (pk *CommandBlockUpdate) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.CommandBlockUpdate{}
	input := standard.(*standardPacket.CommandBlockUpdate)

	p.Block = input.Block
	p.Position = neteaseProtocol.BlockPos(input.Position)
	p.Mode = input.Mode
	p.NeedsRedstone = input.NeedsRedstone
	p.Conditional = input.Conditional
	p.MinecartEntityRuntimeID = input.MinecartEntityRuntimeID
	p.Command = input.Command
	p.LastOutput = input.LastOutput
	p.Name = input.Name
	p.ShouldTrackOutput = input.ShouldTrackOutput
	p.ExecuteOnFirstTick = input.ExecuteOnFirstTick
	p.TickDelay = uint32(input.TickDelay)

	return &p
}

func (pk *CommandBlockUpdate) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.CommandBlockUpdate{}
	input := netease.(*neteasePacket.CommandBlockUpdate)

	p.Block = input.Block
	p.Position = standardProtocol.BlockPos(input.Position)
	p.Mode = input.Mode
	p.NeedsRedstone = input.NeedsRedstone
	p.Conditional = input.Conditional
	p.MinecartEntityRuntimeID = input.MinecartEntityRuntimeID
	p.Command = input.Command
	p.LastOutput = input.LastOutput
	p.Name = input.Name
	p.ShouldTrackOutput = input.ShouldTrackOutput
	p.ExecuteOnFirstTick = input.ExecuteOnFirstTick
	p.TickDelay = int32(input.TickDelay)

	return &p
}
