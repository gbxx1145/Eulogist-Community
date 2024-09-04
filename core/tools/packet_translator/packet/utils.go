package packet

import (
	neteaseProtocol "Eulogist/core/minecraft/protocol"

	standardProtocol "github.com/sandertv/gophertunnel/minecraft/protocol"
)

// 将切片 from([]From) 转换为 []To。
// converter 是用于转换切片内的单个元素的函数
func ConvertSlice[From any, To any](
	from []From,
	converter func(from From) To,
) []To {
	to := make([]To, len(from))
	for index, value := range from {
		to[index] = converter(value)
	}
	return to
}

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.ItemStack
func ConvertToNetEaseItemStack(from standardProtocol.ItemStack) neteaseProtocol.ItemStack {
	return neteaseProtocol.ItemStack{
		ItemType:       neteaseProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}
}

// 将 from 转换为 standardProtocol.ItemStack
func ConvertToStandardItemStack(from neteaseProtocol.ItemStack) standardProtocol.ItemStack {
	return standardProtocol.ItemStack{
		ItemType:       standardProtocol.ItemType(from.ItemType),
		BlockRuntimeID: from.BlockRuntimeID,
		Count:          from.Count,
		NBTData:        from.NBTData,
		CanBePlacedOn:  from.CanBePlacedOn,
		CanBreak:       from.CanBreak,
		HasNetworkID:   from.HasNetworkID,
	}
}

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.ItemInstance
func ConvertToNetEaseItemInstance(from standardProtocol.ItemInstance) neteaseProtocol.ItemInstance {
	return neteaseProtocol.ItemInstance{
		StackNetworkID: from.StackNetworkID,
		Stack:          ConvertToNetEaseItemStack(from.Stack),
	}
}

// 将 from 转换为 standardProtocol.ItemInstance
func ConvertToStandardItemInstance(from neteaseProtocol.ItemInstance) standardProtocol.ItemInstance {
	return standardProtocol.ItemInstance{
		StackNetworkID: from.StackNetworkID,
		Stack:          ConvertToStandardItemStack(from.Stack),
	}
}

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.InventoryAction
func ConvertToNetEaseInventoryAction(from standardProtocol.InventoryAction) neteaseProtocol.InventoryAction {
	return neteaseProtocol.InventoryAction{
		SourceType:    from.SourceType,
		WindowID:      from.WindowID,
		SourceFlags:   from.SourceFlags,
		InventorySlot: from.InventorySlot,
		OldItem:       ConvertToNetEaseItemInstance(from.OldItem),
		NewItem:       ConvertToNetEaseItemInstance(from.NewItem),
	}
}

// 将 from 转换为 standardProtocol.InventoryAction
func ConvertToStandardInventoryAction(from neteaseProtocol.InventoryAction) standardProtocol.InventoryAction {
	return standardProtocol.InventoryAction{
		SourceType:    from.SourceType,
		WindowID:      from.WindowID,
		SourceFlags:   from.SourceFlags,
		InventorySlot: from.InventorySlot,
		OldItem:       ConvertToStandardItemInstance(from.OldItem),
		NewItem:       ConvertToStandardItemInstance(from.NewItem),
	}
}

// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.UseItemTransactionData
func ConvertToNetEaseUseItemTransactionData(from *standardProtocol.UseItemTransactionData) *neteaseProtocol.UseItemTransactionData {
	return &neteaseProtocol.UseItemTransactionData{
		LegacyRequestID: from.LegacyRequestID,
		LegacySetItemSlots: ConvertSlice(
			from.LegacySetItemSlots,
			func(from standardProtocol.LegacySetItemSlot) neteaseProtocol.LegacySetItemSlot {
				return neteaseProtocol.LegacySetItemSlot(from)
			},
		),
		Actions: ConvertSlice(
			from.Actions,
			ConvertToNetEaseInventoryAction,
		),
		ActionType:      from.ActionType,
		BlockPosition:   neteaseProtocol.BlockPos(from.BlockPosition),
		BlockFace:       from.BlockFace,
		HotBarSlot:      from.HotBarSlot,
		HeldItem:        ConvertToNetEaseItemInstance(from.HeldItem),
		Position:        from.Position,
		ClickedPosition: from.ClickedPosition,
		BlockRuntimeID:  from.BlockRuntimeID,
	}
}

