package request

import (
	"fmt"
	"github.com/tangzixiang/goprojectinit/pkg/task"
	"io"
	"net/http"
	"os"
	"path/filepath"

	. "github.com/tangzixiang/goprojectinit/internal/global"
	"github.com/tangzixiang/goprojectinit/pkg/spin"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// 文件下载地址
var (
	DownloadedFile = make(map[string]bool)
)

// DownloadFiles 下载所有远程配置文件
func DownloadAllFiles() error{
	var spinStop func()

	if !spin.Loading() {
		spinStop = spin.Start("downloading...")
	}

	if err := DownloadFile(FileNameConfig);err!= nil {
		return err
	}
	if err := DownloadFile(FileNameDir);err!= nil {
		return err
	}
	if err := DownloadFile(FileNameMainFileTemplate);err!= nil {
		return err
	}

	if spin.Loading() {
		spinStop()
	}

	 Log("config file download success~")
	return nil
}

func downloadAndWriteFile(url string) error {
	var err error
	var resp *http.Response

	// 已经下载过了
	if DownloadedFile[url] {
		return nil
	}

	// 临时文件
	fileName := filepath.Join(TempDir, filepath.Base(url))
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// 下载
	resp, err = http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			Log(fmt.Sprintf("request remote file %v failed: %v", url, err))
		}
	}()

	// 下载内容写入临时文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	DownloadedFile[url] = true
	return nil
}

// DownloadFile 下载指定远程配置文件
func DownloadFile(name string) error{
	return task.Tasks.Call(name).Err()
}

func Init() {
	DealErr(
		task.Tasks.
			HandleFuc(FileNameConfig, func() error {
				return downloadAndWriteFile(RemoteURLConfig)
			}).
			HandleFuc(FileNameDir, func() error {
				return downloadAndWriteFile(RemoteURLDir)
			}).
			HandleFuc(FileNameMainFileTemplate, func() error {
				return downloadAndWriteFile(RemoteURLMainFileTemplate)
			}).Err(),
		true,
	)
}
