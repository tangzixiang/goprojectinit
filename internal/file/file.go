package file

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tangzixiang/goprojectinit/internal/config"
	"github.com/tangzixiang/goprojectinit/internal/global"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// WriteMainTempFile 实例化模板到文件
func WriteMainFileWithTemp(fileNames []string, isToolProject bool) {

	if isToolProject {
		fileNames = fileNames[1:]

		// 将 main 文件写在根目录下
		writeToolMainFile()
	}

	for _, fileName := range fileNames {
		writeCMDMainFile(fileName)
	}
}

func writeCMDMainFile(fileName string) {

	DealErr(os.MkdirAll(filepath.Join("cmd", fileName), global.DirMode), true)

	mainFile, err := os.Create(fmt.Sprintf("%v.go", fileName))
	DealErr(err, true)

	write(mainFile)
}

func writeToolMainFile() {

	mainFile, err := os.Create("main.go")
	DealErr(err, true)

	write(mainFile)
}

func write(file *os.File) {

	temp, err := template.ParseFiles(config.MainFileTemplatePath)
	DealErr(err, true)

	defer (func() {
		if err := file.Close(); err != nil {
			Log(fmt.Sprintf("close file %v failed: %v~", file.Name(), err))
		}
	})()

	DealErr(temp.Execute(file, nil), true)
}
