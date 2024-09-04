package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"
	standardProtocol "Eulogist/core/standard/protocol"
)

// 将 from 转换为 neteaseProtocol.Skin
func ConvertToNetEaseSkin(from standardProtocol.Skin) neteaseProtocol.Skin {
	return neteaseProtocol.Skin{
		SkinID:            from.SkinID,
		PlayFabID:         from.PlayFabID,
		SkinResourcePatch: from.SkinResourcePatch,
		SkinImageWidth:    from.SkinImageWidth,
		SkinData:          from.SkinData,
		Animations: ConvertSlice(
			from.Animations,
			func(from standardProtocol.SkinAnimation) neteaseProtocol.SkinAnimation {
				return neteaseProtocol.SkinAnimation(from)
			},
		),
		CapeImageWidth:            from.CapeImageWidth,
		CapeImageHeight:           from.CapeImageHeight,
		CapeData:                  from.CapeData,
		SkinGeometry:              from.SkinGeometry,
		AnimationData:             from.AnimationData,
		GeometryDataEngineVersion: from.GeometryDataEngineVersion,
		PremiumSkin:               from.PremiumSkin,
		PersonaSkin:               from.PersonaSkin,
		PersonaCapeOnClassicSkin:  from.PersonaCapeOnClassicSkin,
		PrimaryUser:               from.PrimaryUser,
		CapeID:                    from.CapeID,
		FullID:                    from.FullID,
		SkinColour:                from.SkinColour,
		ArmSize:                   from.ArmSize,
		PersonaPieces: ConvertSlice(
			from.PersonaPieces,
			func(from standardProtocol.PersonaPiece) neteaseProtocol.PersonaPiece {
				return neteaseProtocol.PersonaPiece(from)
			},
		),
		PieceTintColours: ConvertSlice(
			from.PieceTintColours,
			func(from standardProtocol.PersonaPieceTintColour) neteaseProtocol.PersonaPieceTintColour {
				return neteaseProtocol.PersonaPieceTintColour(from)
			},
		),
		Trusted:            from.Trusted,
		OverrideAppearance: from.OverrideAppearance,
	}
}

// 将 from 转换为 standardProtocol.Skin
func ConvertToStandardSkin(from neteaseProtocol.Skin) standardProtocol.Skin {
	return standardProtocol.Skin{
		SkinID:            from.SkinID,
		PlayFabID:         from.PlayFabID,
		SkinResourcePatch: from.SkinResourcePatch,
		SkinImageWidth:    from.SkinImageWidth,
		SkinImageHeight:   from.SkinImageHeight,
		SkinData:          from.SkinData,
		Animations: ConvertSlice(
			from.Animations,
			func(from neteaseProtocol.SkinAnimation) standardProtocol.SkinAnimation {
				return standardProtocol.SkinAnimation(from)
			},
		),
		CapeImageWidth:            from.CapeImageWidth,
		CapeImageHeight:           from.CapeImageHeight,
		CapeData:                  from.CapeData,
		SkinGeometry:              from.SkinGeometry,
		AnimationData:             from.AnimationData,
		GeometryDataEngineVersion: from.GeometryDataEngineVersion,
		PremiumSkin:               from.PremiumSkin,
		PersonaSkin:               from.PersonaSkin,
		PersonaCapeOnClassicSkin:  from.PersonaCapeOnClassicSkin,
		PrimaryUser:               from.PrimaryUser,
		CapeID:                    from.CapeID,
		FullID:                    from.FullID,
		SkinColour:                from.SkinColour,
		ArmSize:                   from.ArmSize,
		PersonaPieces: ConvertSlice(
			from.PersonaPieces,
			func(from neteaseProtocol.PersonaPiece) standardProtocol.PersonaPiece {
				return standardProtocol.PersonaPiece(from)
			},
		),
		PieceTintColours: ConvertSlice(
			from.PieceTintColours,
			func(from neteaseProtocol.PersonaPieceTintColour) standardProtocol.PersonaPieceTintColour {
				return standardProtocol.PersonaPieceTintColour(from)
			},
		),
		Trusted:            from.Trusted,
		OverrideAppearance: from.OverrideAppearance,
	}
}
