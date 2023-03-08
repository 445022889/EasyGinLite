package conf

import (
	"fmt"
	"github.com/go-ini/ini"
)

// Mysql mysql的配置结构体
type Mysql struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

// Redis redis的配置结构体
type Redis struct {
	Addr     string `ini:"addr"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

// Config 可以自己再添加，我这只加了redis和mysql
type Config struct {
	Redis `ini:"redis"`
	Mysql `ini:"mysql"`
}

// ProjectConfig 全局都会调用的配置
var ProjectConfig *Config

// InitConfig 初始化配置
func InitConfig() {
	ProjectConfig = new(Config)
	err := ini.MapTo(ProjectConfig, "./config/conf.ini")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
