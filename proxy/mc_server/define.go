package mc_server

import (
	fb_client "Eulogist/core/fb_auth/mv4/client"
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	"Eulogist/core/minecraft/netease/protocol/login"
	"Eulogist/core/minecraft/netease/protocol/packet"
	raknet_wrapper "Eulogist/core/raknet/wrapper"
	"Eulogist/core/tools/skin_process"

	"github.com/google/uuid"
)

type MinecraftServer struct {
	fbClient              *fb_client.Client
	authResponse          *fb_client.AuthResponse
	getCheckNumEverPassed bool

	standardBedrockIdentity uuid.UUID

	identityData *login.IdentityData
	clientData   *login.ClientData

	neteaseUID string
	playerSkin *skin_process.Skin
	serverSkin *neteaseProtocol.Skin
	outfitInfo map[string]*int

	entityUniqueID  int64
	entityRuntimeID uint64

	persistenceData PersistenceData

	*raknet_wrapper.Raknet[packet.Packet]
}

// 描述单个实体的数据
type Entity struct {
	EntityType      string // 该实体的英文 ID
	EntityRuntimeID uint64 // 该实体的运行时 ID
	EntityUniqueID  int64  // 该实体的唯一 ID
}

// 当前用户的持久化数据
type PersistenceData struct {
	WorldEntity []*Entity // 已保存的实体数据
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
