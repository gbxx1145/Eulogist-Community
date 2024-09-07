package mc_server

import (
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	"Eulogist/core/raknet/handshake"
	"Eulogist/core/raknet/marshal"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/packet_translator"
	packet_translate_pool "Eulogist/core/tools/packet_translator/pool"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"
	"Eulogist/core/tools/py_rpc"
	"bytes"
	"fmt"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

// DefaultTranslate 按默认方式翻译数据包为国际版的版本。
// 它仅仅重定向了数据包前端的 ID，
// 并用再次解析的方式保证当前翻译完全正确
func (m *MinecraftServer) DefaultTranslate(
	pk raknet_wrapper.MinecraftPacket[neteasePacket.Packet],
	standardPacketID uint32,
) raknet_wrapper.MinecraftPacket[standardPacket.Packet] {
	// 数据包可能已被修改，
	// 因此此处需要重新编码它的二进制形式
	pk.Bytes = marshal.EncodeNetEasePacket(pk, &m.ShieldID)

	// 从数据包的二进制负载前端读取其在网易协议下的数据包 ID。
	// 这一部分将会被替换为国际版协议下的数据包 ID
	packetBuffer := bytes.NewBuffer(pk.Bytes)
	_ = new(standardPacket.Header).Read(packetBuffer)

	// 取得国际版协议下数据包 ID 的二进制形式
	packetHeader := standardPacket.Header{PacketID: standardPacketID}
	headerBuffer := bytes.NewBuffer([]byte{})
	packetHeader.Write(headerBuffer)

	// 获得该数据包在国际版协议下的二进制负载，
	// 然后将其按国际版协议再次解析
	packetBytes := append(headerBuffer.Bytes(), packetBuffer.Bytes()...)
	result := marshal.DecodeStandardPacket(packetBytes, &m.ShieldID)

	// 返回值
	if result.Packet != nil {
		result.Bytes = nil
	}
	return result
}

/*
数据包过滤器过滤来自租赁服的多个数据包，
然后并将过滤后的多个数据包抄送至客户端。

如果需要，
将根据实际情况由本处的桥接直接发送回应。

writeSinglePacketToClient 指代
用于向客户端抄送数据包的函数。

syncFunc 用于将数据同步到 Minecraft，
它会在 packets 全部被处理完毕后执行，
随后，相应的数据包会被抄送至网易租赁服。

返回的 []error 是一个列表，
分别对应 packets 中每一个数据包的处理成功情况
*/
func (m *MinecraftServer) FiltePacketsAndSendCopy(
	packets []raknet_wrapper.MinecraftPacket[neteasePacket.Packet],
	writePacketsToClient func(packets []raknet_wrapper.MinecraftPacket[standardPacket.Packet]),
	syncFunc func() error,
) (errResults []error, syncError error) {
	// 初始化
	errResults = make([]error, 0)
	sendCopy := make([]raknet_wrapper.MinecraftPacket[standardPacket.Packet], 0)
	// 处理每个数据包
	for _, minecraftPacket := range packets {
		// 初始化
		var shouldSendCopy bool = true
		var err error
		// 根据数据包的类型进行不同的处理
		switch pk := minecraftPacket.Packet.(type) {
		case *neteasePacket.PyRpc:
			err = m.OnPyRpc(pk)
			if err != nil {
				err = fmt.Errorf("FiltePacketsAndSendCopy: %v", err)
			}
			shouldSendCopy = false
		case *neteasePacket.StartGame:
			// 预处理
			entityUniqueID, entityRuntimeID := handshake.HandleStartGame(m.Raknet, pk)
			m.SetEntityUniqueID(entityUniqueID)
			m.SetEntityRuntimeID(entityRuntimeID)
			playerSkin := m.GetPlayerSkin()
			// 发送简要身份证明
			m.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
				Packet: &neteasePacket.NeteaseJson{
					Data: []byte(
						fmt.Sprintf(
							`{"eventName":"LOGIN_UID","resid":"","uid":"%s"}`,
							m.GetNeteaseUID(),
						),
					),
				},
			})
			// 其他组件处理
			if playerSkin == nil {
				m.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
					Packet: &neteasePacket.PyRpc{
						Value:         py_rpc.Marshal(&py_rpc.SyncUsingMod{}),
						OperationType: neteasePacket.PyRpcOperationTypeSend,
					},
				})
			} else {
				// 初始化
				modUUIDs := make([]any, 0)
				outfitInfo := make(map[string]int64, 0)
				// 设置数据
				for modUUID, outfitType := range m.GetOutfitInfo() {
					modUUIDs = append(modUUIDs, modUUID)
					if outfitType != nil {
						outfitInfo[modUUID] = int64(*outfitType)
					}
				}
				// 组件处理
				m.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
					Packet: &neteasePacket.PyRpc{
						Value: py_rpc.Marshal(&py_rpc.SyncUsingMod{
							modUUIDs,
							playerSkin.SkinUUID,
							playerSkin.SkinItemID,
							true,
							outfitInfo,
						}),
						OperationType: neteasePacket.PyRpcOperationTypeSend,
					},
				})
			}
			// 上报自身已完成组件加载，
			// 尽管我们实际上并没有加载任何组件
			m.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
				Packet: &neteasePacket.PyRpc{
					Value:         py_rpc.Marshal(&py_rpc.ClientLoadAddonsFinishedFromGac{}),
					OperationType: neteasePacket.PyRpcOperationTypeSend,
				},
			})
		case *neteasePacket.UpdatePlayerGameType:
			if pk.PlayerUniqueID == m.entityUniqueID {
				// 如果玩家的唯一 ID 与数据包中记录的值匹配，
				// 则向客户端发送 SetPlayerGameType 数据包，
				// 并放弃当前数据包的发送，
				// 以确保 Minecraft 客户端可以正常同步游戏模式更改。
				// 否则，按原样抄送当前数据包
				sendCopy = append(sendCopy, raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					Packet: &standardPacket.SetPlayerGameType{GameType: pk.GameType},
				})
				shouldSendCopy = false
			}
		case *neteasePacket.CreativeContent:
			for index, value := range pk.Items {
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(value.Item.BlockRuntimeID))
				if found {
					pk.Items[index].Item.BlockRuntimeID = int32(standardRuntimeID)
				}
			}
		case *neteasePacket.PlayerList:
			for _, value := range pk.Entries {
				if value.EntityUniqueID == m.GetEntityUniqueID() {
					m.SetServerSkin(&value.Skin)
				}
			}
		case *neteasePacket.PlayStatus:
			if pk.Status == neteasePacket.PlayStatusPlayerSpawn {
				skinData := m.GetServerSkin()
				if skinData == nil || !skinData.Trusted {
					break
				}
				writePacketsToClient([]raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					{
						Packet: &standardPacket.PlayerSkin{
							UUID:        m.GetStandardBedrockIdentity(),
							Skin:        packet_translate_struct.ConvertToStandardSkin(*skinData),
							OldSkinName: "",
							NewSkinName: "",
						},
					},
				})
			}
		case *neteasePacket.AddItemActor:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.Item.Stack.BlockRuntimeID))
			if found {
				pk.Item.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
		case *neteasePacket.UpdateBlock:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(pk.NewBlockRuntimeID)
			if found {
				pk.NewBlockRuntimeID = standardRuntimeID
			}
		case *neteasePacket.UpdateSubChunkBlocks:
			// 初始化
			pks := []raknet_wrapper.MinecraftPacket[standardPacket.Packet]{}
			// 处理前景层的方块
			for _, value := range pk.Blocks {
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(value.BlockRuntimeID)
				if found {
					pks = append(pks, raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
						Packet: &standardPacket.UpdateBlockSynced{
							Position:          standardProtocol.BlockPos(value.BlockPos),
							NewBlockRuntimeID: standardRuntimeID,
							Flags:             value.Flags,
							Layer:             0,
							EntityUniqueID:    int64(value.SyncedUpdateEntityUniqueID),
							TransitionType:    uint64(value.SyncedUpdateType),
						},
					})
				}
			}
			// 处理背景层的方块
			for _, value := range pk.Extra {
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(value.BlockRuntimeID)
				if found {
					pks = append(pks, raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
						Packet: &standardPacket.UpdateBlockSynced{
							Position:          standardProtocol.BlockPos(value.BlockPos),
							NewBlockRuntimeID: standardRuntimeID,
							Flags:             value.Flags,
							Layer:             1,
							EntityUniqueID:    int64(value.SyncedUpdateEntityUniqueID),
							TransitionType:    uint64(value.SyncedUpdateType),
						},
					})
				}
			}
			// 发送多方块更改至客户端，
			// 并指定当前数据包不抄送
			writePacketsToClient(pks)
			shouldSendCopy = false
		case *neteasePacket.UpdateBlockSynced:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(pk.NewBlockRuntimeID)
			if found {
				pk.NewBlockRuntimeID = standardRuntimeID
			}
		case *neteasePacket.LevelEvent:
			switch pk.EventType {
			case neteasePacket.LevelEventParticlesDestroyBlock:
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.EventData))
				if found {
					pk.EventData = int32(standardRuntimeID)
				}
			case neteasePacket.LevelEventParticlesCrackBlock:
				blockFace := pk.EventData >> 24
				blockRuntimeID := pk.EventData & 0xffffff
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(blockRuntimeID))
				if found {
					pk.EventData = int32(standardRuntimeID) | blockFace<<24
				}
			}
		case *neteasePacket.InventorySlot:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.NewItem.Stack.BlockRuntimeID))
			if found {
				pk.NewItem.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
		case *neteasePacket.MobArmourEquipment:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.Helmet.Stack.BlockRuntimeID))
			if found {
				pk.Helmet.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
			standardRuntimeID, found = packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.Chestplate.Stack.BlockRuntimeID))
			if found {
				pk.Chestplate.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
			standardRuntimeID, found = packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.Leggings.Stack.BlockRuntimeID))
			if found {
				pk.Leggings.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
			standardRuntimeID, found = packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.Boots.Stack.BlockRuntimeID))
			if found {
				pk.Boots.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
		case *neteasePacket.MobEquipment:
			standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(pk.NewItem.Stack.BlockRuntimeID))
			if found {
				pk.NewItem.Stack.BlockRuntimeID = int32(standardRuntimeID)
			}
		case *neteasePacket.ItemStackResponse:
			for index, value := range pk.Responses {
				for i, v := range value.ContainerInfo {
					pk.Responses[index].ContainerInfo[i].ContainerID = packet_translate_pool.NetEaseContainerIDStandardContainerID[v.ContainerID]
				}
			}
		case *neteasePacket.ClientBoundMapItemData:
			if pk.Pixels.IsEmpty {
				pk.Height = 0
				pk.Width = 0
			}
		case *neteasePacket.InventoryContent:
			// 初始化
			allExist := true
			pks := []raknet_wrapper.MinecraftPacket[standardPacket.Packet]{}
			// 遍历物品变动表中的每个物品堆栈实例
			for slot, item := range pk.Content {
				// 如果物品的固定网络堆栈 ID 为 0，说明这是个空气，
				// 如果为 -1，说明当前物品未被更改。
				// allExist 指示是否有某个物品未被更改
				if item.Stack.NetworkID == 0 || item.Stack.NetworkID == -1 {
					if item.Stack.NetworkID == -1 {
						allExist = false
					}
					continue
				}
				// 更新每个物品堆栈实例对应的 block runtime id
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(item.Stack.BlockRuntimeID))
				if found {
					pk.Content[slot].Stack.BlockRuntimeID = int32(standardRuntimeID)
				}
				// 当 allExist 为假时，
				// 说明存在某个物品未被更改。
				// 对于国际版，国际版不能识别这样的物品，
				// 因此，我们采用该方式来逐个更新产生变化的槽位
				pks = append(pks, raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					Packet: &standardPacket.InventorySlot{
						WindowID: pk.WindowID,
						Slot:     uint32(slot),
						NewItem:  packet_translate_struct.ConvertToStandardItemInstance(item),
					},
				})
			}
			// 当存在某个物品未被更改时，
			// 我们采用该方式来逐个更新产生变化的槽位
			if !allExist {
				writePacketsToClient(pks)
				shouldSendCopy = false
			}
		default:
			// 默认情况下，
			// 我们需要将数据包同步到客户端
		}
		// 提交子结果
		errResults = append(errResults, err)
		if shouldSendCopy {
			// 确定当前数据包是否处理成功
			if minecraftPacket.Packet == nil {
				continue
			}
			// 查找当前数据包对应的国际版 ID
			standardPacketID, found := packet_translate_pool.NetEasePacketIDToStandardPacketID[minecraftPacket.Packet.ID()]
			if !found {
				continue
			}
			// 确认当前数据包是否需要翻译
			if translator := packet_translator.TranslatorPool[standardPacketID]; translator != nil {
				sendCopy = append(sendCopy, raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					Packet: translator.ToStandardPacket(minecraftPacket.Packet),
				})
				continue
			}
			// 当前数据包无需翻译，
			// 可以直接传输其二进制负载
			sendCopy = append(sendCopy, m.DefaultTranslate(minecraftPacket, standardPacketID))
		}
	}
	// 同步数据并抄送数据包
	err := syncFunc()
	writePacketsToClient(sendCopy)
	// 返回值
	if err != nil {
		return errResults, fmt.Errorf("FiltePacketsAndSendCopy: %v", err)
	} else {
		return errResults, nil
	}
}
