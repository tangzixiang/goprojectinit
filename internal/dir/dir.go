package dir

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tangzixiang/goprojectinit/internal/options"
	"github.com/tangzixiang/goprojectinit/pkg/spin"
	"github.com/tangzixiang/goprojectinit/pkg/utils"
)

var (
	dirMode = os.ModeDir | os.ModePerm
)

// CopyFile 拷贝文件
func CopyFile(targetPath, srcPath string) error {
	return nil
}

// GetProjectPath 获取项目路径
func GetProjectPath(opts *options.HelpOptions) string {
	var projectPath string
	var err error

	if opts.TargetPath == nil {
		projectPath, err = os.Getwd()
	} else {
		projectPath, err = filepath.Abs(*opts.TargetPath)
	}

	utils.DealErr(err)

	return filepath.Join(projectPath, opts.Args.ProjectName)
}

// MakeProjectDir 创建项目目录
func MakeProjectDir(opts *options.HelpOptions, projectPath string) bool {
	var err error
	err = os.Mkdir(projectPath, dirMode)
	if err == nil {
		return true
	}

	// 非目录已存在错误 或则是已存在该目录但不覆盖
	if !os.IsExist(err) || !opts.Cover {
		utils.DealErr(err)
	}

	fmt.Printf("[goprojectinit] are you sure to cover %v directory,type yes or no~\n", projectPath)
	if !ensureCover() {
		os.Exit(1)
		return false
	}

	// 重新创建目录
	stop := spin.Start()
	if err := os.RemoveAll(projectPath); err != nil {
		utils.DealErr(err)
	}
	stop()

	utils.Log(fmt.Sprintf("directory %v remove success~", projectPath))

	err = os.Mkdir(projectPath, dirMode)
	if err != nil {
		utils.DealErr(err)
	}

	utils.Log(fmt.Sprintf("new directory %v success~", projectPath))

	return true
}

func ensureCover() bool {
	text := getScannerText()
	if text != "yes" && text != "y" {
		return false
	}
	return true
}

func getScannerText() string {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return ""
	}

	text := scanner.Text()
	err := scanner.Err()

	if err != nil {
		utils.DealErr(err)
	}

	return text
}
