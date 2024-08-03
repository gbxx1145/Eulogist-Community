package RaknetConnection

import (
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/packet"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
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
func (r *Raknet) SetConnection(connection net.Conn, key *ecdsa.PrivateKey) error {
	r.connection = connection
	r.encoder = packet.NewEncoder(connection)
	r.decoder = packet.NewDecoder(connection)
	r.packets = make(chan MinecraftPacket, 1024)
	r.key = key
	_, _ = rand.Read(r.salt)
	// set value
	return nil
	// return
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

// ...
func (r *Raknet) GetContext() context.Context {
	return r.context
}

// ...
func (r *Raknet) GetShieldID() int32 {
	return r.shieldID.Load()
}

// ...
func (r *Raknet) SetShieldID(shieldID int32) {
	r.shieldID.Store(shieldID)
}

// 从底层 Raknet 不断地读取多个数据包，
// 直到底层 Raknet 连接被关闭。
//
// 此函数应当只被调用一次
func (r *Raknet) ProcessIncomingPackets() {
	for {
		packets, err := r.decoder.Decode()
		if err != nil {
			r.CloseConnection()
			return
			// connection was closed
		}
		// prepare
		for _, data := range packets {
			var pk packet.Packet
			buffer := bytes.NewBuffer(data)
			reader := protocol.NewReader(buffer, r.shieldID.Load(), false)
			// prepare
			packetHeader := packet.Header{}
			packetHeader.Read(buffer)
			packetFunc := packet.ListAllPackets()[packetHeader.PacketID]
			// get header and packet func
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
			// marshal
			select {
			case <-r.context.Done():
				return
			default:
				r.packets <- MinecraftPacket{Packet: pk, Bytes: data}
			}
			// set value
		}
		// process each packet
	}
}

// 从已读取且已解码的数据包池中读取一个数据包。
// 当数据包池没有数据包时，将会阻塞，
// 直到新的已处理数据包抵达
func (r *Raknet) ReadPacket() MinecraftPacket {
	return <-r.packets
}

// 向底层 Raknet 连接写入 Minecraft 数据包 pk。
// useBytes 指代是否要直接写入 pk.Bytes 上的二进制负载
func (r *Raknet) WritePacket(pk MinecraftPacket, useBytes bool) error {
	if useBytes {
		err := r.encoder.Encode([][]byte{pk.Bytes})
		if err != nil {
			return fmt.Errorf("WritePacket: %v", err)
		}
		return nil
	}
	// use bytes to write
	buffer := bytes.NewBuffer([]byte{})
	packetHeader := packet.Header{PacketID: pk.Packet.ID()}
	packetHeader.Write(buffer)
	// get buffer and write packet header
	func() {
		defer func() {
			recover()
		}()
		pk.Packet.Marshal(protocol.NewWriter(buffer, r.shieldID.Load()))
	}()
	// marshal
	err := r.encoder.Encode([][]byte{buffer.Bytes()})
	if err != nil {
		r.CloseConnection()
	}
	// write packet
	return nil
	// return
}
