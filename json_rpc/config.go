package json_rpc

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Url     string `yaml:"url"`
	UserId  string `yaml:"userId"`
	Address string `yaml:"address"`
	Target  string `yaml:"target"`
	Token   string `yaml:"token"`
}

func GetConfig() Config {
	// 1. 读取配置文件内容
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		data, err = os.ReadFile("../config.yaml")
		if err != nil {
			panic(fmt.Sprintf("读取配置文件失败：%v", err))
		}
	}

	// 2. 解析yaml到结构体
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("解析配置失败：%v", err))
	}
	return cfg
}
