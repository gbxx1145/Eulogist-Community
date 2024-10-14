package raknet_wrapper

import (
	"context"
	"crypto/ecdsa"
	"net"
	"sync"
	"sync/atomic"

	"Eulogist/core/minecraft/protocol/packet"
)

// 描述一个简单的，但可以支持不同 Minecraft
// 通信协议的基本 Raknet 连接实例
type Raknet struct {
	connection net.Conn
	shieldID   atomic.Int32

	context context.Context
	cancel  context.CancelFunc

	closed     bool
	closedLock sync.Mutex

	encoder *packet.Encoder
	decoder *packet.Decoder

	key  *ecdsa.PrivateKey
	salt []byte

	packets chan ([]MinecraftPacket)
}

// 描述 Minecraft 数据包
type MinecraftPacket struct {
	// 保存已解码的 Minecraft 数据包
	Packet packet.Packet
	// 保存数据包的二进制形式
	Bytes []byte
}
