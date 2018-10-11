package controllers

var (
	ErrorExistUser   = "ExistUser"
	ErrorSystemError = "SystemError"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) ErrorExistUser() {
	c.Response.Code = 1000
	c.Response.Message = "该用户已存在"
	c.Response.Reason = "插入用户数据失败"
}

func (c *ErrorController) ErrorSystemError() {
	c.Response.Code = 2000
	c.Response.Message = "系统异常,请稍后再试"
	c.Response.Reason = "未知错误,统一处理"
}

type AccountError interface {
	ErrorExistUser()
}
