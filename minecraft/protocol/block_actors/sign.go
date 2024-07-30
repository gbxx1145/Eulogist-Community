package block_actors

import general "Eulogist/minecraft/protocol/block_actors/general_actors"

// 告示牌
type Sign struct {
	general.SignBlockActor `mapstructure:",squash"`
}

// ID ...
func (*Sign) ID() string {
	return IDSign
}
