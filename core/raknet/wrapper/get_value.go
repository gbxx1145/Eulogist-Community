package raknet_wrapper

import (
	"Eulogist/core/minecraft/standard/protocol/packet"
	"context"
	"crypto/ecdsa"
	"sync/atomic"
)

// ...
func (r *Raknet[T]) GetEncoder() *packet.Encoder {
	return r.encoder
}

// ...
func (r *Raknet[T]) GetDecoder() *packet.Decoder {
	return r.decoder
}

// ...
func (r *Raknet[T]) GetKey() *ecdsa.PrivateKey {
	return r.key
}

// ...
func (r *Raknet[T]) GetSalt() []byte {
	return r.salt
}

// ...
func (r *Raknet[T]) GetShieldID() *atomic.Int32 {
	return &r.shieldID
}

// ...
func (r *Raknet[T]) GetContext() context.Context {
	return r.context
}
