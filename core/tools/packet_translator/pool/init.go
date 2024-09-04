package packet_translate_pool

// ...
func init() {
	NetEasePacketIDToStandardPacketID = make(map[uint32]uint32)
	for key, value := range StandardPacketIDToNetEasePacketID {
		NetEasePacketIDToStandardPacketID[value] = key
	}

	NetEaseContainerIDStandardContainerID = make(map[uint8]uint8)
	for key, value := range StandardContainerIDToNetEaseContainerID {
		NetEaseContainerIDStandardContainerID[value] = key
	}
}
