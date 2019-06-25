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
	DirPath          *string `yaml:"dir"          json:"dir"`
	MainFileTempPath *string `yaml:"mainFileTemp" json:"mainFileTemp"`
}

type pdir struct {
	Empty   []string `json:"empty"      yaml:"empty"`
	Tool    []string `json:"tool"       yaml:"tool"`
	Default []string `json:"default"    yaml:"default"`
}

// DirM 初始化目录
var Dirs []string

// PDirs 默认不同项目类型的目录结构
var PDirs pdir

// MainFileTemplatePath 模板文件位置
var MainFileTemplatePath string

// Config 配置项
var Config = new(config)

// 配置文件目录，包绝对路径
var PathDir string

// ParseConfigContentDirs 解析 dir 字段
func ParseConfigContentDirs(configContentDirPath string, projType string) {

	exists, err := PathExists(configContentDirPath)
	DealErr(err, true)

	if !exists {
		DealErr(
			fmt.Errorf("dir config file not find in %q", configContentDirPath), true)
	}

	err = parseFile(configContentDirPath, &Dirs) // 兼容旧版本，尝试解析为单独的数组
	if err != nil {
		err = parseFile(configContentDirPath, &PDirs)
	}
	DealErr(err, true)

	if len(Dirs) == 0 {
		switch projType {
		case "empty":
			Dirs = PDirs.Empty
		case "tool":
			Dirs = PDirs.Tool
		default:
			Dirs = PDirs.Default
		}
	}

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
		DealErr(fmt.Errorf("file not find in %q", configPath), true)
	}

	DealErr(parseFile(configPath, Config), true)

	Log("read init config file success~")
}

// 解析配置文件内容
func parseFile(path string, out interface{}) (err error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	switch filepath.Ext(path) {
	case ".yaml":
		err = yaml.Unmarshal(fileBytes, out)
	case ".json":
		err = json.Unmarshal(fileBytes, out)
	}

	return
}
