package raknet_wrapper

import (
	"context"
	"crypto/ecdsa"
	"net"
	"sync"
	"sync/atomic"

	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// 描述一个简单的，但可以支持不同 Minecraft
// 通信协议的基本 Raknet 连接实例
type Raknet[T any] struct {
	Connection net.Conn
	ShieldID   atomic.Int32

	Context context.Context
	Cancel  context.CancelFunc

	Closed     bool
	ClosedLock sync.Mutex

	Encoder *packet.Encoder
	Decoder *packet.Decoder

	DecodePacket func(buf []byte, shieldID *atomic.Int32) (pk MinecraftPacket[T])
	EncodePacket func(pk MinecraftPacket[T], shieldID *atomic.Int32) (buf []byte)

	Key  *ecdsa.PrivateKey
	Salt []byte

	Packets chan ([]MinecraftPacket[T])
}

// 描述 Minecraft 数据包
type MinecraftPacket[T any] struct {
	// 保存已解码的 Minecraft 数据包
	Packet T
	// 保存数据包的二进制形式。
	//
	// 在发送 MinecraftPacket 时，
	// 将优先采用该字段，除非该字段为空，
	// 则将会尝试编码 Packet 字段，
	// 然后将编码结果发送到 Raknet 上
	Bytes []byte
}
