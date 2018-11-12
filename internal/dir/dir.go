package dir

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tangzixiang/goprojectinit/internal/global"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// GetProjectPath 获取项目路径，返回的路径为绝对路径
func GetProjectPath(targetPath *string, projectName string) string {
	var err error

	projectPath, err := os.Getwd()
	DealErr(err, true)

	if targetPath != nil {

		absPath, err := PathAbs(*targetPath)
		DealErr(err, true)

		exists, err := PathExists(absPath)
		DealErr(err, true)

		if !exists {
			DealErr(errors.New(fmt.Sprintf("config file %v not found", *targetPath)), true)
		}
	}

	return filepath.Join(projectPath, projectName)
}

// MakeProjectPath 创建项目目录,返回是否创建成功
func MakeProjectPath(projectPath string, cover bool) bool {
	var err error

	exist, err := PathExists(projectPath)
	DealErr(err, false)

	if exist {
		if !cover { // 已存在该目录但不需要覆盖
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

	DealErr(os.MkdirAll(projectPath, global.DirMode), true)
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
func MakeProjectSubDir(dirs []string) {

	for _, dir := range dirs {
		DealErr(os.MkdirAll(dir, global.DirMode), true)
	}
}
