# Go 项目工程初始化工具

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tangzixiang/goprojectinit) [![](https://badgen.net/github/branches/tangzixiang/goprojectinit)](https://github.com/tangzixiang/goprojectinit/branches) [![](https://badgen.net/github/stars/tangzixiang/goprojectinit)](https://github.com/tangzixiang/goprojectinit/stargazers) [![](https://badgen.net/github/commits/tangzixiang/goprojectinit)](https://github.com/tangzixiang/goprojectinit/commits/master)[![](https://img.shields.io/badge/release-v1.2.0-brightgreen)](https://github.com/tangzixiang/goprojectinit/releases)


`go get github.com/tangzixiang/goprojectinit`



## 摘要

- 概述
- 使用帮助
- 使用说明
  - 快速创建项目
  - 工具型项目
  - 类库型项目
  - 多服务项目
  - 自定义目录
  - vendor
  - 展示更多调试信息
  - 指定项目位置
  - 覆盖项目路径
- 注意
- 参考



## 概述

**goprojectinit** 是一个可以快速初始化 go 项目环境的工具。



## 使用帮助

```bash
$goprojectinit -h
Usage:
  goprojectinit [OPTIONS] projectname...

Application Options:
  -v, --version      show this tool version
  -b, --verbose      Show verbose debug information
  -c, --cover        if the project path exists ,cover the directory and init the project
  -t, --tool         tool mean this project is a tool project,so the main-file will be placed in project root directory
  -e, --empty        empty mean this project is a empty project or lib project
  -l, --lib          same as --empty
  -n, --usevendor    usevendor mean this project init whit vendor,default use go-modules
  -k, --nokeep       don't add .keep each dir
  -d, --dirs=        mkdir customer set,example: api,internal/configs,internal/model
  -m, --modulename=  modulename use for go modules init file: go.mod,default use project name
  -p, --targetpath=  project should init in the which directory,default is current path,if target directory not exists will be created

Help Options:
  -h, --help         Show this help message

Arguments:
  projectname:       init the project with this name, the first name will be named for project,then all remaining names will be sub service name in cmd directory
```



## 使用说明



### 快速创建项目

最简单的命令如下：

```bash
$ goprojectinit myproject

$ tree myproject/
myproject/
├── README.md
├── api
├── assets
├── build
│   ├── ci
│   └── package
├── cmd
│   └── myproject
│       └── myproject.go
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── go.mod
├── init
├── internal
├── pkg
├── scripts
├── test
├── third_party
├── tools
├── web
└── website

21 directories, 3 files
```

通过这么一个简单的命令，快速的初始化好一个 go 项目的目录环境，这里使用的是默认的目录配置。

上面的例子中项目 `myproject` 的项目入口文件为 `myproject/cmd/myproject/myproject.go`。



### 工具型项目

若新建项目为工具型项目则指定 `-t` 或者 `--tool` 参数即可:

```bash
$ goprojectinit -t mytool
$ tree mytool/
mytool/
├── README.md
├── assets
├── build
│   ├── ci
│   └── package
├── configs
├── docs
├── examples
├── githooks
├── go.mod
├── init
├── internal
├── main.go
├── pkg
├── scripts
└── tes
```

是否工具型项目的区别在于项目根目录是否存在 `main.go` 文件。



### 类库型项目

若新建项目为工具型项目则指定 `-l` 或者 `--lib` 参数即可:

```bash
$ goprojectinit -l mylib
$ tree mylib/
mylib/
├── README.md
├── assets
├── docs
├── examples
├── githooks
├── go.mod
├── scripts
└── test

6 directories, 2 files
```

类库型项目不会主动创建任何 go 文件。



### 多服务项目

**goprojectinit** 支持指定新建项目为多服务项目,只需提供多个 `projectname` 即可,首个 `projectname` 将作为项目名。

```bash
$ goprojectinit myproject subserver1 subserver2
$ tree myproject/
myproject/
├── README.md
├── api
├── assets
├── build
│   ├── ci
│   └── package
├── cmd
│   ├── myproject
│   │   └── myproject.go
│   ├── subserver1
│   │   └── subserver1.go
│   └── subserver2
│       └── subserver2.go
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── go.mod
├── init
├── internal
├── pkg
├── scripts
├── test
├── third_party
├── tools
├── web
└── website

23 directories, 5 files
```



### 自定义目录

若不想使用默认提供的几种目录初始化模板，**goprojectinit** 支持自定义需要初始化的目录，通过 `-d` 或则 `--dirs` 指定。

```bash
$ goprojectinit -d api,service,model myproject
$ tree myproject/
myproject/
├── README.md
├── api
├── cmd
│   └── myproject
│       └── myproject.go
├── go.mod
├── model
└── service

5 directories, 3 file
```



### vendor

**goprojectinit** 支持使用 [govendor](github.com/kardianos/govendor) 作为模块处理，在初始化项目的时候带上 `-n` 或者 `--usevendor` 即可 。

```bash
$ goprojectinit -n myproject
$ tree myproject/
myproject/
├── README.md
├── api
├── assets
├── build
│   ├── ci
│   └── package
├── cmd
│   └── myproject
│       └── myproject.go
├── configs
├── deployments
├── docs
├── examples
├── githooks
├── init
├── internal
├── pkg
├── scripts
├── test
├── third_party
├── tools
├── vendor
│   └── vendor.json
├── web
└── website

22 directories, 3 files
```

注意与 `go-modules` 模式不同的是，使用 `vendor` 模式可能需要新建项目位于 `$GOPATH` 下



### 展示更多调试信息

如果想要能够清楚的看到项目初始化过程的详细信息,只需要附带 `-b` 参数:

```bash
$ goprojectinit myproject -b

[goprojectinit] project path is "~/myproject"
[goprojectinit] make new directory ~/myproject success~
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~
```



### 指定项目位置

如果不想在当前目录下初始化项目只需要通过 `-p` 或者 `--targetpath` 指定目标地址即可:

```bash
tangzixiang$ sudo goprojectinit -p /var/www myproject
```



### 覆盖项目路径

若新建的项目路径下存在同名目录，默认不会进行覆盖，会维持目录原样并退出初始化过程，如果需要进行覆盖则可以通过 `-c` 或者 `--cover` 完成:

```bash
$ goprojectinit -p /var/www -c myproject
```

上面这个示例为，在 `/var/www` 目录下创建默认 `go` 项目 `myproject` 若已存在则删除并重新创建。



##  注意

1. 当项目初始化过程中因为其他外界因素导致中途失败后，工具会自动清除创建到一半的项目。
2. 项目会默认带 `.goprohectinit` 目录用于存放初始化相关内容
3. 初始化默认使用的是 `go-modules` 模式，需要 go 版本的支持，go 在 `go1.11` 版本开始支持




## 参考：

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
