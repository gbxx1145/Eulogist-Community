package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

type EducationSettings struct{}

func (pk *EducationSettings) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.EducationSettings{}
	input := standard.(*standardPacket.EducationSettings)

	p.CodeBuilderDefaultURI = input.CodeBuilderDefaultURI
	p.CodeBuilderTitle = input.CodeBuilderTitle
	p.CanResizeCodeBuilder = input.CanResizeCodeBuilder
	p.DisableLegacyTitleBar = input.DisableLegacyTitleBar
	p.PostProcessFilter = input.PostProcessFilter
	p.ScreenshotBorderPath = input.ScreenshotBorderPath
	p.HasQuiz = input.HasQuiz

	if value, has := input.CanModifyBlocks.Value(); has {
		p.CanModifyBlocks = neteaseProtocol.Option(value)
	}
	if value, has := input.OverrideURI.Value(); has {
		p.OverrideURI = neteaseProtocol.Option(value)
	}
	if value, has := input.ExternalLinkSettings.Value(); has {
		p.ExternalLinkSettings = neteaseProtocol.Option(neteaseProtocol.EducationExternalLinkSettings(value))
	}

	p.Unknown1 = false

	return &p
}

func (pk *EducationSettings) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.EducationSettings{}
	input := netease.(*neteasePacket.EducationSettings)

	p.CodeBuilderDefaultURI = input.CodeBuilderDefaultURI
	p.CodeBuilderTitle = input.CodeBuilderTitle
	p.CanResizeCodeBuilder = input.CanResizeCodeBuilder
	p.DisableLegacyTitleBar = input.DisableLegacyTitleBar
	p.PostProcessFilter = input.PostProcessFilter
	p.ScreenshotBorderPath = input.ScreenshotBorderPath
	p.HasQuiz = input.HasQuiz

	if value, has := input.CanModifyBlocks.Value(); has {
		p.CanModifyBlocks = standardProtocol.Option(value)
	}
	if value, has := input.OverrideURI.Value(); has {
		p.OverrideURI = standardProtocol.Option(value)
	}
	if value, has := input.ExternalLinkSettings.Value(); has {
		p.ExternalLinkSettings = standardProtocol.Option(standardProtocol.EducationExternalLinkSettings(value))
	}

	return &p
}
