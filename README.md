# Go 项目工程初始化工具



`go get github.com/tangzixiang/goprojectinit`



## 概述

**goprojectinit** 是一个可以快速初始化 go 项目环境的工具，该工具具体以下特色：

- 支持自定义初始化的目录环境
- 主动创建项目单个或多个入口文件
- 支持指定项目的初始化路径
- 自动初始化 `git` 环境
- 自动初始化 `vendor` 环境或者自动初始化 `modules`环境 (默认)
- 自动创建 `README.md` 文件



## 使用帮助

```bash
$ goprojectinit -h
Usage:
  goprojectinit [OPTIONS] projectname...

Application Options:
  -v, --version      show this tool version
  -b, --verbose      Show verbose debug information
  -c, --cover        if the project path exists ,cover the directory and init the project
  -t, --istool       istool mean this project is a tool project,so the main-file will be placed in project root directory
  -n, --usevendor    usevendor mean this project init whit vendor,default use go-modules
  -m, --modulename=  modulename use for go modules init file: go.mod,default use project name
  -p, --targetpath=  project should init in the which directory,default is current path,if target directory not exists will be created
  -f, --configfile=  which init-config file should be use,if not set, default file will be download

Help Options:
  -h, --help         Show this help message

Arguments:
  projectname:       init the project with this name, the first name will be named for project,then all remaining names will be sub service name in cmd directory
```



## 使用说明



### 快速创建项目

最简单的命令如下：

```bash
tangzixiang$ goprojectinit myproject
tangzixiang$ tree myproject/
myproject/
├── README.md
├── bin
├── build
│   ├── ci
│   └── package
├── cmd
│   └── myproject
│       └── myproject.go
├── configs
├── deplyments
├── go.mod
├── init
├── internal
│   ├── api
│   ├── config
│   ├── handle
│   ├── model
│   ├── schema
│   ├── service
│   └── utils
├── pkg
│   ├── middleware
│   └── model
└── scripts
```

通过这么一个简单的命令，快速的初始化好一个 go 项目的目录环境，这里使用的是默认的目录配置。

上面的例子中项目 `myproject` 的项目入口文件为 `myproject/cmd/myproject/myproject.go`。



### 工具型项目

若新建项目为工具型项目则指定 `-t` 或者 `--istool` 参数即可:

```bash
tangzixiang$ goprojectinit myproject -t
tangzixiang$ tree myproject/
myproject/
├── README.md
├── bin
├── build
│   ├── ci
│   └── package
├── cmd
├── configs
├── deplyments
├── go.mod
├── init
├── internal
│   ├── api
│   ├── config
│   ├── handle
│   ├── model
│   ├── schema
│   ├── service
│   └── utils
├── main.go
├── pkg
│   ├── middleware
│   └── model
└── scripts
```

是否工具型项目的区别在于项目根目录是否存在 `main.go` 文件。



### 多服务项目

**goprojectinit** 支持指定新建项目为多服务项目,只需提供多个 `projectname` 即可,首个 `projectname` 将作为项目名。

```bash
tangzixiang$ goprojectinit myproject subservice
tangzixiang$ tree myproject
myproject/
├── README.md
├── bin
├── build
│   ├── ci
│   └── package
├── cmd
│   ├── myproject
│   │   └── myproject.go
│   └── subservice
│       └── subservice.go
├── deplyments
├── go.mod
├── init
├── internal
│   ├── api
│   ├── config
│   ├── handle
│   ├── model
│   ├── schema
│   ├── service
│   └── utils
├── main.go
├── pkg
│   ├── middleware
│   └── model
└── scripts
```



### vendor

**goprojectinit** 支持使用 vendor 作为模块处理，在初始化项目的时候带上 `-n` 或者 `--usevendor` 即可 。

```bash
tangzixiang$ goprojectinit myproject -n
tangzixiang$ tree myproject
myproject/
├── .git
├── .gitignore
├── README.md
├── bin
├── build
│   ├── ci
│   └── package
├── cmd
│   └── myproject
│       └── myproject.go
├── deplyments
├── init
├── internal
│   ├── api
│   ├── config
│   ├── handle
│   ├── model
│   ├── schema
│   ├── service
│   └── utils
├── pkg
│   ├── middleware
│   └── model
├── scripts
└── vendor
    └── vendor.json
```



