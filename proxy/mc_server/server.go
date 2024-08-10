package mc_server

import SkinProcess "Eulogist/core/tools/skin_process"

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
func (m *MinecraftServer) GetEntityUniqueID() int64 {
	return m.entityUniqueID
}

// ...
func (m *MinecraftServer) SetEntityUniqueID(entityUniqueID int64) {
	m.entityUniqueID = entityUniqueID
}
