package mc_server

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/tools/skin_process"

	"github.com/google/uuid"
)

// ...
func (m *MinecraftServer) SetStandardBedrockIdentity(standardBedrockIdentity uuid.UUID) {
	m.standardBedrockIdentity = standardBedrockIdentity
}

// ...
func (m *MinecraftServer) GetStandardBedrockIdentity() uuid.UUID {
	return m.standardBedrockIdentity
}

// ...
func (m *MinecraftServer) GetNeteaseUID() string {
	return m.neteaseUID
}

// ...
func (m *MinecraftServer) SetNeteaseUID(neteaseUID string) {
	m.neteaseUID = neteaseUID
}

// ...
func (m *MinecraftServer) InitPlayerSkin() {
	m.playerSkin = &skin_process.Skin{}
}

// ...
func (m *MinecraftServer) GetPlayerSkin() *skin_process.Skin {
	return m.playerSkin
}

// ...
func (m *MinecraftServer) SetPlayerSkin(skin *skin_process.Skin) {
	m.playerSkin = skin
}

// ...
func (m *MinecraftServer) GetServerSkin() *neteaseProtocol.Skin {
	return m.serverSkin
}

// ...
func (m *MinecraftServer) SetServerSkin(serverSkin *neteaseProtocol.Skin) {
	m.serverSkin = serverSkin
}

// ...
func (m *MinecraftServer) GetOutfitInfo() map[string]*int {
	return m.outfitInfo
}

// ...
func (m *MinecraftServer) SetOutfitInfo(outfitInfo map[string]*int) {
	m.outfitInfo = outfitInfo
}

// ...
func (m *MinecraftServer) GetEntityUniqueID() int64 {
	return m.entityUniqueID
}

// ...
func (m *MinecraftServer) SetEntityUniqueID(entityUniqueID int64) {
	m.entityUniqueID = entityUniqueID
}

// ...
func (m *MinecraftServer) GetEntityRuntimeID() uint64 {
	return m.entityRuntimeID
}

// ...
func (m *MinecraftServer) SetEntityRuntimeID(entityRuntimeID uint64) {
	m.entityRuntimeID = entityRuntimeID
}
