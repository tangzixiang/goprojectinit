package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tangzixiang/goprojectinit/pkg/spin"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// 文件下载地址
var (
	RemoteConfigFileURL    = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/project-init.yaml`
	RemoteConfigDirFileURL = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/dir.json`
)

// DownloadFiles 下载远程配置文件
func DownloadFiles() string {
	var err error

	spinStop := spin.Start("downloading...")
	defer spinStop()

	tempDir := os.TempDir()

	configFilePath := filepath.Join(tempDir, filepath.Base(RemoteConfigFileURL))
	configFile ,err := os.Create(configFilePath)
	DealErr(err,true)

	configDirFilePath := filepath.Join(tempDir, filepath.Base(RemoteConfigDirFileURL))
	configDirFile ,err := os.Create(configDirFilePath)
	DealErr(err,true)

	downloadAndWriteFile(RemoteConfigFileURL,configFile)
	downloadAndWriteFile(RemoteConfigDirFileURL,configDirFile)

	return configFilePath
}

func downloadAndWriteFile(url string,file *os.File) {
	var err error
	var resp *http.Response

	resp, err = http.Get(url)
	DealErr(err, true)

	defer func() {
		if err := resp.Body.Close(); err != nil {
			Log(fmt.Sprintf("request remote file %v failed: %v", url, err))
		}
	}()

	_, err = io.Copy(file, resp.Body)
	DealErr(err, true)
}
