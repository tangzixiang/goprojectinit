package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

type config struct {
	DirPath          *string `yaml:"dir" json:"dir"`
	MainFileTempPath *string `yaml:"mainFileTemp" json:"mainFileTemp"`
}

// DirM 初始化目录
var Dirs []string

// MainFileTemplatePath 模板文件位置
var MainFileTemplatePath string

// Config 配置项
var Config = new(config)

// 配置文件目录，包绝对路径
var PathDir string

// ParseConfigContentDirs 解析 dir 字段
func ParseConfigContentDirs(configContentDirPath string) {

	exists, err := PathExists(configContentDirPath)
	DealErr(err, true)

	if !exists {
		DealErr(
			fmt.Errorf("dir config file not find in %q", configContentDirPath), true)
	}

	parseFile(configContentDirPath, &Dirs)

	Log("read dir config file success~")
	Log(fmt.Sprintf("dir's : %q", Dirs))

	return
}

// ParseConfigFile 解析主配置文件
func ParseConfigFile(configPath string) {

	// 获取得到是绝对路径
	exists, err := PathExists(configPath)
	DealErr(err, true)

	// 找不到真实的文件
	if !exists {
		DealErr(
			fmt.Errorf("file not find in %q", configPath), true)
	}

	parseFile(configPath, Config)

	Log("read init config file success~")
}

// 解析配置文件内容
func parseFile(path string, out interface{}) {
	fileBytes, err := ioutil.ReadFile(path)
	DealErr(err, true)

	switch filepath.Ext(path) {
	case ".yaml":
		DealErr(yaml.Unmarshal(fileBytes, out), true)
	case ".json":
		DealErr(json.Unmarshal(fileBytes, out), true)
	}
}
