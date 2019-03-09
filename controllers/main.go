package controllers


type MainController struct {
	BaseController
}

func (c *MainController) Index() {
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
