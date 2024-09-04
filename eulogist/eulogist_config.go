package Eulogist

// 验证服务器地址
const (
	AuthServerAddressFastBuilder = "https://user.fastbuilder.pro"
	AuthServerAddressLiliya233   = "https://liliya233.uk"
)

// EulogistConfig 结构体定义了 Eulogist 的配置信息
type EulogistConfig struct {
	RentalServerCode     string `json:"rental_server_code"`     // 网易租赁服编号
	RentalServerPassword string `json:"rental_server_password"` // 该租赁服对应的密码
	FBToken              string `json:"fb_token"`               // ..
}

// LookUpAuthServerAddress 根据令牌查找认证服务器地址
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
