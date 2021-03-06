package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/go-sql-driver/mysql"
	"myGoProject/models"
)

type AccountController struct {
	BaseController
}

//@router /register [get]
func (c *AccountController) Register() {
	m := models.Response{}
	m.Result = map[string]string{}
	data := []byte(c.GetString("password"))
	md5Password := fmt.Sprintf("%x", md5.Sum(data))
	phoneNo, _ := c.GetInt("phoneNo")
	o := orm.NewOrm()
	user := models.User{PhoneNo: phoneNo, Password: md5Password}
	valid := validation.Validation{}
	b, er := valid.Valid(&user)
	if er != nil {
		println("f")
	}

	if !b {
		for _, err := range valid.Errors {
			println(err.Key, err.Message)
			println("")

		}
	}

	_, err := o.Insert(&user)
	dbError, _ := err.(*mysql.MySQLError)
	if dbError != nil {
		switch dbError.Number {
		case 1062:
			c.Abort(ErrorExistUser)
			break
		default:
			c.Abort(ErrorSystemError)
			break
		}
	}
	c.Response.Result = &user
}
