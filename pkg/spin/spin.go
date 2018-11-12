package spin

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

// 是否加载中
var loading = false

// Start 开始一个新的进度,如果正在加载中返回的参数为 nil,是否加载中应使用 Loading 方法获取
func Start(msg string) StopFunc {
	if loading {
		return nil
	}

	loading = true
	spin := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	spin.Prefix = fmt.Sprintf("[goprojectinit] %v ", msg)
	spin.Color("magenta")
	spin.Start()

	return func() {
		spin.Stop()
		loading = false
	}
}

// Loading 是否加载中
func Loading() bool {
	return loading
}

// StopFunc 进度回调
type StopFunc func()
