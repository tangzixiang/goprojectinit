package main

import (
	"fmt"
	"github.com/tangzixiang/goprojectinit/internal/request"
	"os"
	"path/filepath"

	"github.com/tangzixiang/goprojectinit/internal/config"
	"github.com/tangzixiang/goprojectinit/internal/dir"
	"github.com/tangzixiang/goprojectinit/internal/options"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

func main() {

	opts := options.Parse()
	if opts == nil {
		return
	}

	// 获取项目地址的绝对路径
	projectPath := dir.GetProjectPath(opts)

	Log(fmt.Sprintf("project path is %q", projectPath))

	// 注册异常退出时需要将执行到一半的项目删除
	RegisterExitDeferFunc(func() {
		exists, err := PathExists(projectPath)
		if err != nil || !exists {
			return
		}

		_ = os.RemoveAll(projectPath)
	})

	// 创建项目目录
	if !dir.MakeProjectPathDir(opts, projectPath) {
		return
	}

	var err error
	var configPath string

	if opts.ConfigPath != nil {
		configPath, err = PathAbs(*opts.ConfigPath)
		DealErr(err, true)
	} else {
		configPath = request.DownloadFiles() // 下载配置文件到临时目录
	}

	config.ConfigPath = configPath
	config.ConfigPathDir = filepath.Dir(configPath)

	config.Init(opts)

	// 切换工作目录
	DealErr(os.Chdir(projectPath), true)

	DealErr(os.Mkdir("configs", dir.DirMode), true)
	CopyFileTo("configs", config.ConfigPath)
	CopyFileTo("configs", config.ConfigContentDirPath)

	//  子目录
	dir.MakeProjectSubDir(config.Dirs)
	Log(fmt.Sprintf("projoct %v init success~", opts.Args.ProjectName))
}
