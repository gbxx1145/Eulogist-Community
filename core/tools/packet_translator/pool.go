package packet_translator

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

var StandardPacketIDToNetEasePacketID = map[uint32]uint32{
	standardPacket.IDLogin:                      neteasePacket.IDLogin,
	standardPacket.IDPlayStatus:                 neteasePacket.IDPlayStatus,
	standardPacket.IDServerToClientHandshake:    neteasePacket.IDServerToClientHandshake,
	standardPacket.IDClientToServerHandshake:    neteasePacket.IDClientToServerHandshake,
	standardPacket.IDDisconnect:                 neteasePacket.IDDisconnect,
	standardPacket.IDResourcePacksInfo:          neteasePacket.IDResourcePacksInfo,
	standardPacket.IDResourcePackStack:          neteasePacket.IDResourcePackStack,
	standardPacket.IDResourcePackClientResponse: neteasePacket.IDResourcePackClientResponse,
	standardPacket.IDText:                       neteasePacket.IDText,
	standardPacket.IDSetTime:                    neteasePacket.IDSetTime,
	standardPacket.IDStartGame:                  neteasePacket.IDStartGame,
	standardPacket.IDAddPlayer:                  neteasePacket.IDAddPlayer,
	standardPacket.IDAddActor:                   neteasePacket.IDAddActor,
	standardPacket.IDRemoveActor:                neteasePacket.IDRemoveActor,
	standardPacket.IDAddItemActor:               neteasePacket.IDAddItemActor, // ...
	// ---
	standardPacket.IDTakeItemActor:     neteasePacket.IDTakeItemActor,
	standardPacket.IDMoveActorAbsolute: neteasePacket.IDMoveActorAbsolute,
	standardPacket.IDMovePlayer:        neteasePacket.IDMovePlayer,
	standardPacket.IDPassengerJump:     neteasePacket.IDPassengerJump,
	standardPacket.IDUpdateBlock:       neteasePacket.IDUpdateBlock,
	standardPacket.IDAddPainting:       neteasePacket.IDAddPainting,
	standardPacket.IDTickSync:          neteasePacket.IDTickSync,
	// ---
	standardPacket.IDLevelEvent:           neteasePacket.IDLevelEvent,
	standardPacket.IDMobEffect:            neteasePacket.IDMobEffect,
	standardPacket.IDUpdateAttributes:     neteasePacket.IDUpdateAttributes,
	standardPacket.IDInventoryTransaction: neteasePacket.IDInventoryTransaction,
	standardPacket.IDMobEquipment:         neteasePacket.IDMobEquipment, // ...
	standardPacket.IDMobArmourEquipment:   neteasePacket.IDMobArmourEquipment,
	standardPacket.IDInteract:             neteasePacket.IDInteract,
	standardPacket.IDBlockPickRequest:     neteasePacket.IDBlockPickRequest,
	standardPacket.IDActorPickRequest:     neteasePacket.IDActorPickRequest,
	standardPacket.IDPlayerAction:         neteasePacket.IDPlayerAction,
	// ---
	standardPacket.IDHurtArmour:       neteasePacket.IDHurtArmour,
	standardPacket.IDSetActorData:     neteasePacket.IDSetActorData,
	standardPacket.IDSetActorMotion:   neteasePacket.IDSetActorMotion,
	standardPacket.IDSetActorLink:     neteasePacket.IDSetActorLink,
	standardPacket.IDSetHealth:        neteasePacket.IDSetHealth,
	standardPacket.IDSetSpawnPosition: neteasePacket.IDSetSpawnPosition,
	standardPacket.IDAnimate:          neteasePacket.IDAnimate,
	standardPacket.IDRespawn:          neteasePacket.IDRespawn,
	standardPacket.IDContainerOpen:    neteasePacket.IDContainerOpen,
	standardPacket.IDContainerClose:   neteasePacket.IDContainerClose,
	standardPacket.IDPlayerHotBar:     neteasePacket.IDPlayerHotBar,
	standardPacket.IDInventoryContent: neteasePacket.IDInventoryContent,
	standardPacket.IDInventorySlot:    neteasePacket.IDInventorySlot,
	standardPacket.IDCraftingData:     neteasePacket.IDCraftingData,
	standardPacket.IDMobEquipment:     neteasePacket.IDMobEquipment,
	standardPacket.IDMobEquipment:     neteasePacket.IDMobEquipment,
	standardPacket.IDMobEquipment:     neteasePacket.IDMobEquipment,
}
