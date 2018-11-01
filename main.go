package main

import (
	"fmt"
	"os"

	"github.com/tangzixiang/goprojectinit/internal/config"
	"github.com/tangzixiang/goprojectinit/internal/dir"
	"github.com/tangzixiang/goprojectinit/internal/options"
	"github.com/tangzixiang/goprojectinit/pkg/utils"
)

func main() {

	opts := options.Parse()
	if opts == nil {
		return
	}

	projectPath := dir.GetProjectPath(opts)

	utils.Log(fmt.Sprintf("project path is %v", projectPath))

	if !dir.MakeProjectDir(opts, projectPath) {
		return
	}

	// 切换工作目录
	if err := os.Chdir(projectPath); err != nil {
		utils.DealErr(err)
	}

	var configPath string
	if opts.ConfigPath != nil {
		configPath = *opts.ConfigPath
	} else {
		// 下载配置文件到临时目录
		// 读取其中的配置项
		// 复制到 config_path 中
		//
	}

	options.SetConfigPath(configPath)
	config.Init()

	//  子目录

}
