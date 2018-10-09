package controllers

type MainController struct {
	BaseController
}

//@router /main [get]
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) NestPrepare() {
	println("子准备")
}
