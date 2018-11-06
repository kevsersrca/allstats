package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Website"] = "allstats.in"
	c.Data["Email"] = "kevser.sirca@gmail.com"
	c.TplName = "index.html"
}

func (c *MainController) About() {
	c.TplName = "about.html"
}

func (c *MainController) Services() {
	c.TplName = "services.html"
}

func (c *MainController) Contact() {
	c.TplName = "contact.html"
}