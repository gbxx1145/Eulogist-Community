package RaknetConnection

import (
	fbauth "Eulogist/core/fb_auth/pv4"
	"Eulogist/core/minecraft/protocol"
	"Eulogist/core/minecraft/protocol/login"
	"bytes"
	"crypto/ecdsa"
	_ "embed"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

//go:embed skin_resource_patch.json
var skinResourcePatch []byte

//go:embed skin_geometry.json
var skinGeometry []byte

// ...
func (r *RaknetConnection) EncodeLogin(
	authResponse fbauth.AuthResponse, clientKey *ecdsa.PrivateKey,
) ([]byte, error) {
	identityData := login.IdentityData{}
	clientData := login.ClientData{}

	defaultIdentityData(&identityData)
	defaultClientData(&clientData, authResponse)

	var request []byte
	// We login as an Android device and this will show up in the 'titleId' field in the JWT chain, which
	// we can't edit. We just enforce Android data for logging in.
	setAndroidData(&clientData)

	request = login.Encode(authResponse.ChainInfo, clientData, clientKey)
	identityData, _, _, err := login.Parse(request)
	if err != nil {
		return nil, fmt.Errorf("EncodeLogin: WARNING: Identity data parsing error: %v", err)
	}

	return request, nil
}

// defaultIdentityData edits the IdentityData passed to have defaults set to all fields that were left
// unchanged.
func defaultIdentityData(data *login.IdentityData) {
	if data.Identity == "" {
		data.Identity = uuid.New().String()
	}
	if data.DisplayName == "" {
		data.DisplayName = "Steve"
	}
}

// defaultClientData edits the ClientData passed to have defaults set to all fields that were left unchanged.
func defaultClientData(d *login.ClientData, authResponse fbauth.AuthResponse) {
	rand.Seed(time.Now().Unix())

	d.ServerAddress = authResponse.RentalServerIP
	d.ThirdPartyName = authResponse.BotName
	if d.DeviceOS == 0 {
		d.DeviceOS = protocol.DeviceAndroid
	}
	if d.GameVersion == "" {
		d.GameVersion = protocol.CurrentVersion
	}

	// PhoenixBuilder specific changes.
	// Author: Liliya233, Happy2018new
	if d.GrowthLevel != authResponse.BotLevel {
		d.GrowthLevel = authResponse.BotLevel
	}

	if d.ClientRandomID == 0 {
		d.ClientRandomID = rand.Int63()
	}
	if d.DeviceID == "" {
		d.DeviceID = uuid.New().String()
	}
	if d.LanguageCode == "" {
		// PhoenixBuilder specific changes.
		// Author: Liliya233
		d.LanguageCode = "zh_CN"
		// d.LanguageCode = "en_GB"
	}
	if d.AnimatedImageData == nil {
		d.AnimatedImageData = make([]login.SkinAnimation, 0)
	}
	if d.PersonaPieces == nil {
		d.PersonaPieces = make([]login.PersonaPiece, 0)
	}
	if d.PieceTintColours == nil {
		d.PieceTintColours = make([]login.PersonaPieceTintColour, 0)
	}
	if d.SelfSignedID == "" {
		d.SelfSignedID = uuid.New().String()
	}
	if d.SkinID == "" {
		d.SkinID = uuid.New().String()
	}
	if d.SkinData == "" {
		d.SkinData = base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{0, 0, 0, 255}, 32*64))
		d.SkinImageHeight = 32
		d.SkinImageWidth = 64
	}
	if d.SkinResourcePatch == "" {
		d.SkinResourcePatch = base64.StdEncoding.EncodeToString(skinResourcePatch)
	}
	if d.SkinGeometry == "" {
		d.SkinGeometry = base64.StdEncoding.EncodeToString(skinGeometry)
	}
}

// setAndroidData ensures the login.ClientData passed matches settings you would see on an Android device.
func setAndroidData(data *login.ClientData) {
	data.DeviceOS = protocol.DeviceAndroid
	data.GameVersion = protocol.CurrentVersion
}
