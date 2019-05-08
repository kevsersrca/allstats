package controllers

import (
	"github.com/astaxie/beego/orm"
	"allstats/models"
	"fmt"
)

type PanelController struct {
	BaseController
	Count bool
}

func (c *PanelController) IsServerCount() {
	u := &models.User{Id: c.Userinfo.Id}
	o := orm.NewOrm().QueryTable("Server").Filter("User", u)
	count, _ := o.Count()
	if count == 0 {
		c.Count = false
	}
	c.Count = true
}

func (c *PanelController) Index() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	fmt.Println(c.Count)
	c.Data["Title"] = "Dashboard"
	c.Data["Count"] = c.Count
	c.TplName = "panel/index.html"
}

func (c *PanelController) Servers() {
	c.Data["Count"] = c.Count
	c.Data["Title"] = "Servers"
	c.TplName = "panel/servers.html"
}

func (c *PanelController) ServerShow() {
	//id, _ := c.GetInt64(":id")
	c.Data["Count"] = c.Count
	c.Data["Title"] = "Server 1"
	c.TplName = "panel/show.html"
}

func (c *PanelController) Account() {
	User := &models.User{Id: c.Userinfo.Id}
	if err := orm.NewOrm().Read(User); err == nil {
		c.Data["Name"] = User.Name
		c.Data["Email"] = User.Email
		c.Data["Password"] = User.Password
		c.Data["CreatedDate"] = User.Created
	}
	c.Data["Title"] = "User Profile"
	c.TplName = "panel/account.html"
}

func (c *PanelController) Api() {
	c.Data["Title"] = "API"
	c.TplName = "panel/api.html"
}