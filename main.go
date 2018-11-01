package main

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/jessevdk/go-flags"
	"os"
	"path/filepath"
	"time"
)

type helpOptions struct {
	Verbose    bool    `short:"v" long:"verbose" description:"Show verbose debug information"`
	TargetPath *string `short:"p" long:"targetpath" description:"Project should init in the which directory,default is current path"`
	Cover      bool    `short:"c" long:"cover" description:"if the project path exists ,cover the directory and init the project"`
	Args       struct {
		ProjectName string `positional-arg-name:"projectname" description:"init the project with this name" `
	} `positional-args:"yes" required:"yes"`
}

var (
	opts helpOptions
)

func main() {

	parse := parseOptions()
	if parse == nil {
		return
	}

	projectPath := getProjectPath()

	log(fmt.Sprintf("project path is %v", projectPath))

	dirMode := os.ModeDir | os.ModePerm
	if !makeProjectDir(projectPath, dirMode) {
		return
	}

	// 切换工作目录
	if err := os.Chdir(projectPath); err != nil {
		dealErr(err)
	}

	// 下载配置文件

	// 初始化 Config

	//  子目录

}

func parseOptions() *flags.Parser {
	parse := flags.NewParser(&opts, flags.Default)
	_, err := parse.Parse();
	if err == nil {
		return parse
	}

	// 直接输入 -h 或则 --help
	if err.(*flags.Error).Type == flags.ErrHelp {
		return nil
	}

	fmt.Println()
	parse.WriteHelp(os.Stderr)
	return nil
}

func makeProjectDir(projectPath string, dirMode os.FileMode) bool {
	var err error
	err = os.Mkdir(projectPath, dirMode)
	if err == nil {
		return true
	}

	// 非目录已存在错误 或则是已存在该目录但不覆盖
	if !os.IsExist(err) || !opts.Cover {
		dealErr(err)
	}

	fmt.Printf("[goprojectinit] are you sure to cover %v directory,type yes or no~\n", projectPath)
	if !ensureCover() {
		os.Exit(1)
		return false
	}

	// 重新创建目录
	spin := newSpin()
	spin.Start()
	if err := os.RemoveAll(projectPath); err != nil {
		dealErr(err)
	}
	spin.Stop()
	log(fmt.Sprintf("directory %v remove success~", projectPath))

	err = os.Mkdir(projectPath, dirMode)
	if err != nil {
		dealErr(err)
	}
	log(fmt.Sprintf("new directory %v success~", projectPath))

	return true
}

func ensureCover() bool {
	text := getScannerText()
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
	err := scanner.Err()

	if err != nil {
		dealErr(err)
	}

	return text
}

func getProjectPath() string {
	var projectPath string
	var err error

	if opts.TargetPath == nil {
		projectPath, err = os.Getwd()
	} else {
		projectPath, err = filepath.Abs(*opts.TargetPath)
	}
	dealErr(err)

	return filepath.Join(projectPath, opts.Args.ProjectName)
}

func dealErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[goprojectinit] init failed :%v\n", err)
		os.Exit(1)
	}
}

func log(msg string) {
	if opts.Verbose {
		fmt.Printf(fmt.Sprintf("[goprojectinit] %v\n", msg))
	}
}

func newSpin() *spinner.Spinner {
	spin := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	spin.Prefix = "[goprojectinit] "
	spin.Suffix = "\n"
	spin.Color("magenta")
	return spin
}
