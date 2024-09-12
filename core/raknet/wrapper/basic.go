package raknet_wrapper

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"net"
	"sync/atomic"

	"Eulogist/core/minecraft/standard/protocol/packet"
)

// 初始化一个空的 Raknet。
//
// decodePacket 用于解码数据包，
// 而 encodePacket 则用于编码数据包
func NewRaknet[T any](
	decodePacket func(buf []byte, shieldID *atomic.Int32) (pk MinecraftPacket[T]),
	encodePacket func(pk MinecraftPacket[T], shieldID *atomic.Int32) (buf []byte),
) *Raknet[T] {
	ctx, cancel := context.WithCancel(context.Background())
	return &Raknet[T]{
		Context:      ctx,
		Cancel:       cancel,
		DecodePacket: decodePacket,
		EncodePacket: encodePacket,
	}
}

// 将底层 Raknet 连接设置为 connection，
// 并指定 服务器/客户端 私钥为 key
func (r *Raknet[T]) SetConnection(connection net.Conn, key *ecdsa.PrivateKey) {
	r.Connection = connection
	r.Encoder = packet.NewEncoder(connection)
	r.Decoder = packet.NewDecoder(connection)
	r.Decoder.DisableBatchPacketLimit()
	r.Packets = make(chan []MinecraftPacket[T], 255)
	r.Key = key
	_, _ = rand.Read(r.Salt)
}

// 关闭已建立的 Raknet 底层连接
func (r *Raknet[T]) CloseConnection() {
	r.ClosedLock.Lock()
	defer r.ClosedLock.Unlock()

	r.Cancel()

	if r.Connection != nil {
		r.Connection.Close()
	}

	if !r.Closed && r.Packets != nil {
		close(r.Packets)
		r.Closed = true
	}
}

/*
从底层 Raknet 不断地读取多个数据包，
直到底层 Raknet 连接被关闭。

在大多数情况下，由于我们只需按原样传递数据包，
因此，我们只解码了一部分必须的数据包。

而对于其他的数据包，我们不作额外处理，
而是仅仅地保留它们的二进制负载

另，此函数应当只被调用一次
*/
func (r *Raknet[T]) ProcessIncomingPackets() {
	// 确保该函数不会返回恐慌
	defer func() {
		recover()
	}()
	// 不断处理到来的一个或多个数据包
	for {
		// 从底层 Raknet 连接读取数据包
		packets, err := r.Decoder.Decode()
		if err != nil {
			// 此时从底层 Raknet 连接读取数据包遭遇了错误，
			// 因此我们认为连接已被关闭
			r.CloseConnection()
			return
		}
		// 处理每个数据包
		packetSlice := make([]MinecraftPacket[T], len(packets))
		for index, data := range packets {
			pk := r.DecodePacket(data, &r.ShieldID)
			packetSlice[index] = pk
		}
		// 提交
		select {
		case <-r.Context.Done():
			r.CloseConnection()
			return
		default:
			r.Packets <- packetSlice
		}
	}
}

/*
从已读取且已解码的数据包池中读取多个数据包。

当数据包池没有数据包时，将会阻塞，
直到新的已处理数据包抵达。

在大多数情况下，由于我们只需按原样传递数据包，
因此，在读取时，我们只解码了一部分必须的数据包，
而对于其他的数据包，我们将仅仅地保留它们的二进制负载
*/
func (r *Raknet[T]) ReadPackets() []MinecraftPacket[T] {
	return <-r.Packets
}

// 向底层 Raknet 连接写多个 Minecraft 数据包 pk。
// WritePackets 会优先采用每个数据包的二进制负载，
// 除非负载为空，则此时再转而编码对应的数据包，
// 然后写入到 Raknet 底层连接
func (r *Raknet[T]) WritePackets(pk []MinecraftPacket[T]) {
	// 准备
	packetBytes := make([][]byte, len(pk))
	for index, singlePacket := range pk {
		// 先采用当前数据包的二进制负载
		if len(singlePacket.Bytes) > 0 {
			packetBytes[index] = singlePacket.Bytes
			continue
		}
		// 此时当前数据包不存在已编码的二进制负载，
		// 因此我们主动编码它
		packetBytes[index] = r.EncodePacket(singlePacket, &r.ShieldID)
	}
	// 将数据包写入底层 Raknet 连接
	encodeError := r.Encoder.Encode(packetBytes)
	if encodeError != nil {
		// 此时向底层 Raknet 连接写入数据包遭遇了错误，
		// 因此我们认为连接已被关闭
		r.CloseConnection()
	}
}

// 向底层 Raknet 连接写单个 Minecraft 数据包 pk。
// WriteSinglePacket 会优先采用它的二进制负载，
// 除非负载为空，则此时再转而编码该数据包为二进制形式，
// 然后再写入到 Raknet 底层连接
func (r *Raknet[T]) WriteSinglePacket(pk MinecraftPacket[T]) {
	r.WritePackets([]MinecraftPacket[T]{pk})
}
