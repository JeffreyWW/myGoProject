package userControllers

import (
	"myGoProject/common"
	"myGoProject/controllers"
)

type LoginController struct {
	controllers.BaseController
}

// @router /login [get]
func (c *LoginController) Login() {
	println("进来了")
	token := jwtTool.CreateToken()
	c.Ctx.WriteString(token)

}
