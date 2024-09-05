package packet_translator

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

// 数据包翻译器，
// 用于完成 网易 <-> 国际版 相关数据包的翻译工作
type Translator interface {
	// 将国际版 Minecraft 数据包翻译为网易版
	ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet
	// 将网易版 Minecraft 数据包翻译为国际版
	ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet
}
