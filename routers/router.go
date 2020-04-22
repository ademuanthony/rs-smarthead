package routers

import (
	"smarthead/controllers"
	"github.com/astaxie/beego"
)

func init() { 
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pricing",&controllers.MainController{},"*:Pricing")
	beego.Router("/about",&controllers.MainController{},"*:About")
	beego.Router("/contact",&controllers.MainController{},"*:Contact")
	beego.Router("/support",&controllers.MainController{},"*:Support")
	beego.Router("/google6058e3992c01a0e3.html", &controllers.MainController{}, "*:GConsole")
	//User
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/signup", &controllers.UserController{}, "*:Register")
}
