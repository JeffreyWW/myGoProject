package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"myGoProject/common"
	"myGoProject/models"
)

type AccountController struct {
	BaseController
}

//@router /register [get]
func (c *AccountController) Register() {
	//c.Ctx.WriteString("fuck")
	//
	c.Abort("Test")

	println("手机号是" + c.GetString("phoneNo"))
	data := []byte(c.GetString("password"))
	md5Password := fmt.Sprintf("%x", md5.Sum(data))
	println(md5Password)
	phoneNo, _ := c.GetInt("phoneNo")
	o := orm.NewOrm()
	user := models.User{PhoneNo: phoneNo, Password: md5Password}
	id, err := o.Insert(&user)
	println(id, err)

	c.Ctx.WriteString(jwtTool.CreateToken())

}

func (c *AccountController) Login() {

}
