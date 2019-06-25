package file

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tangzixiang/goprojectinit/internal/global"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

// WriteMainTempFile 实例化模板到文件
func WriteMainFileWithTemp(fileNames []string, projType string, mainFileTemplatePath string) {

	switch projType {
	case "empty": // do not thing
	case "tool":
		fileNames = fileNames[1:]

		// 将 main 文件写在根目录下
		writeToolMainFile(mainFileTemplatePath)
		fallthrough
	default:
		for _, fileName := range fileNames {
			writeCMDMainFile(fileName, mainFileTemplatePath)
		}
	}

}

func writeCMDMainFile(fileName string, mainFileTemplatePath string) {

	cmdPath := filepath.Join("cmd", fileName)
	DealErr(os.MkdirAll(cmdPath, global.DirMode), true)

	mainFile, err := os.Create(fmt.Sprintf("%v.go", filepath.Join(cmdPath, fileName)))
	DealErr(err, true)

	write(mainFile, mainFileTemplatePath)
}

func writeToolMainFile(mainFileTemplatePath string) {

	mainFile, err := os.Create("main.go")
	DealErr(err, true)

	write(mainFile, mainFileTemplatePath)
}

func write(file *os.File, mainFileTemplatePath string) {

	temp, err := template.ParseFiles(mainFileTemplatePath)
	DealErr(err, true)

	defer (func() {
		if err := file.Close(); err != nil {
			Log(fmt.Sprintf("close file %v failed: %v~", file.Name(), err))
		}
	})()

	DealErr(temp.Execute(file, nil), true)
}
