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

//go:embed skin_resource_patch.json
var skinResourcePatch []byte

//go:embed skin_geometry.json
var skinGeometry []byte

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

// 从 url 指定的网址下载文件，
// 并处理为有效的皮肤数据。
//
// skinData 指代皮肤的一维的密集像素矩阵，
// skinGeometryData 指代皮肤的骨架信息，
// skinWidth 和 skinHight 则分别指代皮肤的
// 宽度和高度。
//
// TODO: 支持 4D 皮肤
func ProcessFileToSkin(url string) (
	skinData []byte, skinGeometryData []byte,
	skinWidth int, skinHight int,
	err error,
) {
	// prepare
	var skinImageData []byte
	// download skin file from remote server
	res, err := DownloadFile(url)
	if err != nil {
		return nil, nil, 0, 0, fmt.Errorf("ProcessFileToSkin: %v", err)
	}
	// get skin data
	if len(res) >= 4 && bytes.Equal(res[0:4], []byte("PK\x03\x04")) {
		// TODO: 支持 4D 皮肤
		{
			// skinImageData, skinGeometryData, err = ConvertZIPToSkin(res)
			skinImageData, _, err = ConvertZIPToSkin(res)
			if err != nil {
				return nil, nil, 0, 0, fmt.Errorf("ProcessFileToSkin: %v", err)
			}
			skinGeometryData = skinGeometry
		}
	} else {
		skinImageData, skinGeometryData = res, skinGeometry
	}
	// decode to image
	img, err := ConvertToPNG(skinImageData)
	if err != nil {
		return nil, nil, 0, 0, fmt.Errorf("ProcessFileToSkin: %v", err)
	}
	// encode to pixels and return
	return EncodeImageToBytes(img), skinGeometryData, img.Bounds().Dx(), img.Bounds().Dy(), nil
}

// 从 zipData 指代的 ZIP 二进制数据负载提取皮肤数据。
// skinImageData 代表皮肤的 PNG 二进制形式，
// skinGeometry 代表皮肤的骨架信息。
//
// TODO: 支持 4D 皮肤
func ConvertZIPToSkin(zipData []byte) (skinImageData []byte, skinGeometryData []byte, err error) {
	// prepare
	skinImageBuffer := bytes.NewBuffer([]byte{})
	skinGeometryBuffer := bytes.NewBuffer([]byte{})
	// create reader
	reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
	}
	// find skin contents
	for _, file := range reader.File {
		// skin data
		if strings.Contains(file.Name, ".png") {
			r, err := file.Open()
			if err != nil {
				return nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			defer func() {
				err = r.Close()
				if err != nil {
					skinImageData, skinGeometryData, err = nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
				}
			}()
			_, err = io.Copy(skinImageBuffer, r)
			if err != nil {
				return nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
		}
		// skin geometry
		if strings.Contains(file.Name, "geometry.json") {
			r, err := file.Open()
			if err != nil {
				return nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
			defer func() {
				err = r.Close()
				if err != nil {
					skinImageData, skinGeometryData, err = nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
				}
			}()
			_, err = io.Copy(skinGeometryBuffer, r)
			if err != nil {
				return nil, nil, fmt.Errorf("ConvertZIPToSkin: %v", err)
			}
		}
	}
	// return
	return skinImageBuffer.Bytes(), skinGeometryBuffer.Bytes(), nil
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

// 将 img 指代的图像编码为一维密集像素矩阵
func EncodeImageToBytes(img image.Image) []byte {
	result := make([]byte, 0)
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pix := img.(*image.NRGBA).Pix
			index := (y*dy + x) * 4
			result = append(
				result,
				pix[index], pix[index+1], pix[index+2], pix[index+3],
			)
		}
	}
	return result
}
