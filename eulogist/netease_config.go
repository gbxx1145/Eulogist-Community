package Eulogist

type NetEaseConfig struct {
	RoomInfo   RoomInfo   `json:"room_info"`
	PlayerInfo PlayerInfo `json:"player_info"`
	SkinInfo   SkinInfo   `json:"skin_info"`
	Misc       Misc       `json:"misc"`
}

type RoomInfo struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type PlayerInfo struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Urs      string `json:"urs"`
}

type SkinInfo struct {
	SkinPath string `json:"skin"`
	Slim     bool   `json:"slim"`
}

type Misc struct {
	MultiplayerGameType int `json:"multiplayer_game_type"`
}

func DefaultNetEaseConfig() NetEaseConfig {
	return NetEaseConfig{
		RoomInfo: RoomInfo{IP: "127.0.0.1", Port: 19132},
		PlayerInfo: PlayerInfo{
			UserID:   -1,
			UserName: "Eulogist",
			Urs:      "***",
		},
		Misc: Misc{MultiplayerGameType: 1000},
	}
}