// 将 from 转换为 standardProtocol.UseItemTransactionData
func ConvertToStandardUseItemTransactionData(from *neteaseProtocol.UseItemTransactionData) *standardProtocol.UseItemTransactionData {
	return &standardProtocol.UseItemTransactionData{
		LegacyRequestID: from.LegacyRequestID,
		LegacySetItemSlots: ConvertSlice(
			from.LegacySetItemSlots,
			func(from neteaseProtocol.LegacySetItemSlot) standardProtocol.LegacySetItemSlot {
				return standardProtocol.LegacySetItemSlot(from)
			},
		),
		Actions: ConvertSlice(
			from.Actions,
			ConvertToStandardInventoryAction,
		),
		ActionType:      from.ActionType,
		BlockPosition:   standardProtocol.BlockPos(from.BlockPosition),
		BlockFace:       from.BlockFace,
		HotBarSlot:      from.HotBarSlot,
		HeldItem:        ConvertToStandardItemInstance(from.HeldItem),
		Position:        from.Position,
		ClickedPosition: from.ClickedPosition,
		BlockRuntimeID:  from.BlockRuntimeID,
	}
}

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.ItemDescriptor
func ToNetEaseItemDescriptor(
	from standardProtocol.ItemDescriptor,
) neteaseProtocol.ItemDescriptor {
	switch d := from.(type) {
	case *standardProtocol.InvalidItemDescriptor:
		return &neteaseProtocol.InvalidItemDescriptor{}
	case *standardProtocol.DefaultItemDescriptor:
		return &neteaseProtocol.DefaultItemDescriptor{
			NetworkID:     d.NetworkID,
			MetadataValue: d.MetadataValue,
		}
	case *standardProtocol.MoLangItemDescriptor:
		return &neteaseProtocol.MoLangItemDescriptor{
			Expression: d.Expression,
			Version:    d.Version,
		}
	case *standardProtocol.ItemTagItemDescriptor:
		return &neteaseProtocol.ItemTagItemDescriptor{
			Tag: d.Tag,
		}
	case *standardProtocol.DeferredItemDescriptor:
		return &neteaseProtocol.DeferredItemDescriptor{
			Name:          d.Name,
			MetadataValue: d.MetadataValue,
		}
	case *standardProtocol.ComplexAliasItemDescriptor:
		return &neteaseProtocol.ComplexAliasItemDescriptor{
			Name: d.Name,
		}
	}

	return &neteaseProtocol.InvalidItemDescriptor{}
}

// 将 from 转换为 standardProtocol.ItemDescriptor
func ToStandardItemDescriptor(
	from neteaseProtocol.ItemDescriptor,
) standardProtocol.ItemDescriptor {
	switch d := from.(type) {
	case *neteaseProtocol.InvalidItemDescriptor:
		return &standardProtocol.InvalidItemDescriptor{}
	case *neteaseProtocol.DefaultItemDescriptor:
		return &standardProtocol.DefaultItemDescriptor{
			NetworkID:     d.NetworkID,
			MetadataValue: d.MetadataValue,
		}
	case *neteaseProtocol.MoLangItemDescriptor:
		return &standardProtocol.MoLangItemDescriptor{
			Expression: d.Expression,
			Version:    d.Version,
		}
	case *neteaseProtocol.ItemTagItemDescriptor:
		return &standardProtocol.ItemTagItemDescriptor{
			Tag: d.Tag,
		}
	case *neteaseProtocol.DeferredItemDescriptor:
		return &standardProtocol.DeferredItemDescriptor{
			Name:          d.Name,
			MetadataValue: d.MetadataValue,
		}
	case *neteaseProtocol.ComplexAliasItemDescriptor:
		return &standardProtocol.ComplexAliasItemDescriptor{
			Name: d.Name,
		}
	}

	return &standardProtocol.InvalidItemDescriptor{}
}

// ---------------------------------------------------------------------------

// 将 from 转换为 neteaseProtocol.ItemDescriptorCount
func ToNetEaseItemDescriptorCount(
	from standardProtocol.ItemDescriptorCount,
) neteaseProtocol.ItemDescriptorCount {
	return neteaseProtocol.ItemDescriptorCount{
		Descriptor: ToNetEaseItemDescriptor(from.Descriptor),
		Count:      from.Count,
	}
}

// 将 from 转换为 standardProtocol.ItemDescriptorCount
func ToStandardItemDescriptorCount(
	from neteaseProtocol.ItemDescriptorCount,
) standardProtocol.ItemDescriptorCount {
	return standardProtocol.ItemDescriptorCount{
		Descriptor: ToStandardItemDescriptor(from.Descriptor),
		Count:      from.Count,
	}
}
