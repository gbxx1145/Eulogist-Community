package RaknetConnection

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
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
	// 获取 HTTP 响应
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("DownloadFile: %v", err)
	}
	defer httpResponse.Body.Close()
	// 读取文件数据
	result, err = io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("DownloadFile: %v", err)
	}
	// 返回值
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
*/
func ProcessURLToSkin(url string) (skin *Skin, err error) {
	skin = &Skin{}
	// 从远程服务器下载皮肤文件
	res, err := DownloadFile(url)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// 获取皮肤数据
	skin.SkinImageData, skin.SkinGeometry = res, defaultSkinGeometry
	if len(res) >= 4 && bytes.Equal(res[0:4], []byte("PK\x03\x04")) {
		// TODO: 支持 4D 皮肤
		// 将 ZIP 文件转换为皮肤数据
		if err = ConvertZIPToSkin(skin, res); err != nil {
			return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
		}
	}
	// 将皮肤数据解码为图片
	img, err := ConvertToPNG(skin.SkinImageData)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// 将图片编码为像素并返回
	skin.SkinPixels = img.(*image.NRGBA).Pix
	skin.SkinWidth, skin.SkinHight = img.Bounds().Dx(), img.Bounds().Dy()
	// 返回值
	return
}

// 从 zipData 指代的 ZIP 二进制数据负载提取皮肤数据。
// skinImageData 代表皮肤的 PNG 二进制形式，
// skinGeometry 代表皮肤的骨架信息。
//
// TODO: 支持 4D 皮肤
func ConvertZIPToSkin(skin *Skin, zipData []byte) (err error) {
	// 准备缓冲区
	skinImageBuffer := bytes.NewBuffer([]byte{})
	skinGeometryBuffer := bytes.NewBuffer([]byte{})
	// 创建 ZIP 读取器
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return fmt.Errorf("ConvertZIPToSkin: %v", err)
	}
	// 设置默认资源路径
	skin.SkinResourcePatch = defaultSkinResourcePatch
	// 查找皮肤内容
	for _, file := range reader.File {
		// 皮肤数据
		if strings.HasSuffix(file.Name, ".png") && !strings.HasSuffix(file.Name, "_bloom.png") {
			r, err := file.Open()
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			defer r.Close()
			skin.SkinImageData, err = io.ReadAll(r)
			if err != nil {
				return fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
		}
		// 皮肤骨架信息
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
	// 返回皮肤数据
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
