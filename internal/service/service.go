package service

import (
	"fmt"
	"github.com/tangzixiang/goprojectinit/pkg/spin"
	"os"
	"path/filepath"

	"github.com/tangzixiang/goprojectinit/internal/config"
	"github.com/tangzixiang/goprojectinit/internal/dir"
	"github.com/tangzixiang/goprojectinit/internal/file"
	. "github.com/tangzixiang/goprojectinit/internal/global"
	"github.com/tangzixiang/goprojectinit/internal/options"
	"github.com/tangzixiang/goprojectinit/internal/request"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

func Run() {

	opts := options.Parse()
	if opts == nil {
		return
	}

	// 获取项目地址的绝对路径
	projectPath := dir.GetProjectPath(opts.TargetPath, opts.Args.ProjectName[0])

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
	if !dir.MakeProjectPath(projectPath, opts.Cover) {
		return
	}

	// 切换工作目录
	DealErr(os.Chdir(projectPath), true)
	DealErr(os.Mkdir("configs", DirMode), true)

	// 检查配置文件
	checkConfigPath(opts.ConfigPath)

	// 检查缺失的文件并下载
	checkConfigContent()

	//  子目录
	dir.MakeProjectSubDir(config.Dirs)

	// 创建项目入口文件
	file.WriteMainFileWithTemp(opts.Args.ProjectName, opts.IsTool)

	Log(fmt.Sprintf("projoct %v init success~", opts.Args.ProjectName[0]))
}

func checkConfigPath(path *string) {
	var configPath, configPathDir string
	var err error

	if path == nil {
		request.DownloadAllFiles() // 下载所有配置文件到临时目录

		configPath, configPathDir = filepath.Join(TempDir, FileNameConfig), TempDir
	} else {
		configPath, err = PathAbs(*path)
		DealErr(err, true)

		configPathDir = filepath.Dir(configPath)
	}

	config.PathDir = configPathDir
	config.ParseConfigFile(configPath)

	CopyFileTo("configs", configPath)

	return
}

func checkConfigContent() {
	var shouldDownloadFile []string
	var dirPath, templatePath string

	// dir 文件
	if config.Config.DirPath == nil {
		Log("dir config file not found")
		shouldDownloadFile = append(shouldDownloadFile, FileNameDir)

		dirPath = filepath.Join(TempDir, FileNameDir)
	} else {
		dirPath = *config.Config.DirPath

		if !filepath.IsAbs(dirPath) {
			dirPath = filepath.Join(config.PathDir, dirPath)
		}
	}
	config.ParseConfigContentDirs(dirPath)
	CopyFileTo("configs", dirPath)

	// main 目标文件
	if config.Config.MainFileTempPath == nil {
		Log("main file template not found")
		shouldDownloadFile = append(shouldDownloadFile, FileNameMainFileTemplate)

		templatePath = filepath.Join(TempDir, FileNameMainFileTemplate)
	} else {
		absPath, err := PathAbs(*config.Config.MainFileTempPath)
		DealErr(err, true)

		exists, err := PathExists(absPath)
		DealErr(err, true)

		if !exists {
			DealErr(
				fmt.Errorf("file not find in %q", absPath), true)
		}

		templatePath = absPath
	}
	config.MainFileTemplatePath = templatePath
	CopyFileTo("configs", templatePath)

	if !spin.Loading() {
		spinStop := spin.Start("downloading...")
		defer spinStop()
	}

	for _, fileName := range shouldDownloadFile {
		request.DownloadFile(fileName)
		Log(fmt.Sprintf("file %v download success~", fileName))
	}
}
