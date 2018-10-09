package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["myGoProject/controllers:MainController"] = append(beego.GlobalControllerRouter["myGoProject/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/main`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
