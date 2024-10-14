package persistence_data

import (
	neteaseLogin "Eulogist/core/minecraft/netease/protocol/login"
	standardLogin "Eulogist/core/minecraft/standard/protocol/login"
)

// 来自 Minecraft 客户端的登录数据
type LoginDataClientSide struct {
	IdentityData *standardLogin.IdentityData // Minecraft 客户端的身份证明
	ClientData   *standardLogin.ClientData   // Minecraft 客户端的客户端数据
}

// 来自网易账户的登录数据
type LoginDataServerSide struct {
	IdentityData *neteaseLogin.IdentityData // 网易账户的身份证明
	ClientData   *neteaseLogin.ClientData   // 网易账户的客户端数据
}

// 记录 Minecraft 客户端和网易账户的登录数据
type LoginData struct {
	Client          LoginDataClientSide // 来自 Minecraft 客户端的登录数据
	Server          LoginDataServerSide // 来自网易账户的登录数据
	PlayerUniqueID  int64               // 当前网易账户在当前租赁服所对应的唯一 ID
	PlayerRuntimeID uint64              // 当前网易账户在当前租赁服所对应的运行时 ID
}
