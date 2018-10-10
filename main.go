package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myGoProject/controllers"
	"myGoProject/models"
	_ "myGoProject/routers"
)

func init() {
	dataBaseConfig()
	beego.ErrorController(&controllers.ErrorController{})
}
func main() {
	beego.Run()
}

func dataBaseConfig() {
	/**默认自带,可以不写下面这句*/
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:crfchina@tcp(localhost:3306)/goProject?charset=utf8", 30)
	orm.RegisterModel(new(models.User))

	/**自动创建表,个人建议还是先建好表再处理逻辑*/
	//orm.RunSyncdb("default", false, true)
}
