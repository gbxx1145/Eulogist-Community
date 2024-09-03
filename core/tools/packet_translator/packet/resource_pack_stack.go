package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/tools/packet_translator"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type ResourcePackStack struct{}

func (pk *ResourcePackStack) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.ResourcePackStack{}
	input := standard.(*standardPacket.ResourcePackStack)

	p.TexturePackRequired = input.TexturePackRequired
	p.BaseGameVersion = input.BaseGameVersion
	p.ExperimentsPreviouslyToggled = input.ExperimentsPreviouslyToggled
	p.Unknown1 = false
	p.Unknown2 = false

	p.BehaviourPacks = packet_translator.ConvertSlice(
		input.BehaviourPacks,
		func(from standardProtocol.StackResourcePack) neteaseProtocol.StackResourcePack {
			return neteaseProtocol.StackResourcePack(from)
		},
	)
	p.TexturePacks = packet_translator.ConvertSlice(
		input.TexturePacks,
		func(from standardProtocol.StackResourcePack) neteaseProtocol.StackResourcePack {
			return neteaseProtocol.StackResourcePack(from)
		},
	)
	p.Experiments = packet_translator.ConvertSlice(
		input.Experiments,
		func(from standardProtocol.ExperimentData) neteaseProtocol.ExperimentData {
			return neteaseProtocol.ExperimentData(from)
		},
	)

	return &p
}

func (pk *ResourcePackStack) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.ResourcePackStack{}
	input := netease.(*neteasePacket.ResourcePackStack)

	p.TexturePackRequired = input.TexturePackRequired
	p.BaseGameVersion = input.BaseGameVersion
	p.ExperimentsPreviouslyToggled = input.ExperimentsPreviouslyToggled

	p.BehaviourPacks = packet_translator.ConvertSlice(
		input.BehaviourPacks,
		func(from neteaseProtocol.StackResourcePack) standardProtocol.StackResourcePack {
			return standardProtocol.StackResourcePack(from)
		},
	)
	p.TexturePacks = packet_translator.ConvertSlice(
		input.TexturePacks,
		func(from neteaseProtocol.StackResourcePack) standardProtocol.StackResourcePack {
			return standardProtocol.StackResourcePack(from)
		},
	)
	p.Experiments = packet_translator.ConvertSlice(
		input.Experiments,
		func(from neteaseProtocol.ExperimentData) standardProtocol.ExperimentData {
			return standardProtocol.ExperimentData(from)
		},
	)

	return &p
}
