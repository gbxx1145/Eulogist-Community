package RaknetConnection

import (
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/login"
	"Eulogist/core/minecraft/protocol/packet"
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// 处理传入的 RequestNetworkSettings 数据包。
// 如果不支持协议版本，它将返回错误。
// 否则，以 NetworkSettings 作为响应，
// 并且为底层 Raknet 连接启用数据包压缩
func (r *Raknet) HandleRequestNetworkSettings(pk *packet.RequestNetworkSettings) error {
	// 检查网络协议版本
	if pk.ClientProtocol != protocol.CurrentProtocol {
		status := packet.PlayStatusLoginFailedClient
		if pk.ClientProtocol > protocol.CurrentProtocol {
			// 此时服务器已过期，所以我们需要更新 status 的值
			status = packet.PlayStatusLoginFailedServer
		}
		_ = r.WritePacket(MinecraftPacket{Packet: &packet.PlayStatus{Status: status}}, false)
		return fmt.Errorf("HandleRequestNetworkSettings: Connected with an incompatible protocol: expected protocol = %v, client protocol = %v", protocol.CurrentProtocol, pk.ClientProtocol)
	}
	// 发送 NetworkSettings 数据包以响应客户端
	if err := r.WritePacket(MinecraftPacket{Packet: &packet.NetworkSettings{
		CompressionThreshold:    1,
		CompressionAlgorithm:    0,
		ClientThrottle:          false,
		ClientThrottleThreshold: 0,
		ClientThrottleScalar:    0,
	}}, false); err != nil {
		return fmt.Errorf("HandleRequestNetworkSettings: error sending network settings: %v", err)
	}
	// 为数据包传输启用压缩
	r.encoder.EnableCompression(packet.DefaultCompression)
	r.decoder.EnableCompression(packet.DefaultCompression)
	// 返回值
	return nil
}

// 处理传入的登录数据包。
// 它验证并解码数据包中找到的登录请求，
// 如果无法成功完成，则返回错误
func (r *Raknet) HandleLogin(pk *packet.Login) error {
	// 准备
	var (
		err        error
		authResult login.AuthResult
	)
	// 解析登录请求
	_, _, authResult, err = login.Parse(pk.ConnectionRequest)
	if err != nil {
		return fmt.Errorf("HandleLogin: parse login request: %w", err)
	}
	// 启用加密
	if err := r.EnableEncryption(authResult.PublicKey); err != nil {
		return fmt.Errorf("HandleLogin: error enabling encryption: %v", err)
	}
	// 返回值
	return nil
}

// 为创建的底层 Raknet 连接启用加密。
// 它向客户端发送未加密的握手数据包，
// 然后为底层 Raknet 连接启用数据包加密
func (r *Raknet) EnableEncryption(clientPublicKey *ecdsa.PublicKey) error {
	// 创建 JWT 签名器
	signer, _ := jose.NewSigner(jose.SigningKey{Key: r.key, Algorithm: jose.ES384}, &jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]any{"x5u": login.MarshalPublicKey(&r.key.PublicKey)},
	})
	// 生成并序列化 JWT
	serverJWT, err := jwt.Signed(signer).Claims(saltClaims{Salt: base64.RawStdEncoding.EncodeToString(r.salt)}).CompactSerialize()
	if err != nil {
		return fmt.Errorf("EnableEncryption: compact serialise server JWT: %w", err)
	}
	// 发送 ServerToClientHandshake 数据包
	if err := r.WritePacket(MinecraftPacket{
		Packet: &packet.ServerToClientHandshake{JWT: []byte(serverJWT)},
	}, false); err != nil {
		return fmt.Errorf("EnableEncryption: error sending ServerToClientHandshake packet: %v", err)
	}
	// 计算公钥
	x, _ := clientPublicKey.Curve.ScalarMult(clientPublicKey.X, clientPublicKey.Y, r.key.D.Bytes())
	sharedSecret := append(bytes.Repeat([]byte{0}, 48-len(x.Bytes())), x.Bytes()...)
	keyBytes := sha256.Sum256(append(r.salt, sharedSecret...))
	// 为数据包传输启用加密
	r.encoder.EnableEncryption(keyBytes)
	r.decoder.EnableEncryption(keyBytes)
	// 返回值
	return nil
}
