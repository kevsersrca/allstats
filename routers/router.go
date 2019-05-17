package routers

import (
	"allstats/controllers"

	"github.com/astaxie/beego"
	"allstats/controllers/manage"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/about", &controllers.MainController{}, "get:About")
	beego.Router("/services", &controllers.MainController{}, "get:Services")
	beego.Router("/contact", &controllers.MainController{}, "get:Contact")
	//Auth
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/signup", &controllers.LoginController{}, "get,post:Signup")

	//Panel
	beego.Router("/panel", &controllers.PanelController{}, "get:Index")
	beego.Router("/panel/servers", &controllers.PanelController{}, "get:Servers")
	beego.Router("/panel/servers/create", &controllers.PanelController{}, "post:Create")
	beego.Router("/panel/server/:id([0-9]+)", &manage.ManageController{}, "get:Show")
	beego.Router("/panel/account", &controllers.PanelController{}, "get:Account")
	beego.Router("/panel/api", &controllers.PanelController{}, "get:Api")

	//Api
	beego.Router("/api/agent.json", &controllers.ApiController{}, "get,post:Agent")

}
