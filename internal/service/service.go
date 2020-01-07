package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tangzixiang/goprojectinit/internal/config"
	"github.com/tangzixiang/goprojectinit/internal/dir"
	"github.com/tangzixiang/goprojectinit/internal/file"
	. "github.com/tangzixiang/goprojectinit/internal/global"
	"github.com/tangzixiang/goprojectinit/internal/options"
	"github.com/tangzixiang/goprojectinit/internal/request"
	"github.com/tangzixiang/goprojectinit/pkg/spin"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

func Run() {

	opts := options.Parse()
	if opts == nil {
		return
	}

	// 获取项目地址的绝对路径
	projectName := opts.Args.ProjectName[0]
	projectPath := dir.GetProjectPath(opts.TargetPath, projectName)

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

	configsDirPath := filepath.Join(projectPath, ".goprojectinit")
	DealErr(os.Mkdir(configsDirPath, DirMode), true)

	// 项目类型
	projType := projectType(opts)

	// 解析需要初始化的目录
	config.ParseDirs(projType, opts.Dirs)

	// 初始化文件下载环境
	request.Init()

	// 检查配置文件
	DealErr(request.DownloadAllFiles(), true) // 下载所有配置文件到临时目录

	// 检查缺失的文件并下载
	DealErr(checkConfigContent(configsDirPath, projType), true)

	// 切换工作目录
	DealErr(os.Chdir(projectPath), true)

	// 检查 git 环境并初始化
	DealErr(checkGitEnv(), true)

	// 新建 README.md
	DealErr(touchREADME(), true)

	// 初始化 vendor 环境 或则 go modules 环境
	DealErr(checkEnv(opts.UseVendor, opts.ModulesName, projectName), true)

	//  子目录
	dir.MakeProjectSubDir(config.Dirs, opts.NoKeep)

	// 创建项目入口文件
	file.WriteMainFileWithTemp(opts.Args.ProjectName, projType, config.MainFileTemplatePath)

	Log(fmt.Sprintf("projoct %v init success~", opts.Args.ProjectName[0]))
}

func projectType(ops *options.HelpOptions) string {
	switch {
	case ops.Lib:
		fallthrough
	case ops.Empty:
		return "empty"
	case ops.Tool:
		return "tool"
	default:
		return "default"
	}
}

func checkEnv(useVendor bool, moduleName, projectName string) error {
	if useVendor {
		return checkVendor()
	} else {
		if moduleName == "" {
			return checkGOModules(projectName)
		}
		return checkGOModules(moduleName)
	}
}

func checkVendor() error {
	var err error

	if _, err = exec.LookPath("govendor"); err != nil {

		var spinStop func()
		if !spin.Loading() {
			spinStop = spin.Start("downloading...")
		}

		// 需要下载 govendor
		err = ExecCommand("go", "get", "-u", "github.com/kardianos/govendor")
		if spin.Loading() {
			spinStop()
		}

		if err != nil {
			return err
		}
	}

	if err = ExecCommand("govendor", "init"); err != nil {
		return err
	}

	return nil
}

func checkGOModules(moduleName string) error {

	if err := ExecCommand("go", "mod", "init", moduleName); err != nil {
		Log("my be you go version not support go-modules,if want use vendor please run command with -n")
		return err
	}

	if err := ExecCommand("go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}

func touchREADME() error {
	return ExecCommand("touch", "README.md")
}

func checkGitEnv() error {
	var err error

	if err = request.DownloadFile(FileNameGitIgnore); err != nil {
		return err
	}

	if _, err = exec.LookPath("git"); err != nil {
		return err
	}

	if err = ExecCommand("git", "init"); err != nil {
		return err
	}

	return CopyFileTo(".", filepath.Join(TempDir, FileNameGitIgnore))
}

func checkConfigContent(configsDirPath string, projType string) error {
	var shouldDownloadFile []string
	var templatePath string
	var err error

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
			templatePath = filepath.Join(TempDir, templatePath)
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
			break
		}
	}

	if spin.Loading() {
		spinStop()
	}

	if err != nil {
		return err
	}

	if len(shouldDownloadFile) > 0 {
		Log(fmt.Sprintf("file %q download success~", shouldDownloadFile))
	}

	config.MainFileTemplatePath = templatePath
	err = CopyFileTo(configsDirPath, templatePath)
	if err != nil {
		return err
	}

	err = CopyFileTo(configsDirPath, filepath.Join(TempDir, FileNameGitIgnore))
	if err != nil {
		return err
	}

	return nil
}
