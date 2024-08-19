package mc_server

import SkinProcess "Eulogist/core/tools/skin_process"

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
	m.playerSkin = &SkinProcess.Skin{}
}

// ...
func (m *MinecraftServer) GetPlayerSkin() *SkinProcess.Skin {
	return m.playerSkin
}

// ...
func (m *MinecraftServer) SetPlayerSkin(skin *SkinProcess.Skin) {
	m.playerSkin = skin
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
