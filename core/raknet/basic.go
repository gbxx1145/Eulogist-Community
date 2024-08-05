package RaknetConnection

import (
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/packet"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"net"

	"github.com/pterm/pterm"
)

// 初始化一个空的 Raknet
func NewRaknet() *Raknet {
	ctx, cancel := context.WithCancel(context.Background())
	return &Raknet{
		context: ctx,
		cancel:  cancel,
	}
}

// 将底层 Raknet 连接设置为 connection，
// 并指定 服务器/客户端 私钥为 key
func (r *Raknet) SetConnection(connection net.Conn, key *ecdsa.PrivateKey) {
	r.connection = connection
	r.encoder = packet.NewEncoder(connection)
	r.decoder = packet.NewDecoder(connection)
	r.packets = make(chan []MinecraftPacket, 255)
	r.key = key
	_, _ = rand.Read(r.salt)
}

// 关闭已建立的 Raknet 底层连接
func (r *Raknet) CloseConnection() {
	r.closedLock.Lock()
	defer r.closedLock.Unlock()

	r.cancel()
	r.connection.Close()

	if !r.closed {
		close(r.packets)
		r.closed = true
	}
}

// 获取当前的上下文
func (r *Raknet) GetContext() context.Context {
	return r.context
}

// 获取当前的 Shield ID
func (r *Raknet) GetShieldID() int32 {
	return r.shieldID.Load()
}

// 设置 Shield ID
func (r *Raknet) SetShieldID(shieldID int32) {
	r.shieldID.Store(shieldID)
}

// 从底层 Raknet 不断地读取多个数据包，
// 直到底层 Raknet 连接被关闭。
//
// 此函数应当只被调用一次
func (r *Raknet) ProcessIncomingPackets() {
	for {
		// 从底层 Raknet 连接读取数据包
		packets, err := r.decoder.Decode()
		if err != nil {
			// 此时从底层 Raknet 连接读取数据包遭遇了错误，
			// 因此我们认为连接已被关闭
			r.CloseConnection()
			return
		}
		// 处理每个数据包
		packetSlice := make([]MinecraftPacket, len(packets))
		for index, data := range packets {
			// 准备读取数据包
			var pk packet.Packet
			buffer := bytes.NewBuffer(data)
			reader := protocol.NewReader(buffer, r.shieldID.Load(), false)
			// 获取数据包头和数据包处理函数
			packetHeader := packet.Header{}
			packetHeader.Read(buffer)
			packetFunc := packet.ListAllPackets()[packetHeader.PacketID]
			// 序列化数据包
			func() {
				defer func() {
					r := recover()
					if r != nil {
						pterm.Warning.Printf("ProcessIncomingPackets: %v\n", err)
					}
				}()
				switch packetHeader.PacketID {
				case packet.IDRequestNetworkSettings, packet.IDNetworkSettings:
				case packet.IDLogin:
				case packet.IDServerToClientHandshake, packet.IDClientToServerHandshake:
				case packet.IDStartGame, packet.IDPyRpc:
				case packet.IDUpdatePlayerGameType:
				default:
					return
				}
				pk = packetFunc()
				pk.Marshal(reader)
			}()
			// 同步数据包到待存区
			select {
			case <-r.context.Done():
				r.CloseConnection()
				return
			default:
				packetSlice[index] = MinecraftPacket{Packet: pk, Bytes: data}
			}
		}
		// 提交
		r.packets <- packetSlice
	}
}

// 从已读取且已解码的数据包池中读取多个数据包。
// 当数据包池没有数据包时，将会阻塞，
// 直到新的已处理数据包抵达
func (r *Raknet) ReadPackets() []MinecraftPacket {
	return <-r.packets
}

// 向底层 Raknet 连接写多个 Minecraft 数据包 pk。
// useBytes 指代是否要直接采用这些数据包的二进制负载，
// 然后写入到底层 Raknet 连接
func (r *Raknet) WritePackets(pk []MinecraftPacket, useBytes bool) {
	// 准备
	packetBytes := make([][]byte, len(pk))
	// 如果考虑使用字节直接写入
	if useBytes {
		for index, singlePacket := range pk {
			packetBytes[index] = singlePacket.Bytes
		}
		err := r.encoder.Encode(packetBytes)
		if err != nil {
			// 此时向底层 Raknet 连接写入数据包遭遇了错误，
			// 因此我们认为连接已被关闭
			r.CloseConnection()
		}
		return
	}
	// 处理多个数据包
	for index, singlePacket := range pk {
		// 获取缓冲区并写入数据包头
		buffer := bytes.NewBuffer([]byte{})
		packetHeader := packet.Header{PacketID: singlePacket.Packet.ID()}
		packetHeader.Write(buffer)
		// 序列化数据包
		func() {
			defer func() {
				recover()
			}()
			singlePacket.Packet.Marshal(protocol.NewWriter(buffer, r.shieldID.Load()))
		}()
		packetBytes[index] = buffer.Bytes()
	}
	// 写入数据包
	err := r.encoder.Encode(packetBytes)
	if err != nil {
		// 此时向底层 Raknet 连接写入数据包遭遇了错误，
		// 因此我们认为连接已被关闭
		r.CloseConnection()
	}
}

// 向底层 Raknet 连接写单个 Minecraft 数据包 pk。
// useBytes 指代是否要直接写入 pk.Bytes 上的二进制负载
func (r *Raknet) WriteSinglePacket(pk MinecraftPacket, useBytes bool) {
	r.WritePackets([]MinecraftPacket{pk}, useBytes)
}
