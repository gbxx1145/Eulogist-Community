package mc_server

import RaknetConnection "Eulogist/core/raknet"

// ...
func (m *MinecraftServer) InitPlayerSkin() {
	m.playerSkin = &RaknetConnection.Skin{}
}

// ...
func (m *MinecraftServer) GetPlayerSkin() *RaknetConnection.Skin {
	return m.playerSkin
}

// ...
func (m *MinecraftServer) SetPlayerSkin(skin *RaknetConnection.Skin) {
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
