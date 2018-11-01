package options

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/tangzixiang/goprojectinit/pkg/utils"
)

type HelpOptions struct {
	Verbose    bool    `short:"v" long:"verbose" description:"Show verbose debug information"`
	TargetPath *string `short:"p" long:"targetpath" description:"Project should init in the which directory,default is current path"`
	Cover      bool    `short:"c" long:"cover" description:"if the project path exists ,cover the directory and init the project"`
	ConfigPath *string `short:"f" long:"configfile" description:"which init-config file should be use,default is project-path/configs/config.yaml will be download"`
	Args       struct {
		ProjectName string `positional-arg-name:"projectname" description:"init the project with this name" `
	} `positional-args:"yes" required:"yes"`
}

var (
	opts       HelpOptions
	configPath = `configs/project-init.yaml`
)

// SetConfigPath 设置配置文件地址
func SetConfigPath(path string) {
	configPath = path
}

// GetConfigPath 获取配置文件地址
func GetConfigPath() string {
	return configPath
}

// Parse 开始解析命令行参数
func Parse() *HelpOptions {
	parse := parseOptions()

	if parse != nil {
		utils.Verbose = opts.Verbose
	}

	return &opts
}

func parseOptions() *flags.Parser {
	parse := flags.NewParser(&opts, flags.Default)
	_, err := parse.Parse()
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
