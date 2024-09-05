package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
)

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
