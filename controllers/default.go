package controllers

type MainController struct {
	BaseController
}

//@router /* [post]
func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Abort("503")
	//c.TplName = "dist/index.html"

}

//@router /* [get]
func (c *MainController) TTT() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.Abort("503")
	c.TplName = "dist/index.html"

}
