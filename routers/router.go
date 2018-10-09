package routers

import (
	"github.com/astaxie/beego"
	"myGoProject/controllers"
	"myGoProject/controllers/userControllers"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&userControllers.LoginController{})
}
