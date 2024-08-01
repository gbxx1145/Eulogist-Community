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
// 并且启用为底层 Raknet 连接启用数据包压缩功能。
func (r *RaknetConnection) HandleRequestNetworkSettings(pk *packet.RequestNetworkSettings) error {
	if pk.ClientProtocol != protocol.CurrentProtocol {
		status := packet.PlayStatusLoginFailedClient
		if pk.ClientProtocol > protocol.CurrentProtocol {
			// 此时服务器已过期，所以我们需要更新 status 的值
			status = packet.PlayStatusLoginFailedServer
		}
		_ = r.WritePacket(MinecraftPacket{Packet: &packet.PlayStatus{Status: status}}, false)
		return fmt.Errorf("HandleRequestNetworkSettings: Connected with an incompatible protocol: expected protocol = %v, client protocol = %v", protocol.CurrentProtocol, pk.ClientProtocol)
	}

	if err := r.WritePacket(MinecraftPacket{Packet: &packet.NetworkSettings{
		CompressionThreshold:    1,
		CompressionAlgorithm:    0,
		ClientThrottle:          false,
		ClientThrottleThreshold: 0,
		ClientThrottleScalar:    0,
	}}, false); err != nil {
		return fmt.Errorf("HandleRequestNetworkSettings: error sending network settings: %v", err)
	}

	r.encoder.EnableCompression(packet.DefaultCompression)
	r.decoder.EnableCompression(packet.DefaultCompression)
	return nil
}

// 处理传入的登录数据包。
// 它验证并解码数据包中找到的登录请求，
// 如果无法成功完成，则返回错误
func (r *RaknetConnection) HandleLogin(pk *packet.Login) error {
	var (
		err        error
		authResult login.AuthResult
	)
	_, _, authResult, err = login.Parse(pk.ConnectionRequest)
	if err != nil {
		return fmt.Errorf("HandleLogin: parse login request: %w", err)
	}

	if err := r.EnableEncryption(authResult.PublicKey); err != nil {
		return fmt.Errorf("HandleLogin: error enabling encryption: %v", err)
	}
	return nil
}

// 为创建的底层 Raknet 连接启用加密。
// 它向客户端发送未加密的握手数据包，然后启用加密。
func (r *RaknetConnection) EnableEncryption(clientPublicKey *ecdsa.PublicKey) error {
	signer, _ := jose.NewSigner(jose.SigningKey{Key: r.key, Algorithm: jose.ES384}, &jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]any{"x5u": login.MarshalPublicKey(&r.key.PublicKey)},
	})
	// We produce an encoded JWT using the header and payload above.
	// Then, we send the JWT in a ServerToClient-
	// Handshake packet so that the client can initialise encryption.
	serverJWT, err := jwt.Signed(signer).Claims(saltClaims{Salt: base64.RawStdEncoding.EncodeToString(r.salt)}).CompactSerialize()
	if err != nil {
		return fmt.Errorf("EnableEncryption: compact serialise server JWT: %w", err)
	}
	// get server JWT
	if err := r.WritePacket(MinecraftPacket{
		Packet: &packet.ServerToClientHandshake{JWT: []byte(serverJWT)},
	}, false); err != nil {
		return fmt.Errorf("EnableEncryption: error sending ServerToClientHandshake packet: %v", err)
	}
	// write server to client hand shake
	x, _ := clientPublicKey.Curve.ScalarMult(clientPublicKey.X, clientPublicKey.Y, r.key.D.Bytes())
	sharedSecret := append(bytes.Repeat([]byte{0}, 48-len(x.Bytes())), x.Bytes()...)
	keyBytes := sha256.Sum256(append(r.salt, sharedSecret...))
	// calculate the key bytes
	r.encoder.EnableEncryption(keyBytes)
	r.decoder.EnableEncryption(keyBytes)
	// enable encryption
	return nil
	// return
}
