package mc_server

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	neteasePacket "Eulogist/core/minecraft/netease/protocol/packet"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	"Eulogist/core/raknet/handshake"
	"Eulogist/core/raknet/marshal"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/packet_translator"
	packet_translate_pool "Eulogist/core/tools/packet_translator/pool"
	packet_translate_struct "Eulogist/core/tools/packet_translator/struct"
	"Eulogist/core/tools/py_rpc"
	"Eulogist/proxy/persistence_data"
	"bytes"
	"fmt"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"

	"github.com/google/uuid"
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
	pk.Bytes = marshal.EncodeNetEasePacket(pk, m.Conn.GetShieldID())

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
	result := marshal.DecodeStandardPacket(packetBytes, m.Conn.GetShieldID())

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
			// 初始化变量
			m.PersistenceData.LoginData.PlayerUniqueID, m.PersistenceData.LoginData.PlayerRuntimeID = handshake.HandleStartGame(m.Conn, pk)
			m.PersistenceData.BotDimension = persistence_data.BotDimensionData{
				Dimension:  pk.Dimension, // 保存当前的维度信息
				ChangeDown: true,         // 玩家此时已经位于对应的维度，因此被认为已完成维度更改
			}
			playerSkin := m.PersistenceData.SkinData.NeteaseSkin
			// 处理维度信息
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
			// 发送简要身份证明
			m.Conn.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
				Packet: &neteasePacket.NeteaseJson{
					Data: []byte(
						fmt.Sprintf(
							`{"eventName":"LOGIN_UID","resid":"","uid":"%d"}`,
							m.PersistenceData.LoginData.Server.IdentityData.Uid,
						),
					),
				},
			})
			// 其他组件处理
			if playerSkin == nil {
				m.Conn.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
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
				for modUUID, outfitType := range m.PersistenceData.BotComponent {
					modUUIDs = append(modUUIDs, modUUID)
					if outfitType != nil {
						outfitInfo[modUUID] = int64(*outfitType)
					}
				}
				// 组件处理
				m.Conn.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
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
			m.Conn.WriteSinglePacket(raknet_wrapper.MinecraftPacket[neteasePacket.Packet]{
				Packet: &neteasePacket.PyRpc{
					Value:         py_rpc.Marshal(&py_rpc.ClientLoadAddonsFinishedFromGac{}),
					OperationType: neteasePacket.PyRpcOperationTypeSend,
				},
			})
		case *neteasePacket.AddPlayer:
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.AddPlayer = append(
					m.PersistenceData.BotDimension.DataCache.AddPlayer,
					*packet_translator.TranslatorPool[pk.ID()].ToStandardPacket(pk).(*standardPacket.AddPlayer),
				)
			}
		case *neteasePacket.AddActor:
			m.PersistenceData.AddWorldEntity(persistence_data.EntityData{
				EntityType:      pk.EntityType,
				EntityRuntimeID: pk.EntityRuntimeID,
				EntityUniqueID:  pk.EntityUniqueID,
			})
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.AddActor = append(
					m.PersistenceData.BotDimension.DataCache.AddActor,
					*packet_translator.TranslatorPool[pk.ID()].ToStandardPacket(pk).(*standardPacket.AddActor),
				)
			}
			if pk.EntityType == "minecraft:falling_block" {
				if entityFlags, ok := pk.EntityMetadata[neteaseProtocol.EntityDataKeyFlags].(int64); ok {
					pk.EntityMetadata[neteaseProtocol.EntityDataKeyFlags] = entityFlags ^ 0x1000000000000 // Ingore collision
				}
				if fallingBlockRuntimeID, ok := pk.EntityMetadata[neteaseProtocol.EntityDataKeyVariant].(int32); ok {
					standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(fallingBlockRuntimeID))
					if found {
						pk.EntityMetadata[neteaseProtocol.EntityDataKeyVariant] = int32(standardRuntimeID)
					}
				}
			}
		case *neteasePacket.SetActorData:
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.SetActorData = append(
					m.PersistenceData.BotDimension.DataCache.SetActorData,
					*packet_translator.TranslatorPool[pk.ID()].ToStandardPacket(pk).(*standardPacket.SetActorData),
				)
			}
			entity := m.PersistenceData.GetWorldEntityByRuntimeID(pk.EntityRuntimeID)
			if entity == nil || entity.EntityType != "minecraft:falling_block" {
				break
			}
			if entityFlags, ok := pk.EntityMetadata[neteaseProtocol.EntityDataKeyFlags].(int64); ok {
				pk.EntityMetadata[neteaseProtocol.EntityDataKeyFlags] = entityFlags ^ 0x1000000000000 // Ingore collision
			}
			if fallingBlockRuntimeID, ok := pk.EntityMetadata[neteaseProtocol.EntityDataKeyVariant].(int32); ok {
				standardRuntimeID, found := packet_translator.ConvertToStandardBlockRuntimeID(uint32(fallingBlockRuntimeID))
				if found {
					pk.EntityMetadata[neteaseProtocol.EntityDataKeyVariant] = int32(standardRuntimeID)
				}
			}
		case *neteasePacket.RemoveActor:
			m.PersistenceData.DeleteWorldEntityByUniqueID(pk.EntityUniqueID)
		case *neteasePacket.AddItemActor:
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.AddItemActor = append(
					m.PersistenceData.BotDimension.DataCache.AddItemActor,
					*packet_translator.TranslatorPool[pk.ID()].ToStandardPacket(pk).(*standardPacket.AddItemActor),
				)
			}
		case *neteasePacket.AddPainting:
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.AddPainting = append(
					m.PersistenceData.BotDimension.DataCache.AddPainting,
					standardPacket.AddPainting(*pk),
				)
			}
		case *neteasePacket.AddVolumeEntity:
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.AddVolumeEntity = append(
					m.PersistenceData.BotDimension.DataCache.AddVolumeEntity,
					*packet_translator.TranslatorPool[pk.ID()].ToStandardPacket(pk).(*standardPacket.AddVolumeEntity),
				)
			}
		case *neteasePacket.RemoveVolumeEntity:
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
		case *neteasePacket.ChangeDimension:
			// 同步维度数据
			lastDimensionID := m.PersistenceData.BotDimension.Dimension
			m.PersistenceData.BotDimension = persistence_data.BotDimensionData{
				Dimension:  pk.Dimension,
				Position:   pk.Position,
				Respawn:    pk.Respawn,
				ChangeDown: lastDimensionID <= neteasePacket.DimensionEnd && pk.Dimension <= neteasePacket.DimensionEnd,
			}
			// 如果目标维度和原本维度有一个不是原版维度，
			// 则赞颂者需要进行额外处理
			if !m.PersistenceData.BotDimension.ChangeDown {
				writePacketsToClient([]raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					{
						Packet: &standardPacket.ChangeDimension{
							Dimension: m.PersistenceData.BotDimension.GetTransferDimensionID(lastDimensionID, pk.Dimension),
							Position:  pk.Position,
							Respawn:   pk.Respawn,
						},
					},
				})
				shouldSendCopy = false
			}
		case *neteasePacket.PositionTrackingDBServerBroadcast:
			if dimension, ok := pk.Payload["dim"].(int32); ok {
				if dimension > neteasePacket.DimensionEnd {
					pk.Payload["dim"] = int32(neteasePacket.DimensionOverworld)
				}
			}
		case *neteasePacket.SpawnParticleEffect:
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
		case *neteasePacket.LevelChunk:
			// 当赞颂者需要处理非主世界维度的特殊情况时，
			// 需要暂存一部分生物群落数据。
			// 这些数据将在用户最终抵达正确维度时由赞颂者发送，
			// 随后赞颂者再清除暂存的这些数据
			if !m.PersistenceData.BotDimension.ChangeDown {
				m.PersistenceData.BotDimension.DataCache.LevelChunk = append(
					m.PersistenceData.BotDimension.DataCache.LevelChunk,
					standardPacket.LevelChunk{
						Position:        standardProtocol.ChunkPos(pk.Position),
						HighestSubChunk: pk.HighestSubChunk,
						SubChunkCount:   pk.SubChunkCount,
						CacheEnabled:    pk.CacheEnabled,
						BlobHashes:      pk.BlobHashes,
						RawPayload:      pk.RawPayload,
					},
				)
			}
		case *neteasePacket.SubChunk:
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
		case *neteasePacket.UpdatePlayerGameType:
			if pk.PlayerUniqueID == m.PersistenceData.LoginData.PlayerUniqueID {
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
				if value.EntityUniqueID == m.PersistenceData.LoginData.PlayerUniqueID {
					m.PersistenceData.SkinData.ServerSkin = &value.Skin
				}
			}
		case *neteasePacket.PlayStatus:
			if pk.Status == neteasePacket.PlayStatusPlayerSpawn {
				// 初始化变量
				var playerUUID uuid.UUID
				skinData := m.PersistenceData.SkinData.ServerSkin
				// 判断皮肤是否存在且皮肤是否可信
				if skinData == nil || !skinData.Trusted {
					break
				}
				// 解析 Minecraft 客户端处原本的玩家 UUID
				playerUUID, err = uuid.Parse(m.PersistenceData.LoginData.Client.IdentityData.Identity)
				if err != nil {
					err = fmt.Errorf("FiltePacketsAndSendCopy: %v", err)
					break
				}
				// 同步 Minecraft 客户端处的皮肤为网易账户对应的皮肤
				writePacketsToClient([]raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
					{
						Packet: &standardPacket.PlayerSkin{
							UUID:        playerUUID,
							Skin:        packet_translate_struct.ConvertToStandardSkin(*skinData),
							OldSkinName: "",
							NewSkinName: "",
						},
					},
				})
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
			if pk.Dimension > neteasePacket.DimensionEnd {
				pk.Dimension = neteasePacket.DimensionOverworld
			}
		case *neteasePacket.InventoryContent:
			// 初始化
			allExist := true
			pks := []raknet_wrapper.MinecraftPacket[standardPacket.Packet]{}
			// 遍历物品变动表中的每个物品堆栈实例
			for slot, item := range pk.Content {
				// 如果物品的固定网络堆栈 ID 为 -1，
				// 说明当前物品未被更改
				if item.Stack.NetworkID == -1 {
					allExist = false
					continue
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
			// 针对副手物品的特殊处理，
			// 因为客户端看起来只接受有更改的物品
			if pk.WindowID == neteaseProtocol.WindowIDOffHand && allExist {
				newerPacket := standardPacket.InventoryContent{
					WindowID: pk.WindowID,
					Content:  make([]standardProtocol.ItemInstance, 1),
				}
				newerPacket.Content[0].Stack = packet_translate_struct.ConvertToStandardItemStack(pk.Content[0].Stack)
				newerPacket.Content[0].Stack.NBTData = map[string]any{"Happy2018new": "Liliya233"}
				sendCopy = append(
					sendCopy,
					[]raknet_wrapper.MinecraftPacket[standardPacket.Packet]{
						{Packet: &newerPacket},
						{Packet: packet_translator.TranslatorPool[standardPacket.IDInventoryContent].ToStandardPacket(pk)},
					}...,
				)
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
