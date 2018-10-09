package routers

import (
	"github.com/astaxie/beego"
	"myGoProject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
