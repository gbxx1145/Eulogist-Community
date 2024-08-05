package RaknetConnection

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/login"
	"Eulogist/core/minecraft/protocol/packet"
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"gopkg.in/square/go-jose.v2/jwt"
)

// HandleNetworkSettings
// 接收 NetworkSettings 数据包并处理其内容。
//
// 这会为后续的数据包传输启用压缩，
// 然后，我们会构造并发送 Login 数据包至服务器
func (r *Raknet) HandleNetworkSettings(
	pk *packet.NetworkSettings,
	authResponse fbauth.AuthResponse,
	skin *Skin,
) (identityData *login.IdentityData, clientData *login.ClientData, err error) {
	// 准备
	var loginRequest []byte
	// 为底层 Raknet 连接启用数据包压缩
	alg, ok := packet.CompressionByID(pk.CompressionAlgorithm)
	if !ok {
		return nil, nil, fmt.Errorf("HandleNetworkSettings: unknown compression algorithm: %v", pk.CompressionAlgorithm)
	}
	r.encoder.EnableCompression(alg)
	r.decoder.EnableCompression(alg)
	// 编码登录请求
	loginRequest, identityData, clientData, err = r.EncodeLogin(authResponse, r.key, skin)
	if err != nil {
		return nil, nil, fmt.Errorf("HandleNetworkSettings: %v", err)
	}
	// 发送登录请求
	err = r.WritePacket(MinecraftPacket{
		Packet: &packet.Login{
			ClientProtocol:    protocol.CurrentProtocol,
			ConnectionRequest: loginRequest,
		},
	}, false)
	if err != nil {
		return nil, nil, fmt.Errorf("HandleNetworkSettings: %v", err)
	}
	// 返回值
	return
}

// HandleServerToClientHandshake
// 处理从服务器收到的 ServerToClientHandshake 包，
// 并为后续的数据传输启用加密
func (r *Raknet) HandleServerToClientHandshake(pk *packet.ServerToClientHandshake) error {
	// 解析 JWT 令牌
	tok, err := jwt.ParseSigned(string(pk.JWT))
	if err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: parse server token: %w", err)
	}
	// 获取公钥并进行解码
	raw, _ := tok.Headers[0].ExtraHeaders["x5u"].(string)
	pub := new(ecdsa.PublicKey)
	if err := login.ParsePublicKey(raw, pub); err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: parse server public key: %w", err)
	}
	// 验证并提取 Claims 和 Salt
	var c saltClaims
	if err := tok.Claims(pub, &c); err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: verify claims: %w", err)
	}
	c.Salt = strings.TrimRight(c.Salt, "=")
	salt, err := base64.RawStdEncoding.DecodeString(c.Salt)
	if err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: error base64 decoding ServerToClientHandshake salt: %v", err)
	}
	// 计算共享密钥
	x, _ := pub.Curve.ScalarMult(pub.X, pub.Y, r.key.D.Bytes())
	sharedSecret := append(bytes.Repeat([]byte{0}, 48-len(x.Bytes())), x.Bytes()...)
	// 创建加密密钥
	keyBytes := sha256.Sum256(append(salt, sharedSecret...))
	// 启用加密
	r.encoder.EnableEncryption(keyBytes)
	r.decoder.EnableEncryption(keyBytes)
	// 发送回应的 ClientToServerHandshake 包
	_ = r.WritePacket(MinecraftPacket{Packet: &packet.ClientToServerHandshake{}}, false)
	// 返回值
	return nil
}

// HandleStartGame 处理 StartGame 数据包，
// 用于表示玩家已加入游戏
func (r *Raknet) HandleStartGame(pk *packet.StartGame) (entityUniqueID int64) {
	entityUniqueID = pk.EntityUniqueID

	for _, item := range pk.Items {
		if item.Name == "minecraft:shield" {
			r.shieldID.Store(int32(item.RuntimeID))
		}
	}

	return entityUniqueID
}
