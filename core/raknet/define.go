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

// ...
type SkinManifest struct {
	Header struct {
		UUID string `json:"uuid"`
	} `json:"header"`
}

// 描述皮肤信息
type Skin struct {
	// 储存皮肤数据的二进制负载。
	// 对于普通皮肤，这是一个二进制形式的 PNG；
	// 对于高级皮肤(如 4D 皮肤)，
	// 这是一个压缩包形式的二进制表示
	FullSkinData []byte
	// 皮肤的 UUID
	SkinUUID string
	// 皮肤项目的 UUID
	SkinItemID string
	// 皮肤的一维密集像素矩阵
	SkinPixels []byte
	// 皮肤的骨架信息
	SkinGeometry []byte
	// SkinResourcePatch 是一个 JSON 编码对象，
	// 其中包含一些指向皮肤所具有的几何形状的字段。
	// 它包含的 JSON 对象指定动画的几何形状，
	// 以及播放器的默认皮肤的组合方式
	SkinResourcePatch []byte
	// 皮肤的宽度
	SkinWidth int
	// 皮肤的高度
	SkinHight int
}
