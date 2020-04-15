package routers

import (
	"smarthead/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pricing",&controllers.MainController{},"*:Pricing")
	beego.Router("/support",&controllers.MainController{},"*:Support")
}
