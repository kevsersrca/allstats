package controllers


type PanelController struct {
	BaseController
}

func (c *PanelController) Index() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	c.TplName = "panel/index.html"
}