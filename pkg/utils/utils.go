package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var Verbose = false

var deferFuncs []func()

// RegisterExitDeferFunc 注册异常退出时的提前处理器，提前处理器为非异步执行，程序将会在所有提前处理器按注册顺序执行完毕后退出
func RegisterExitDeferFunc(deferfunc func()) {
	deferFuncs = append(deferFuncs, deferfunc)
}

// DealErr 处理异常并直接退出程序
func DealErr(err error, doDeferFuncs bool) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[goprojectinit] failed :%v\n", err)

		if len(deferFuncs) > 0 && doDeferFuncs {
			for _, deferFunc := range deferFuncs {
				deferFunc()
			}
		}
		os.Exit(1)
	}
}

// Log 输出日志，当且仅当 Verbose 为 true 时输出
func Log(msg string) {
	if Verbose {
		fmt.Printf(fmt.Sprintf("[goprojectinit] %v\n", msg))
	}
}

// PathAbs 获取文件的绝对路径
func PathAbs(path string) (string, error) {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return "", err
		}
	}

	return path, nil
}

// PathExists 判断指定路径文件或文件夹是否存在,返回的错误为系统级调用错误，文件是否存在由第一个返回参数确认，应该优先判断异常再判断是否文件存在
func PathExists(path string) (bool, error) {
	var err error

	if _, err = os.Stat(path); err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// CopyFile 拷贝文件
func CopyFile(targetFilePath, srcFilePath string) error {
	return nil
}

// CopyFileTo 将文件拷贝打目录
func CopyFileTo(targetDir, filePath string) error {
	return exec.Command("cp", filePath, targetDir).Run()
}