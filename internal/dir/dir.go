package dir

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tangzixiang/goprojectinit/internal/options"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// 权限信息
var (
	DirMode = os.ModeDir | os.ModePerm
)

// GetProjectPath 获取项目路径，返回的路径为绝对路径
func GetProjectPath(opts *options.HelpOptions) string {
	var projectPath string
	var err error

	if opts.TargetPathDir == nil {
		projectPath, err = os.Getwd()
	} else {
		projectPath, err = filepath.Abs(*opts.TargetPathDir)
		if err != nil && !os.IsNotExist(err) {
			DealErr(err, false)
		}
	}

	DealErr(err, false)

	return filepath.Join(projectPath, opts.Args.ProjectName)
}

// MakeProjectPathDir 创建项目目录,返回是否创建成功
func MakeProjectPathDir(opts *options.HelpOptions, projectPath string) bool {
	var err error

	exist, err := PathExists(projectPath)
	DealErr(err, false)

	if exist {
		if !opts.Cover { // 已存在该目录但不需要覆盖
			DealErr(errors.New(fmt.Sprintf("file or directory %v was exists", projectPath)), false)
		}

		fmt.Printf("[goprojectinit] are you sure to cover %v directory,type yes or no~\n", projectPath)
		if !ensureCover() {
			os.Exit(1)
			return false
		}

		// 确认需要重新创建目录
		DealErr(os.RemoveAll(projectPath), false)
		Log(fmt.Sprintf("directory %v remove success~", projectPath))
	}

	DealErr(os.MkdirAll(projectPath, DirMode), true)
	Log(fmt.Sprintf("make new directory %v success~", projectPath))

	return true
}

func ensureCover() bool {
	text := strings.ToLower(getScannerText())
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
	DealErr(scanner.Err(), false)

	return text
}

// MakeProjectSubDir 创建项目子目录
func MakeProjectSubDir(dirs []string){
	for _, dir := range dirs {
		DealErr(os.MkdirAll(dir,DirMode),true)
	}
}