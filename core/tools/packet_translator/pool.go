package packet_translator

import (
	"Eulogist/core/tools/packet_translator/packet"

	standardPacket "Eulogist/core/minecraft/standard/protocol/packet"
)

// ...
var TranslatorPool = map[uint32]Translator{
	standardPacket.IDAddActor:                 &packet.AddActor{},
	standardPacket.IDAddEntity:                &packet.AddEntity{},
	standardPacket.IDAddPlayer:                &packet.AddPlayer{},
	standardPacket.IDAddVolumeEntity:          &packet.AddVolumeEntity{},
	standardPacket.IDAnimate:                  &packet.Animate{},
	standardPacket.IDChangeMobProperty:        &packet.ChangeMobProperty{},
	standardPacket.IDClientBoundDebugRenderer: &packet.ClientBoundDebugRenderer{},
	standardPacket.IDClientBoundMapItemData:   &packet.ClientBoundMapItemData{},
	standardPacket.IDCommandBlockUpdate:       &packet.CommandBlockUpdate{},
	standardPacket.IDCommandRequest:           &packet.CommandRequest{},
	standardPacket.IDCompletedUsingItem:       &packet.CompletedUsingItem{},
	standardPacket.IDContainerOpen:            &packet.ContainerOpen{},
	standardPacket.IDCraftingData:             &packet.CraftingData{},
	standardPacket.IDInventoryTransaction:     &packet.InventoryTransaction{},
	standardPacket.IDItemStackRequest:         &packet.ItemStackRequest{},
	standardPacket.IDOnScreenTextureAnimation: &packet.OnScreenTextureAnimation{},
	standardPacket.IDPlaySound:                &packet.PlaySound{},
	standardPacket.IDPlayerAuthInput:          &packet.PlayerAuthInput{},
	standardPacket.IDPlayerEnchantOptions:     &packet.PlayerEnchantOptions{},
	standardPacket.IDPlayerList:               &packet.PlayerList{},
	standardPacket.IDPlayerSkin:               &packet.PlayerSkin{},
	standardPacket.IDRemoveEntity:             &packet.RemoveEntity{},
	standardPacket.IDRemoveVolumeEntity:       &packet.RemoveVolumeEntity{},
	standardPacket.IDRequestChunkRadius:       &packet.RequestChunkRadius{},
	standardPacket.IDRequestPermissions:       &packet.RequestPermissions{},
	standardPacket.IDResourcePackStack:        &packet.ResourcePackStack{},
	standardPacket.IDSimpleEvent:              &packet.SimpleEvent{},
	standardPacket.IDStartGame:                &packet.StartGame{},
	standardPacket.IDSubChunk:                 &packet.SubChunk{},
	standardPacket.IDText:                     &packet.Text{},
	standardPacket.IDUpdateBlockSynced:        &packet.UpdateBlockSynced{},
	standardPacket.IDEducationSettings:        &packet.EducationSettings{},
	standardPacket.IDLessonProgress:           &packet.LessonProgress{},
}
