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

// 处理来自服务器的 HandleNetworkSettings 数据包，
// 并为后续数据包的传递启用压缩功能，
// 然后发送 Login 数据包
func (r *Raknet) HandleNetworkSettings(
	pk *packet.NetworkSettings, authResponse fbauth.AuthResponse,
) error {
	alg, ok := packet.CompressionByID(pk.CompressionAlgorithm)
	if !ok {
		return fmt.Errorf("HandleNetworkSettings: unknown compression algorithm: %v", pk.CompressionAlgorithm)
	}
	// get compression
	r.encoder.EnableCompression(alg)
	r.decoder.EnableCompression(alg)
	// enable compression
	loginRequest, err := r.EncodeLogin(authResponse, r.key)
	if err != nil {
		return fmt.Errorf("HandleNetworkSettings: %v", err)
	}
	err = r.WritePacket(MinecraftPacket{
		Packet: &packet.Login{
			ClientProtocol:    protocol.CurrentProtocol,
			ConnectionRequest: loginRequest,
		},
	}, false)
	if err != nil {
		return fmt.Errorf("HandleNetworkSettings: %v", err)
	}
	// send login packet
	return nil
	// return
}

// HandleServerToClientHandshake 处理传入的 ServerToClientHandshake 包。
// 这会使当前与服务器建立的 Raknet 会话被加密。
// 我们通过使用数据包中公开的服务器的散列和公钥来完成通信加密
func (r *Raknet) HandleServerToClientHandshake(pk *packet.ServerToClientHandshake) error {
	tok, err := jwt.ParseSigned(string(pk.JWT))
	if err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: parse server token: %w", err)
	}
	//lint:ignore S1005 Double assignment is done explicitly to prevent panics.
	raw, _ := tok.Headers[0].ExtraHeaders["x5u"]
	kStr, _ := raw.(string)

	pub := new(ecdsa.PublicKey)
	if err := login.ParsePublicKey(kStr, pub); err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: parse server public key: %w", err)
	}

	var c saltClaims
	if err := tok.Claims(pub, &c); err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: verify claims: %w", err)
	}
	c.Salt = strings.TrimRight(c.Salt, "=")
	salt, err := base64.RawStdEncoding.DecodeString(c.Salt)
	if err != nil {
		return fmt.Errorf("HandleServerToClientHandshake: error base64 decoding ServerToClientHandshake salt: %v", err)
	}

	x, _ := pub.Curve.ScalarMult(pub.X, pub.Y, r.key.D.Bytes())
	// Make sure to pad the shared secret up to 96 bytes.
	sharedSecret := append(bytes.Repeat([]byte{0}, 48-len(x.Bytes())), x.Bytes()...)

	keyBytes := sha256.Sum256(append(salt, sharedSecret...))

	// Finally we enable encryption for the enc and dec using the secret pubKey bytes we produced.
	r.encoder.EnableEncryption(keyBytes)
	r.decoder.EnableEncryption(keyBytes)

	// We write a ClientToServerHandshake packet (which has no payload) as a response.
	_ = r.WritePacket(MinecraftPacket{Packet: &packet.ClientToServerHandshake{}}, false)
	return nil
}

// HandleStartGame 处理传入的 StartGame 数据包。
// 这是玩家已被添加到世界的信号，并且它获得了大部分专用属性
func (r *Raknet) HandleStartGame(pk *packet.StartGame) (entityUniqueID int64) {
	entityUniqueID = pk.EntityUniqueID

	for _, item := range pk.Items {
		if item.Name == "minecraft:shield" {
			r.shieldID.Store(int32(item.RuntimeID))
		}
	}

	return
}
