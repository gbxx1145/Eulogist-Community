package ModPC

import (
	"Eulogist/minecraft/protocol"
	"Eulogist/minecraft/protocol/packet"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/sandertv/go-raknet"
)

// 在指定的 IP 地址及端口号上创建 Raknet 侦听连接，
// 并初始化管道 s.Connected 以供后续操作。
//
// 任何 Minecraft 客户端可以连接到该地址
func (s *Server) CreateListener(address string) error {
	listener, err := raknet.Listen(address)
	if err != nil {
		return fmt.Errorf("CreateListener: %v", err)
	}
	s.listener = listener
	s.connected = make(chan struct{}, 1)
	s.closed = make(chan struct{}, 1)
	return nil
}

// 等待 Minecraft 客户端连接到服务器
func (s *Server) WaitConnect() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return fmt.Errorf("WaitConnect: %v", err)
	}
	// accept connection
	s.connection = conn
	s.encoder = packet.NewEncoder(conn)
	s.decoder = packet.NewDecoder(conn)
	s.packets = make(chan MinecraftPacket, 1024)
	s.serverKey, _ = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	_, _ = rand.Read(s.salt)
	s.connected <- struct{}{}
	// set value
	return nil
	// return
}

// 从底层 Raknet 不断地读取多个数据包，
// 直到底层 Raknet 连接被关闭。
//
// 此函数应当只被调用一次
func (s *Server) ProcessIncomingPackets() {
	for {
		packets, err := s.decoder.Decode()
		if err != nil {
			close(s.packets)
			s.closed <- struct{}{}
			return
			// connection was closed
		}
		// prepare
		for _, data := range packets {
			buffer := bytes.NewBuffer(data)
			reader := protocol.NewReader(buffer, 0, false)
			// prepare
			packetHeader := packet.Header{}
			packetHeader.Read(buffer)
			packetFunc := packet.ListAllPackets()[packetHeader.PacketID]
			// get header and packet func
			pk := packetFunc()
			pk.Marshal(reader)
			// marshal
			s.packets <- MinecraftPacket{Packet: pk, Bytes: data}
			// set value
		}
		// process each packet
	}
}

// 从已读取且已解码的数据包池中读取一个数据包。
// 当数据包池没有数据包时，将会阻塞，
// 直到新的已处理数据包抵达
func (s *Server) ReadPacket() MinecraftPacket {
	return <-s.packets
}

// 向底层 Raknet 连接写入 Minecraft 数据包 pk。
// useBytes 指代是否直接写入 pk.Bytes 上二进制负载
func (s *Server) WritePacket(pk MinecraftPacket, useBytes bool) error {
	if useBytes {
		err := s.encoder.Encode([][]byte{pk.Bytes})
		if err != nil {
			return fmt.Errorf("WritePacket: %v", err)
		}
		return nil
	}
	// use bytes to write
	buffer := bytes.NewBuffer([]byte{})
	packetHeader := packet.Header{PacketID: pk.Packet.ID()}
	packetHeader.Write(buffer)
	pk.Packet.Marshal(protocol.NewWriter(buffer, 0))
	err := s.encoder.Encode([][]byte{buffer.Bytes()})
	if err != nil {
		s.closed <- struct{}{}
	}
	// marshal and write
	return nil
	// return
}
