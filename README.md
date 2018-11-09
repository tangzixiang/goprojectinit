# Go 项目工程初始化工具



`go get github.com/tangzixiang/goprojectinit`



## 概述

goprojectinit 是一个可以快速初始化 go 项目环境的工具，该工具支持以下特点：

- 支持自定义需要初始化的目录环境
- 支持灵活指定项目进行初始化的路径



## 使用帮助

```bash
$ goprojectinit -h
Usage:
  goprojectinit [OPTIONS] projectname

Application Options:
  -v, --verbose        Show verbose debug information
  -p, --targetpathdir= Project should init in the which directory,default is current path,if target directory not exists will be created
  -c, --cover          if the project path exists ,cover the directory and init the project
  -f, --configfile=    which init-config file should be use,default is project-path/configs/config.yaml will be download

Help Options:
  -h, --help           Show this help message

Arguments:
  projectname:         init the project with this name
```



## 快速入门

最简单的命令如下：

```bash
$ goprojectinit myproject
$ tree myproject
```





参考：

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)