package spin

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

// StopFunc 进度回调
type StopFunc func()

// Start 开始一个新的进度
func Start(msg string) StopFunc {
	spin := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	spin.Prefix = fmt.Sprintf("[goprojectinit] %v ",msg)
	spin.Color("magenta")
	spin.Start()

	return func() {
		spin.Stop()
	}
}
