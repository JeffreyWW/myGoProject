package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myGoProject/controllers:AccountController"] = append(beego.GlobalControllerRouter["myGoProject/controllers:AccountController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/register`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["myGoProject/controllers:MainController"] = append(beego.GlobalControllerRouter["myGoProject/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/*`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["myGoProject/controllers:MainController"] = append(beego.GlobalControllerRouter["myGoProject/controllers:MainController"],
		beego.ControllerComments{
			Method:           "TTT",
			Router:           `/*`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["myGoProject/controllers:TestController"] = append(beego.GlobalControllerRouter["myGoProject/controllers:TestController"],
		beego.ControllerComments{
			Method:           "Test",
			Router:           `/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
