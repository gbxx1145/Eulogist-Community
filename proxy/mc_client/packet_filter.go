package mc_client

import (
	"Eulogist/core/raknet/marshal"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/packet_translator"
	packet_translate_pool "Eulogist/core/tools/packet_translator/pool"
	"bytes"
	"fmt"

	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"

	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

// DefaultTranslate 按默认方式翻译数据包为网易的版本。
// 它仅仅重定向了数据包前端的 ID，
// 并用再次解析的方式保证当前翻译完全正确
func (m *MinecraftClient) DefaultTranslate(
	pk raknet_wrapper.MinecraftPacket[standardPacket.Packet],
	neteasePacketID uint32,
) raknet_wrapper.MinecraftPacket[neteasePacket.Packet] {
	// 数据包可能已被修改，
	// 因此此处需要重新编码它的二进制形式
	pk.Bytes = marshal.EncodeStandardPacket(pk, &m.ShieldID)

	// 从数据包的二进制负载前端读取其在国际版协议下的数据包 ID。
	// 这一部分将会被替换为网易版协议下的数据包 ID
	packetBuffer := bytes.NewBuffer(pk.Bytes)
	_ = new(neteasePacket.Header).Read(packetBuffer)

	// 取得网易版协议下数据包 ID 的二进制形式
	packetHeader := neteasePacket.Header{PacketID: neteasePacketID}
	headerBuffer := bytes.NewBuffer([]byte{})
	packetHeader.Write(headerBuffer)

	// 获得该数据包在国际版协议下的二进制负载，
	// 然后将其按国际版协议再次解析，然后返回值
	packetBytes := append(headerBuffer.Bytes(), packetBuffer.Bytes()...)
	return marshal.DecodeNetEasePacket(packetBytes, &m.ShieldID)
}

/*
数据包过滤器过滤来自 Minecraft 客户端的多个数据包，
然后并将过滤后的多个数据包抄送至网易租赁服。

如果需要，
将根据实际情况由本处的桥接直接发送回应。

writePacketsToServer 指代
用于向客户端抄送数据包的函数。

syncFunc 用于将数据同步到网易租赁服，
它会在 packets 全部被处理完毕后执行，
随后，相应的数据包会被抄送至客户端。

返回的 []error 是一个列表，
分别对应 packets 中每一个数据包的处理成功情况
*/
func (m *MinecraftClient) FiltePacketsAndSendCopy(
	packets []raknet_wrapper.MinecraftPacket[standardPacket.Packet],
	writePacketsToServer func(packets []raknet_wrapper.MinecraftPacket[neteasePacket.Packet]),
	syncFunc func() error,
) (errResults []error, syncError error) {
	// 初始化
	errResults = make([]error, 0)
	sendCopy := make([]raknet_wrapper.MinecraftPacket[neteasePacket.Packet], 0)
	// 处理每个数据包
	for _, minecraftPacket := range packets {
		// 初始化
		var shouldSendCopy bool = true
		var err error
		// 根据数据包的类型进行不同的处理
		switch pk := minecraftPacket.Packet.(type) {
		case *standardPacket.ClientCacheStatus:
			writePacketsToServer([]raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
				{Packet: &neteasePacket.ClientCacheStatus{Enabled: false}},
			})
			shouldSendCopy = false
		case *standardPacket.InventoryTransaction:
			data, ok := pk.TransactionData.(*standardProtocol.UseItemTransactionData)
			if ok {
				standardRuntimeID, found := packet_translator.ConvertToNetEaseBlockRuntimeID(data.BlockRuntimeID)
				if found {
					data.BlockRuntimeID = standardRuntimeID
				}
			}
		case *standardPacket.CraftingEvent:
			for index, value := range pk.Input {
				standardRuntimeID, found := packet_translator.ConvertToNetEaseBlockRuntimeID(uint32(value.Stack.BlockRuntimeID))
				if found {
					pk.Input[index].Stack.BlockRuntimeID = int32(standardRuntimeID)
				}
			}
			for index, value := range pk.Output {
				standardRuntimeID, found := packet_translator.ConvertToNetEaseBlockRuntimeID(uint32(value.Stack.BlockRuntimeID))
				if found {
					pk.Output[index].Stack.BlockRuntimeID = int32(standardRuntimeID)
				}
			}
		case *standardPacket.MobEquipment:
			standardRuntimeID, found := packet_translator.ConvertToNetEaseBlockRuntimeID(uint32(pk.NewItem.Stack.BlockRuntimeID))
			if found {
				pk.NewItem.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
		case *standardPacket.RequestPermissions:
			pk.PermissionLevel = pk.PermissionLevel / 2
		default:
			// 默认情况下，我们需要将
			// 数据包同步到网易租赁服
		}
		// 提交子结果
		errResults = append(errResults, err)
		if shouldSendCopy {
			// 取得当前数据包相关联的 ID
			standardPacketID := minecraftPacket.Packet.ID()
			neteasePacketID := packet_translate_pool.StandardPacketIDToNetEasePacketID[standardPacketID]
			// 确认当前数据包是否需要翻译
			if translator := packet_translator.TranslatorPool[standardPacketID]; translator != nil {
				sendCopy = append(sendCopy, raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
					Packet: translator.ToNetEasePacket(minecraftPacket.Packet),
				})
				continue
			}
			// 当前数据包无需翻译，
			// 可以直接传输其二进制负载
			sendCopy = append(sendCopy, m.DefaultTranslate(minecraftPacket, neteasePacketID))
		}
	}
	// 同步数据并抄送数据包
	err := syncFunc()
	writePacketsToServer(sendCopy)
	// 返回值
	if err != nil {
		return errResults, fmt.Errorf("FiltePacketsAndSendCopy: %v", err)
	} else {
		return errResults, nil
	}
}
