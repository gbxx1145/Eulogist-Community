package packet_translate_struct

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
)

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
