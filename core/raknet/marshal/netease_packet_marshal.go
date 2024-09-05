package marshal

import (
	"Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/minecraft/netease/protocol/packet"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"bytes"
	"sync/atomic"
)

// 将负载为 buf 的网易 Minecraft 数据包解码，
// 并指定解码时使用的 shieldID
func DecodeNetEasePacket(buf []byte, shieldID *atomic.Int32) (
	pk raknet_wrapper.MinecraftPacket[packet.Packet],
) {
	// 初始化
	var minecraftPacket packet.Packet
	buffer := bytes.NewBuffer(buf)
	reader := protocol.NewReader(buffer, shieldID.Load(), false)
	// 获取数据包头和数据包处理函数
	packetHeader := packet.Header{}
	packetHeader.Read(buffer)
	packetFunc := packet.ListAllPackets()[packetHeader.PacketID]
	if packetFunc == nil {
		return raknet_wrapper.MinecraftPacket[packet.Packet]{Bytes: buf}
	}
	// 序列化数据包
	func() {
		defer func() {
			r := recover()
			if r != nil {
				minecraftPacket = nil
			}
		}()
		minecraftPacket = packetFunc()
		minecraftPacket.Marshal(reader)
	}()
	// 返回值
	return raknet_wrapper.MinecraftPacket[packet.Packet]{Packet: minecraftPacket, Bytes: buf}
}

// 将网易 Minecraft 数据包 pk 编码，
// 并指定编码时使用的 shieldID
func EncodeNetEasePacket(
	pk raknet_wrapper.MinecraftPacket[packet.Packet],
	shieldID *atomic.Int32,
) (buf []byte) {
	buffer := bytes.NewBuffer([]byte{})
	packetHeader := packet.Header{PacketID: pk.Packet.ID()}
	packetHeader.Write(buffer)
	func() {
		defer func() {
			recover()
		}()
		pk.Packet.Marshal(protocol.NewWriter(buffer, shieldID.Load()))
	}()
	return buffer.Bytes()
}
