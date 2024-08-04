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

// 从 url 指定的网址下载文件，
// 并处理为有效的皮肤数据，
// 然后保存在 skin 中
func ProcessURLToSkin(url string) (skin *Skin, err error) {
	// 初始化，
	// 然后从远程服务器下载皮肤文件
	skin = &Skin{}
	res, err := DownloadFile(url)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// 获取皮肤数据
	{
		// 如果这是一个普通的皮肤，
		// 那么 res 就是该皮肤的 PNG 二进制形式，
		// 并且该皮肤使用的骨架格式为默认格式
		skin.SkinImageData, skin.SkinGeometry = res, defaultSkinGeometry
		// 如果这是一个高级的皮肤(比如 4D 皮肤)，
		// 那么 res 是一个压缩包，
		// 我们需要处理这个压缩包以得到皮肤文件
		if len(res) >= 4 && bytes.Equal(res[0:4], []byte("PK\x03\x04")) {
			if err = ConvertZIPToSkin(skin, res); err != nil {
				return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
			}
		}
	}
	// 将皮肤 PNG 二进制形式解码为图片
	img, err := ConvertToPNG(skin.SkinImageData)
	if err != nil {
		return nil, fmt.Errorf("ProcessURLToSkin: %v", err)
	}
	// 返回值
	skin.SkinPixels = img.(*image.NRGBA).Pix
	skin.SkinWidth, skin.SkinHight = img.Bounds().Dx(), img.Bounds().Dy()
	return
}

// 从 zipData 指代的 ZIP 二进制数据负载提取皮肤数据，
// 并把处理好的皮肤数据保存在 skin 中
func ConvertZIPToSkin(skin *Skin, zipData []byte) (err error) {
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
	// 返回值
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
