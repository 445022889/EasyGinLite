package utils

import (
	conf "EasyGinLite/config"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MysqlInit 这里初始化一个mysql的操作实例，具体使用方法 https://goframe.org/pages/viewpage.action?pageId=1114686
func MysqlInit(Debug bool) (md gdb.DB) {
	gdb.SetConfig(gdb.Config{"default": gdb.ConfigGroup{gdb.ConfigNode{
		Host:  conf.ProjectConfig.Mysql.Host,
		Port:  conf.ProjectConfig.Mysql.Port,
		User:  conf.ProjectConfig.Mysql.User,
		Pass:  conf.ProjectConfig.Mysql.Password,
		Name:  conf.ProjectConfig.Mysql.Database,
		Type:  "mysql",
		Debug: Debug,
	}}})
	md = g.DB()
	return
}
