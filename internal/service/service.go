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

	configsDirPath:= filepath.Join(projectPath, "configs")
	DealErr(os.Mkdir(configsDirPath, DirMode), true)

	// 初始化文件下载环境
	request.Init()

	// 检查配置文件
	DealErr(checkConfigPath(opts.ConfigPath,configsDirPath), true)

	// 检查缺失的文件并下载
	DealErr(checkConfigContent(configsDirPath), true)

	// 切换工作目录
	DealErr(os.Chdir(projectPath), true)

	//  子目录
	dir.MakeProjectSubDir(config.Dirs)

	// 创建项目入口文件
	file.WriteMainFileWithTemp(opts.Args.ProjectName, opts.IsTool)

	Log(fmt.Sprintf("projoct %v init success~", opts.Args.ProjectName[0]))
}

func checkConfigPath(path *string,configsDirPath string) error {
	var configPath, configPathDir string
	var err error

	if path == nil {
		DealErr(request.DownloadAllFiles(), true) // 下载所有配置文件到临时目录

		configPath, configPathDir = filepath.Join(TempDir, FileNameConfig), TempDir
	} else {
		configPath, err = PathAbs(*path)
		if err != nil {
			return err
		}
		configPathDir = filepath.Dir(configPath)
	}

	config.PathDir = configPathDir
	config.ParseConfigFile(configPath)

	return CopyFileTo(configsDirPath, configPath)
}

func checkConfigContent(configsDirPath string) error {
	var shouldDownloadFile []string
	var dirPath, templatePath string
	var err error

	// dir 文件
	if config.Config.DirPath == nil {
		if !request.DownloadedFile[RemoteURLDir] {
			Log("dir config file not found")
			shouldDownloadFile = append(shouldDownloadFile, FileNameDir)
		}

		dirPath = filepath.Join(TempDir, FileNameDir)
	} else {
		dirPath = *config.Config.DirPath

		if !filepath.IsAbs(dirPath) {
			dirPath = filepath.Join(config.PathDir, dirPath)
		}

		exists, err := PathExists(dirPath)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("file not find in %q", dirPath)
		}
	}

	// main 目标文件
	if config.Config.MainFileTempPath == nil {

		if !request.DownloadedFile[RemoteURLMainFileTemplate] {
			Log("main file template not found")
			shouldDownloadFile = append(shouldDownloadFile, FileNameMainFileTemplate)
		}

		templatePath = filepath.Join(TempDir, FileNameMainFileTemplate)
	} else {
		templatePath = *config.Config.MainFileTempPath

		if !filepath.IsAbs(templatePath) {
			templatePath = filepath.Join(config.PathDir, templatePath)
		}

		exists, err := PathExists(templatePath)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("file not find in %q", templatePath)
		}
	}

	var spinStop func()
	if !spin.Loading() {
		spinStop = spin.Start("downloading...")
	}

	for _, fileName := range shouldDownloadFile {
		if err = request.DownloadFile(fileName); err != nil {
			return err
		}
	}

	if spin.Loading() {
		spinStop()
	}

	if len(shouldDownloadFile) > 0 {
		Log(fmt.Sprintf("file %q download success~", shouldDownloadFile))
	}

	config.ParseConfigContentDirs(dirPath)
	err = CopyFileTo(configsDirPath, dirPath)
	if err != nil {
		return err
	}

	config.MainFileTemplatePath = templatePath
	err = CopyFileTo(configsDirPath, templatePath)
	if err != nil {
		return err
	}

	return nil
}
