# Go 项目工程初始化工具



`go get github.com/tangzixiang/goprojectinit`



## 概述

**goprojectinit** 是一个可以快速初始化 go 项目环境的工具，该工具具体以下特色：

- 支持自定义需要初始化的目录环境。
- 主动创建好项目单个或多个入口文件。
- 支持指定项目的初始化路径，指定的项目目录已存在不会自动覆盖，若需要覆盖只需要添加一个参数即可。



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
tangzixiang$ goprojectinit myproject
tangzixiang$ tree myproject/
myproject/
├── bin
├── build
│   ├── ci
│   └── package
├── cmd
├── configs
│   ├── dir.json
│   └── project-init.yaml
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
└── scripts
```

通过这么一个简单的命令，快速的初始化好一个 go 项目的目录环境，这里使用的是默认的目录配置



如果想要能够清楚的看到项目初始化过程的信息,只需要附带 `-v` 参数:

```bash
tangzixiang$ goprojectinit myproject -v

[goprojectinit] project path is "/XXX/tangzixiang/myproject"
[goprojectinit] make new directory /XXX/tangzixiang/myproject success~
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~
```



如果不想在当前目录下初始化项目只需要通过 `-p` 或者 `--targetpathdir` 指定目标地址即可:

```bash
tangzixiang$ sudo goprojectinit myproject -v -p /var/www

[goprojectinit] project path is "/var/www/myproject"
[goprojectinit] make new directory /var/www/myproject success~
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~
```



如果需要配置自定义的目录环境需要使用 `-f`  或者 `--configfile` 指定配置文件地址:

```bash
tangzixiang$ tree configs/ ## 自定义配置文件
configs/
├── dir.json          ## 目录列表
└── project-init.yaml ## 配置文件

tangzixiang$ cat configs/project-init.yaml
# 项目初始化目录结构，相对路径以当前配置文件所在目录为基准
dir: dir.json

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

tangzixiang$ goprojectinit myproject -v -f configs/project-init.yaml -p /var/www
[goprojectinit] project path is "/var/www/myproject"
[goprojectinit] make new directory /var/www/myproject success~
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~

tangzixiang$ sudo tree /var/www
/var/www
└── myproject
    ├── cmd
    ├── configs
    │   ├── dir.json
    │   └── project-init.yaml
    └── internal
        ├── api
        ├── config
        ├── model
        ├── service
        └── utils

9 directories, 2 files
```



若新建的项目路径下存在同名目录，默认不会进行覆盖，会维持目录原样并退出初始化过程，如果需要进行覆盖则可以通过 `-c` 或者 `--cover` 完成:

```bash
tangzixiang$ ls -l /var/www/
drwxr-xr-x  5 xxx  xxx  160 11  9 21:01 myproject

tangzixiang$ sudo goprojectinit myproject -v -f configs/project-init.yaml -p /var/www -c
Password:
[goprojectinit] project path is "/var/www/myproject"
[goprojectinit] are you sure to cover /var/www/myproject directory,type yes or no~
yes
//...忽略部分输出日志
[goprojectinit] projoct myproject init success~

tangzixiang:tangzixiang tangzixiang$ ls -l /var/www/
total 0
drwxr-xr-x  5 xxx  xxx  160 11  9 21:17 myproject
```



##  注意

1. 当项目初始化过程中因为其他外界因素导致中途失败后，工具会自动清除创建到一半的项目。

2. 项目会默认带 `configs` 目录以及 `cmd` 目录

   1. 工具会将指定的初始化配置文件同步至 `configs` 目录下。
   2. 工具会根据使用 `goprojectinit` 命令指定的 `projectname`  参数在 `cmd` 目录下创建同样数量及对应名字的目录。



## TODO

1. 自动初始化 `git` 环境
2. 自动初始化 `vendor` 环境，并且支持指定 `vendor.json`
3. 工具初始化参数同时命令行及配置文件
4. 支持自动初始化 model 文件




## 参考：

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)