package config

import (
	"github.com/tangzixiang/goprojectinit/internal/options"
	"github.com/tangzixiang/goprojectinit/pkg/utils"
)

// Config 配置
type Config struct {
	DirPath string `yaml:"dir"`
	ConfigPath string
}

// 初始化目录
var DirM = make(map[string]string)


func Init(){

	configPath := options.GetConfigPath()

	utils.Log(configPath)
}