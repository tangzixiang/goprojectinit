package global

import (
	"os"
)

// 常量
const (
	DirMode = os.ModeDir | os.ModePerm // 目录权限信息

	RemoteURLMainFileTemplate = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/main-file.temp`
	RemoteURLGitIgnore        = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/.gitignore`

	FileNameMainFileTemplate = `main-file.temp`
	FileNameGitIgnore        = `.gitignore`
)

// TempDir 临时目录
var TempDir = os.TempDir()

// FileNameUrlM 需要下载的文件
var FileNameUrlM = map[string]string{
	FileNameMainFileTemplate: RemoteURLMainFileTemplate,
	FileNameGitIgnore:        RemoteURLGitIgnore,
}

// 初始化目录
var (
	DefaultPath = []string{
		"api",
		"assets",
		"build/package",
		"build/ci",
		"cmd",
		"configs",
		"docs",
		"deployments",
		"examples",
		"githooks",
		"init",
		"internal",
		"pkg",
		"scripts",
		"test",
		"third_party",
		"tools",
		"web",
		"website",
	}

	ToolPath = []string{
		"build/package",
		"build/ci",
		"configs",
		"scripts",
		"internal",
		"pkg",
		"assets",
		"docs",
		"examples",
		"githooks",
		"init",
		"test",
		"docs",
	}

	EmptyPath = []string{
		"scripts",
		"assets",
		"docs",
		"examples",
		"githooks",
		"test",
	}
)
