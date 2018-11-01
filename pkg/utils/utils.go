package utils

import (
	"fmt"
	"os"
)

var Verbose  = false

// DealErr 处理异常并退出程序
func DealErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[goprojectinit] init failed :%v\n", err)
		os.Exit(1)
	}
}

// Log 输出日志
func Log(msg string) {
	if Verbose {
		fmt.Printf(fmt.Sprintf("[goprojectinit] %v\n", msg))
	}
}