### 展示更多调试信息

如果想要能够清楚的看到项目初始化过程的详细信息,只需要附带 `-b` 参数:

```bash
tangzixiang$ goprojectinit myproject -b

[goprojectinit] project path is "/XXX/tangzixiang/myproject"
[goprojectinit] make new directory /XXX/tangzixiang/myproject success~
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~
```



### 指定项目位置

如果不想在当前目录下初始化项目只需要通过 `-p` 或者 `--targetpath` 指定目标地址即可:

```bash
tangzixiang$ sudo goprojectinit myproject -p /var/www
```



### 自定义配置内容

配置文件可以指定如下内容:

1. `dir` : 指定项目目录结构，该参数指定结构文件路径，支持 `yaml` 以及 `json` 两种格式，可选。
2. `mainFileTemp` : main 文件的模板，该参数指定结构文件路径，可选。



上述配置项中路径参数可以是绝对路径或者是处于当前配置文件目录下的相对路径，若配置项未指定则会自动下载并使用默认配置举例如下:

```bash
tangzixiang$ tree configs/ ## 自定义配置文件
configs/
├── dir.json 			## 目录列表
├── main-file.temp  	## 模板文件
└── project-init.yaml 	## 配置文件
```



`project-init.yaml` 文件内容示例如下:

```bash
tangzixiang$ cat configs/project-init.yaml
## 项目初始化目录结构,可选
dir: dir.json
## main 文件的模板,可选
mainFileTemp: main-file.temp
```



`dir.json` 文件内容示例如下:

```bash
tangzixiang$ cat configs/dir.json
[
  "cmd",
  "configs",
  "internal/api",
  "internal/config",
  "internal/service",
  "internal/model",
  "internal/utils"
]
```



`main-file.temp` 文件内容示例如下:

```bash
tangzixiang$ cat configs/main-file.temp
{{/* main file */}}
package main

func main() {
    // todo everything
}
```



指定配置文件需要使用 `-f`  或者 `--configfile` :

```bash
tangzixiang$ goprojectinit myproject -f configs/project-init.yaml
```



### 覆盖项目路径

若新建的项目路径下存在同名目录，默认不会进行覆盖，会维持目录原样并退出初始化过程，如果需要进行覆盖则可以通过 `-c` 或者 `--cover` 完成:

```bash
tangzixiang$ ls -l /var/www/
drwxr-xr-x  5 xxx  xxx  160 11  9 21:01 myproject

tangzixiang$ goprojectinit myproject -f configs/project-init.yaml -p /var/www -c
[goprojectinit] are you sure to cover /var/www/myproject directory,type yes or no~
yes

tangzixiang:tangzixiang tangzixiang$ ls -l /var/www/
total 0
drwxr-xr-x  5 xxx  xxx  160 11  9 21:17 myproject
```

上面这个示例可以看到目录的最后变更时间发生了变化



##  注意

1. 当项目初始化过程中因为其他外界因素导致中途失败后，工具会自动清除创建到一半的项目。

2. 项目会默认带 `.configs` 目录以及 `cmd` 目录

   - 工具会将指定的初始化配置文件同步至 `.configs` 目录下。
   
   - 工具会根据使用 `goprojectinit` 命令指定的 `projectname`  参数在 `cmd` 目录下创建同样数量及对应名字的目录。
   
3. 初始化默认使用的是 `go-modules` 模式，需要 go 版本的支持，go 在 `go1.11` 版本开始支持
    - 初始化的 `go.mod` 文件的第一行为 `module {moduleName}`, 使用的 **moduleName** 默认为项目名，需要自定义的话可以附带 `-m` 参数指定自定义**moduleName** 。
    
    - 使用 `vendor` 模式需要带上 `-n` 参数，且确保项目位于 `GOPATH` 下。



## TODO

1. 支持指定 `vendor.json`
2. 工具初始化参数同时支持命令行及配置文件
3. 支持自动初始化 model 文件




## 参考：

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
