package packet

import (
	neteasePacket "Eulogist/core/minecraft/protocol/packet"

	standardPacket "Eulogist/core/standard/protocol/packet"
)

type RequestPermissions struct{}

func (pk *RequestPermissions) ToNetEasePacket(standard standardPacket.Packet) neteasePacket.Packet {
	p := neteasePacket.RequestPermissions{}
	input := standard.(*standardPacket.RequestPermissions)

	p.EntityUniqueID = input.EntityUniqueID
	p.PermissionLevel = int32(input.PermissionLevel)
	p.RequestedPermissions = input.RequestedPermissions

	return &p
}

func (pk *RequestPermissions) ToStandardPacket(netease neteasePacket.Packet) standardPacket.Packet {
	p := standardPacket.RequestPermissions{}
	input := netease.(*neteasePacket.RequestPermissions)

	p.EntityUniqueID = input.EntityUniqueID
	p.PermissionLevel = uint8(input.PermissionLevel)
	p.RequestedPermissions = input.RequestedPermissions

	return &p
}
