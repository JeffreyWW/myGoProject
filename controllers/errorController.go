package controllers

type Error struct {
	Code int
	error
	reason  string
	message string
}

func (e *Error) Error() string {
	return "fuck"
}

type ErrorController struct {
	BaseController
}

func (c *ErrorController) ErrorTest() {

	println("进来这里了")
}

func (c *ErrorController) Finish() {
	c.Ctx.WriteString("fuck")

	println("结束")
}
