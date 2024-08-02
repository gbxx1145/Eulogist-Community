package RaknetConnection

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"
	"os"

	_ "image/png"
)

// 从 url 指定的网址下载图片，
// 并返回该图片的二进制形式
func DownloadImage(url string) (result []byte, err error) {
	// get http response
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("DownloadImage: %v", err)
	}
	defer httpResponse.Body.Close()
	// read image data
	result, err = io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("DownloadImage: %v", err)
	}
	os.WriteFile("test.png", result, 0600)
	// return
	return
}

// 将 imageData 解析为 PNG 图片
func ConvertToPNG(imageData []byte) (image.Image, error) {
	buffer := bytes.NewBuffer(imageData)
	img, _, err := image.Decode(buffer)
	if err != nil {
		return nil, fmt.Errorf("ConvertToPNG: %v", err)
	}
	return img, nil
}

// ...
func ConvertUint32RGBAToBytes(r uint32, g uint32, b uint32, a uint32) []byte {
	if a == 0 {
		return []byte{0, 0, 0, 0}
	} else {
		return []byte{
			byte(r * 0xff / (a >> 8) >> 8),
			byte(g * 0xff / (a >> 8) >> 8),
			byte(b * 0xff / (a >> 8) >> 8),
			byte(a >> 8),
		}
	}
}

// 将 img 指代的图像编码为一维密集像素矩阵
func EncodeImageToBytes(img image.Image) []byte {
	result := make([]byte, 0)
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			color := img.At(x, y)
			R, G, B, A := color.RGBA()
			result = append(result, ConvertUint32RGBAToBytes(R, G, B, A)...)
		}
	}
	return result
}
