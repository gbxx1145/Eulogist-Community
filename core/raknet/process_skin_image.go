package RaknetConnection

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"slices"
	"strings"

	_ "embed"
)

//go:embed default_skin_resource_patch.json
var defaultSkinResourcePatch []byte

//go:embed default_skin_geometry.json
var defaultSkinGeometry []byte

// 从 url 指定的网址下载文件，
// 并返回该文件的二进制形式
func DownloadFile(url string) (result []byte, err error) {
	// get http response
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("DownloadFile: %v", err)
	}
	defer httpResponse.Body.Close()
	// read image data
	result, err = io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("DownloadFile: %v", err)
	}
	// return
	return
}

/*
从 url 指定的网址下载文件，
并处理为有效的皮肤数据。

skinImageData 指代皮肤的 PNG 二进制形式，
skinData 指代皮肤的一维的密集像素矩阵，
skinGeometryData 指代皮肤的骨架信息，
skinWidth 和 skinHight 则分别指代皮肤的
宽度和高度。

TODO: 支持 4D 皮肤
*/
func ProcessURLToSkin(url string) (skin *Skin, err error) {
	skin = &Skin{}
	// download skin file from remote server
	res, err := DownloadFile(url)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// get skin data
	if len(res) >= 4 && bytes.Equal(res[0:4], []byte("PK\x03\x04")) {
		// TODO: 支持 4D 皮肤
		if err = ConvertZIPToSkin(skin, res); err != nil {
			return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
		}
	} else {
		skin.SkinImageData, skin.SkinGeometry = res, defaultSkinGeometry
	}
	// decode to image
	img, err := ConvertToPNG(skin.SkinImageData)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// encode to pixels and return
	skin.SkinPixels = img.(*image.NRGBA).Pix
	skin.SkinWidth, skin.SkinHight = img.Bounds().Dx(), img.Bounds().Dy()
	return
}

// 从 zipData 指代的 ZIP 二进制数据负载提取皮肤数据。
// skinImageData 代表皮肤的 PNG 二进制形式，
// skinGeometry 代表皮肤的骨架信息。
//
// TODO: 支持 4D 皮肤
func ConvertZIPToSkin(skin *Skin, zipData []byte) (err error) {
	// prepare
	skinImageBuffer := bytes.NewBuffer([]byte{})
	// create reader
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return fmt.Errorf("ConvertZIPToSkin: %v", err)
	}
	// set default resource patch
	skin.SkinResourcePatch = defaultSkinResourcePatch
	// find skin contents
	for _, file := range reader.File {
		// skin data
		if strings.HasSuffix(file.Name, ".png") && !strings.HasSuffix(file.Name, "_bloom.png") {
			r, err := file.Open()
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			defer r.Close()
			_, err = io.Copy(skinImageBuffer, r)
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
		}
		// skin geometry
		if strings.HasSuffix(file.Name, "geometry.json") {
			r, err := file.Open()
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			defer r.Close()
			geometryData, err := io.ReadAll(r)
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			ProcessGeometry(skin, geometryData)
		}
	}
	// return
	skin.SkinImageData = skinImageBuffer.Bytes()
	return
}

// 将 imageData 解析为 PNG 图片
func ConvertToPNG(imageData []byte) (image.Image, error) {
	buffer := bytes.NewBuffer(imageData)
	img, err := png.Decode(buffer)
	if err != nil {
		return nil, fmt.Errorf("ConvertToPNG: %v", err)
	}
	return img, nil
}

type SkinGeometryBone struct {
	Cubes         []any     `json:"cubes"`
	Name          string    `json:"name"`
	Parent        string    `json:"parent,omitempty"`
	Pivot         []float64 `json:"pivot"`
	RenderGroupID int       `json:"render_group_id,omitempty"`
	Rotation      []float64 `json:"rotation,omitempty"`
}

type SkinGeometry struct {
	Bones               []*SkinGeometryBone `json:"bones"`
	TextureHeight       int                 `json:"textureheight"`
	TextureWidth        int                 `json:"texturewidth"`
	VisibleBoundsHeight float64             `json:"visible_bounds_height,omitempty"`
	VisibleBoundsOffset []float64           `json:"visible_bounds_offset,omitempty"`
	VisibleBoundsWidth  float64             `json:"visible_bounds_width,omitempty"`
}

func ProcessGeometry(skin *Skin, rawData []byte) (err error) {
	/* Layer 1 */
	geometryMap := map[string]json.RawMessage{}
	if err = json.Unmarshal(rawData, &geometryMap); err != nil {
		return fmt.Errorf("ProcessGeometry: %v", err)
	}
	if len(geometryMap) != 1 {
		return fmt.Errorf("ProcessGeometry: invaild len in geometry map")
	}
	// setup resource patch and get geometry data
	var skinGeometry json.RawMessage
	var geometryName string
	for k, v := range geometryMap {
		geometryName = k
		skinGeometry = v
	}
	skin.SkinResourcePatch = bytes.ReplaceAll(
		skin.SkinResourcePatch,
		[]byte("geometry.humanoid.custom"),
		[]byte(geometryName),
	)
	/* Layer 2 */
	geometry := &SkinGeometry{}
	if err = json.Unmarshal(skinGeometry, geometry); err != nil {
		return fmt.Errorf("ProcessGeometry: %v", err)
	}
	// handle bones
	hasRoot := false
	renderGroupNames := []string{"leftArm", "rightArm"}
	for _, bone := range geometry.Bones {
		// setup parent
		switch bone.Name {
		case "waist", "leftLeg", "rightLeg":
			bone.Parent = "root"
		case "head":
			bone.Parent = "body"
		case "leftArm", "rightArm":
			bone.Parent = "body"
			bone.RenderGroupID = 1
		case "body":
			bone.Parent = "waist"
		case "root":
			hasRoot = true
		}
		// setup render group
		if slices.Contains(renderGroupNames, bone.Parent) {
			bone.RenderGroupID = 1
			renderGroupNames = append(renderGroupNames, bone.Name)
		}
	}
	if !hasRoot {
		geometry.Bones = append(geometry.Bones, &SkinGeometryBone{
			Name:  "root",
			Pivot: []float64{0, 0, 0},
		})
	}
	// return
	skin.SkinGeometry, _ = json.Marshal(map[string]any{
		"format_version": "1.8.0",
		geometryName:     geometry,
	})
	return
}
