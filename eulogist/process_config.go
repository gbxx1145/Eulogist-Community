package Eulogist

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

// 检查 path 对应路径的文件是否存在。
// 如果不存在，或该路径指向一个文件夹，
// 则返回假，否则返回真
func FileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !fileInfo.IsDir()
}

// 将 content 以 JSON 形式写入到 path 指代的文件处
func WriteJsonFile(path string, content any) error {
	contentBytes, _ := json.Marshal(content)
	// marshal
	buffer := bytes.NewBuffer([]byte{})
	json.Indent(buffer, contentBytes, "", "	")
	// indent json
	err := os.WriteFile(path, buffer.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("WriteJsonFile: %v", err)
	}
	// write json to file
	return nil
	// return
}

// 在当前目录读取 赞颂者 的配置文件。
// 如果没有对应的文件，
// 则将尝试生成默认配置文件。
//
// 生成默认配置文件期间需要从控制台读取用户输入，
// 读取的内容包括需要赞颂的租赁服号及其密码，
// 以及 FastBuilder 原生验证服务器的 Token
func ReadEulogistConfig() (*EulogistConfig, error) {
	var cfg EulogistConfig

	if !FileExist("eulogist_config.json") {
		config, err := GenerateEulogistConfig()
		if err != nil {
			return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
		}
		return config, nil
	}

	fileBytes, err := os.ReadFile("eulogist_config.json")
	if err != nil {
		return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
	}

	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("ReadEulogistConfig: %v", err)
	}

	return &cfg, nil
}

// 在当前目录生成 赞颂者 的默认配置文件，
// 并返回该配置文件。
//
// 此函数会从控制台读取用户输入，
// 读取的内容包括需要赞颂的租赁服号及其密码，
// 以及 FastBuilder 原生验证服务器的 Token
func GenerateEulogistConfig() (config *EulogistConfig, err error) {
	cfg := DefaultEulogistConfig()

	pterm.Info.Printf("Type your rental server code: ")
	fmt.Scanln(&cfg.RentalServerCode)

	pterm.Info.Printf("Type your rental server password: ")
	fmt.Scanln(&cfg.RentalServerPassword)

	pterm.Info.Printf("Type your FB token of FastBuilder Auth Server: ")
	fmt.Scanln(&cfg.FBToken)

	err = WriteJsonFile("eulogist_config.json", cfg)
	if err != nil {
		return nil, fmt.Errorf("GenerateEulogistConfig: %v", err)
	}

	return &cfg, nil
}

//go:embed steve.png
var steveSkin []byte

// 根据赞颂者的配置 config，
// 在当前目录下生成用于启动 NEMC PC 的配置文件，
// 并返回该配置文件的绝对路径
func GenerateNetEaseConfig(config *EulogistConfig) (configPath string, err error) {
	cfg := DefaultNetEaseConfig()

	cfg.RoomInfo.IP = config.ServerIP
	cfg.RoomInfo.Port = config.ServerPort

	if !FileExist(config.SkinPath) {
		currentPath, _ := os.Getwd()
		cfg.SkinInfo.SkinPath = fmt.Sprintf(`%s\steve.png`, currentPath)
		err = os.WriteFile("steve.png", steveSkin, 0600)
		if err != nil {
			return "", fmt.Errorf("GenerateNetEaseConfig: %v", err)
		}
	} else {
		cfg.SkinInfo.SkinPath = config.SkinPath
	}

	err = WriteJsonFile("netease.cppconfig", cfg)
	if err != nil {
		return "", fmt.Errorf("GenerateNetEaseConfig: %v", err)
	}

	configPath, _ = os.Getwd()
	configPath = fmt.Sprintf(`%s\netease.cppconfig`, configPath)

	return
}
