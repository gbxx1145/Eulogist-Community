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

type Raknet struct {
	connection net.Conn
	shieldID   atomic.Int32

	context context.Context
	cancel  context.CancelFunc

	closed     bool
	closedLock sync.Mutex

	encoder      *packet.Encoder
	decoder      *packet.Decoder
	shouldDecode bool

	key  *ecdsa.PrivateKey
	salt []byte

	packets chan (MinecraftPacket)
}

type MinecraftPacket struct {
	Packet packet.Packet
	Bytes  []byte
}
