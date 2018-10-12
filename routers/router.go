package routers

import (
	"github.com/astaxie/beego"
	"myGoProject/controllers"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.AccountController{})
	beego.Include(&controllers.TestController{})
}
