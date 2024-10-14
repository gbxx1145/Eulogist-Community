package raknet_wrapper

import (
	"Eulogist/core/minecraft/standard/protocol/packet"
	"context"
	"crypto/ecdsa"
	"sync/atomic"
)

// ...
func (r *Raknet[T]) Encoder() *packet.Encoder {
	return r.encoder
}

// ...
func (r *Raknet[T]) Decoder() *packet.Decoder {
	return r.decoder
}

// ...
func (r *Raknet[T]) Key() *ecdsa.PrivateKey {
	return r.key
}

// ...
func (r *Raknet[T]) Salt() []byte {
	return r.salt
}

// ...
func (r *Raknet[T]) ShieldID() *atomic.Int32 {
	return &r.shieldID
}

// ...
func (r *Raknet[T]) Context() context.Context {
	return r.context
}
