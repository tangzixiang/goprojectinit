package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	. "github.com/tangzixiang/goprojectinit/internal/global"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

type config struct {
	MainFileTempPath *string `yaml:"mainFileTemp" json:"mainFileTemp"`
}

// DirM 初始化目录
var Dirs []string

// MainFileTemplatePath 模板文件位置
var MainFileTemplatePath string

// Config 配置项
var Config = new(config)

// ParseDirs 解析需要的目录
func ParseDirs(projType string, dirs string) {

	switch projType {
	case "empty":
		Dirs = EmptyPath
	case "tool":
		Dirs = ToolPath
	}

	if dirs != "" {
		dirs = strings.TrimSpace(dirs)
		splits := strings.Split(dirs, ",")
		if len(splits) > 0 {
			for _, s := range splits {
				Dirs = append(Dirs, strings.TrimSpace(s))
			}
		}
	}

	if len(Dirs) == 0 {
		Dirs = DefaultPath
	}

	Log(fmt.Sprintf("dir's : %q", Dirs))

	return
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
