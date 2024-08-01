package Eulogist

const (
	LaunchTypeNormal int = iota
	LaunchTypeNoOperation
)

const (
	AuthServerAddressFastBuilder = "https://user.fastbuilder.pro"
	AuthServerAddressLiliya233   = "https://liliya233.uk"
)

type EulogistConfig struct {
	LaunchType int    `json:"launch_type"`
	NEMCPath   string `json:"nemc_program_path"`
	SkinPath   string `json:"skin_path"`

	ServerIP   string `json:"server_ip"`
	ServerPort int    `json:"server_port"`

	RentalServerCode     string `json:"rental_server_code"`
	RentalServerPassword string `json:"rental_server_password"`
	FBToken              string `json:"fb_token"`
}

func DefaultEulogistConfig() EulogistConfig {
	return EulogistConfig{
		LaunchType: LaunchTypeNormal,
		NEMCPath:   `D:\MCLDownload\MinecraftBENetease\windowsmc\Minecraft.Windows.exe`,
		ServerIP:   "127.0.0.1",
		ServerPort: 19132,
	}
}

func LookUpAuthServerAddress(token string) string {
	if len(token) < 3 {
		return AuthServerAddressFastBuilder
	}

	switch token[:3] {
	case "w9/":
		return AuthServerAddressFastBuilder
	case "y8/":
		return AuthServerAddressLiliya233
	default:
		return AuthServerAddressFastBuilder
	}
}
