package persistence_data

import (
	neteaseProtocol "Eulogist/core/minecraft/netease/protocol"
	standardProtocol "Eulogist/core/minecraft/standard/protocol"
	"Eulogist/core/tools/skin_process"
)

// 用户的皮肤数据
type SkinData struct {
	NeteaseSkin *skin_process.Skin     // 赞颂者处理的皮肤结果
	ServerSkin  *neteaseProtocol.Skin  // 租赁服返回的最终皮肤信息
	ClientSkin  *standardProtocol.Skin // 客户端原本提交的皮肤信息
}
