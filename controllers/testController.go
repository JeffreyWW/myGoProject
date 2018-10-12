package controllers

type TestController struct {
	BaseController
}

// @router /test [post]
func (c *TestController) Test() {
	c.TplName = "index.tpl"
	//c.Abort("405")
}
