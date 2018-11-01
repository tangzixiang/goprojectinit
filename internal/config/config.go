package config

// Config 配置
type Config struct {
	DirPath string `yaml:"dir"`
}

// 初始化目录
var DirM = make(map[string]string)


func init() {

}
