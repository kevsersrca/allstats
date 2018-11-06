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
}
