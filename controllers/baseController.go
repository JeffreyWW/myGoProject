package controllers

import (
	"github.com/astaxie/beego"
	"myGoProject/models"
)

//type NestPreparer interface {
//	NestPrepare()
//}

type BaseController struct {
	beego.Controller
	Response models.Response
}

func (c *BaseController) Finish() {
	c.Data["json"] = &c.Response
	c.ServeJSON()
}

//func (c *BaseController) Prepare() {
//	println("准备")
//	if controller, ok := c.AppController.(NestPreparer); ok {
//		controller.NestPrepare()
//	}
//}
