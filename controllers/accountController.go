package controllers

import (
	"crypto/md5"
	"fmt"
	"myGoProject/common"
)

type AccountController struct {
	BaseController
}

//@router /register [get]
func (c *AccountController) Register() {
	println("手机号是" + c.GetString("phoneNo"))
	data := []byte(c.GetString("password"))
	md5Password := fmt.Sprintf("%x", md5.Sum(data))
	println(md5Password)
	c.Ctx.WriteString(jwtTool.CreateToken())

}

func (c *AccountController) Login() {

}
