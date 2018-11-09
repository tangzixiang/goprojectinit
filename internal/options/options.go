package options

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

type HelpOptions struct {
	Verbose       bool    `short:"v" long:"verbose" description:"Show verbose debug information"`
	TargetPathDir *string `short:"p" long:"targetpathdir" description:"Project should init in the which directory,default is current path,if target directory not exists will be created"`
	Cover         bool    `short:"c" long:"cover" description:"if the project path exists ,cover the directory and init the project"`
	ConfigPath    *string `short:"f" long:"configfile" description:"which init-config file should be use,if not set, default file will be download"`
	Args          struct {
		ProjectName string `positional-arg-name:"projectname" description:"init the project with this name" `
	} `positional-args:"yes" required:"yes"`
}

var (
	opts HelpOptions
)

// Parse 开始解析命令行参数
func Parse() *HelpOptions {
	parse := parseOptions()

	if parse == nil {
		return nil
	}

	Verbose = opts.Verbose
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
