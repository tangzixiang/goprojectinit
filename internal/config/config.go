package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/tangzixiang/goprojectinit/internal/options"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

type config struct {
	DirPath *string `yaml:"dir" json:"dir"`
}

// DirM 初始化目录
var Dirs []string
// Config 配置项
var Config = new(config)

// 配置文件目录，绝对路径
var (
	ConfigPathDir string
	ConfigPath    string
	ConfigContentDirPath string
)

func Init(opts *options.HelpOptions) {
	parseConfig(opts)
	parseConfigContentDirs(opts)
}

func parseConfigContentDirs(opts *options.HelpOptions) {

	if Config.DirPath == nil {
		Log("dir config file not found")
		return
	}

	ConfigContentDirPath = *Config.DirPath
	if !filepath.IsAbs(ConfigContentDirPath) {
		ConfigContentDirPath = filepath.Join(ConfigPathDir,ConfigContentDirPath)
	}

	exists, err := PathExists(ConfigContentDirPath)
	DealErr(err, true)

	if !exists {
		DealErr(
			fmt.Errorf("dir config file not find in %q", ConfigContentDirPath), true)
	}

	parseFile(ConfigContentDirPath, &Dirs)

	Log("read dir config file success~")
	Log(fmt.Sprintf("dir's : %q", Dirs))
}

func parseConfig(opts *options.HelpOptions) {

	// 获取得到是绝对路径
	exists, err := PathExists(ConfigPath)
	DealErr(err, true)

	// 找不到真实的文件
	if !exists {
		DealErr(
			fmt.Errorf("file not find in %q", ConfigPath), true)
	}

	parseFile(ConfigPath, Config)

	Log("read init config file success~")
}

// 解析配置文件
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
