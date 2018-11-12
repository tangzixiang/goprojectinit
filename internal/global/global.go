package global

import (
	"os"
)

const (
	DirMode = os.ModeDir | os.ModePerm // 目录权限信息

	RemoteURLConfig           = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/project-init.yaml`
	RemoteURLDir              = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/dir.json`
	RemoteURLMainFileTemplate = `https://raw.githubusercontent.com/tangzixiang/goprojectinit/master/configs/main-file.temp`

	FileNameConfig           = `project-init.yaml`
	FileNameDir              = `dir.json`
	FileNameMainFileTemplate = `main-file.temp`
)

var (
	TempDir = os.TempDir()
)
