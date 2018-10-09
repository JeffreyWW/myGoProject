package controllers

import "github.com/astaxie/beego"

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	println("准备")

	if controller, ok := c.AppController.(NestPreparer); ok {
		controller.NestPrepare()
	}

}
