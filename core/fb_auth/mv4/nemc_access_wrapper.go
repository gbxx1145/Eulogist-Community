package fbauth

import (
	"context"
	"encoding/base64"
)

type AccessWrapper struct {
	ServerCode     string
	ServerPassword string
	Token          string
	Client         *Client
	Username       string
	Password       string
}

func NewAccessWrapper(Client *Client, ServerCode, ServerPassword, Token, username, password string) *AccessWrapper {
	return &AccessWrapper{
		Client:         Client,
		ServerCode:     ServerCode,
		ServerPassword: ServerPassword,
		Token:          Token,
		Username:       username,
		Password:       password,
	}
}

func (aw *AccessWrapper) GetAccess(ctx context.Context, publicKey []byte) (authResponse AuthResponse, err error) {
	pubKeyData := base64.StdEncoding.EncodeToString(publicKey)
	authResponse, err = aw.Client.Auth(ctx, aw.ServerCode, aw.ServerPassword, pubKeyData, aw.Token, aw.Username, aw.Password)
	if err != nil {
		return AuthResponse{}, err
	}
	return
}
