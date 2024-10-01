package marshal

import (
	"Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/minecraft/netease/protocol/packet"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"bytes"
	"fmt"
	"runtime/debug"
	"sync/atomic"

	"github.com/pterm/pterm"
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
			if r == nil {
				return
			}
			if packetFunc == nil {
				pterm.Warning.Printf(
					"DecodeNetEasePacket: Failed to unmarshal packet which numbered %d, and the error log is %v\n\n[Stack Info]\n%s\n",
					packetHeader.PacketID, r, string(debug.Stack()),
				)
				fmt.Println()
			} else {
				pterm.Warning.Printf(
					"DecodeNetEasePacket: Failed to unmarshal packet %T, and the error log is %v\n\n[Stack Info]\n%s\n",
					packetFunc(), r, string(debug.Stack()),
				)
				fmt.Println()
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
			r := recover()
			if r != nil {
				pterm.Warning.Printf(
					"EncodeNetEasePacket: Failed to marshal packet %T, and the error log is %v\n\n[Stack Info]\n%s\n",
					pk.Packet, r, string(debug.Stack()),
				)
				fmt.Println()
			}
		}()
		pk.Packet.Marshal(protocol.NewWriter(buffer, shieldID.Load()))
	}()
	return buffer.Bytes()
}
