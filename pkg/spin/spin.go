package spin

import (
	"time"

	"github.com/briandowns/spinner"
)

// StopFunc 进度回调
type StopFunc func()

// Start 开始一个新的进度
func Start() StopFunc {
	spin := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	spin.Prefix = "[goprojectinit] "
	spin.Suffix = "\n"
	spin.Color("magenta")
	spin.Start()

	return func() {
		spin.Stop()
	}
}
