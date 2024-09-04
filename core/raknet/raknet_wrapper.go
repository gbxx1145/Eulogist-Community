package raknet_connection

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"
	"Eulogist/core/raknet/marshal"
	raknet_wrapper "Eulogist/core/raknet/wrapper"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

// 获取一个基于网易 Minecraft 协议的 Raknet 包装器
func NewNetEaseRaknetWrapper() *raknet_wrapper.Raknet[neteasePacket.Packet] {
	return raknet_wrapper.NewRaknet(marshal.DecodeNetEasePacket, marshal.EncodeNetEasePacket)
}

// 获取一个基于国际版 Minecraft 协议的 Raknet 包装器
func NewStandardRaknetWrapper() *raknet_wrapper.Raknet[standardPacket.Packet] {
	return raknet_wrapper.NewRaknet(marshal.DecodeStandardPacket, marshal.EncodeStandardPacket)
}
