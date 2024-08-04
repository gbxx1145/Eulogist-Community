package RaknetConnection

import (
	"Eulogist/core/minecraft/protocol/packet"
	"context"
	"crypto/ecdsa"
	"net"
	"sync"
	"sync/atomic"
)

// saltClaims 是保存服务器在
// ServerToClientHandshake
// 数据包中发送的 salt 的声明
type saltClaims struct {
	Salt string `json:"salt"`
}

// 描述一个基本的 Raknet 连接实例
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

	packets chan (MinecraftPacket)
}

// 描述 Minecraft 数据包
type MinecraftPacket struct {
	Packet packet.Packet
	Bytes  []byte
}

// 描述皮肤信息
type Skin struct {
	SkinImageData     []byte // 皮肤的 PNG 二进制形式
	SkinPixels        []byte // 皮肤的一维密集像素矩阵
	SkinGeometry      []byte // 皮肤的骨架信息
	SkinResourcePatch []byte // ...
	SkinWidth         int    // 皮肤的宽度
	SkinHight         int    // 皮肤的高度
}
