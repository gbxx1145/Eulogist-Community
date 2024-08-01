package Eulogist

const (
	LaunchTypeNormal int = iota
	LaunchTypeNoOperation
)

const (
	AuthServerTypeFastBuilder int = iota
	AuthServerTypeLiliya233
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
	AuthServerType       int    `json:"auth_server_type"`
}

func DefaultEulogistConfig() EulogistConfig {
	return EulogistConfig{
		LaunchType:     LaunchTypeNormal,
		NEMCPath:       `D:\MCLDownload\MinecraftBENetease\windowsmc\Minecraft.Windows.exe`,
		ServerIP:       "127.0.0.1",
		ServerPort:     19132,
		AuthServerType: AuthServerTypeFastBuilder,
	}
}

func LookUpAuthServerAddress(id int) string {
	switch id {
	case AuthServerTypeFastBuilder:
		return AuthServerAddressFastBuilder
	case AuthServerTypeLiliya233:
		return AuthServerAddressLiliya233
	default:
		return AuthServerAddressFastBuilder
	}
}
