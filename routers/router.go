package routers

import (
	"allstats/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/services", &controllers.MainController{}, "get:Services")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/signup", &controllers.LoginController{}, "get,post:Signup")

	//Panel
	beego.Router("/panel", &controllers.PanelController{}, "get:Index")

	//Api
	beego.Router("/api/agent.json", &controllers.ApiController{}, "post:Index")

}
