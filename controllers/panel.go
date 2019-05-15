package controllers

import (
	"allstats/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"io"
	"crypto/rand"
)

type PanelController struct {
	BaseController
	Count bool
}


func (c *PanelController) IsServerCount() {
	u := &models.User{Id: c.Userinfo.Id}
	o := orm.NewOrm().QueryTable("Server").Filter("Account", u)
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
	c.IsServerCount()
	c.Data["Title"] = "Dashboard"
	c.Data["Count"] = c.Count
	c.TplName = "panel/index.html"
}

func RandomHash(length int) string {
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	new_pword := make([]byte, length)
	random_data := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, random_data); err != nil {
			beego.Error(err)
		}
		for _, c := range random_data {
			if c >= maxrb {
				continue
			}
			new_pword[i] = chars[c%clen]
			i++
			if i == length {
				return string(new_pword)
			}
		}
	}
	return "123qwe"
}

func (c *PanelController) Servers() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	c.IsServerCount()
	flash := beego.NewFlash()
	o := orm.NewOrm().QueryTable("Server").Filter("Account", c.Userinfo)
	var server []orm.Params
	_, err := o.Values(&server, "Id","Uuid", "Account__Name", "Name", "Systemlimit", "DiskLimit","RamLimit","Created", "Status" )
	if err != nil {
		beego.Error(err)
		flash.Error("Hata!")
		flash.Store(&c.Controller)
		return
	}
	c.Data["Count"] = c.Count
	c.Data["Title"] = "Servers"
	c.Data["server"] = server
	c.TplName = "panel/servers.html"
}

func (c *PanelController) Create() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	c.IsServerCount()
	c.Data["Title"] = "Servers"
	c.Data["Count"] = c.Count
	c.TplName = "panel/servers.html"
	flash := beego.NewFlash()

	var server models.Server
	server.Name = c.GetString("name")
	server.Uuid = RandomHash(32)
	server.Systemlimit = c.GetString("loadusage")
	server.RamLimit = c.GetString("ramusage")
	server.DiskLimit = c.GetString("diskusage")
	server.Account = c.Userinfo
	server.Status = "Active"
	_, err := orm.NewOrm().Insert(&server)
	if err != nil {
		beego.Error(err)
		flash.Error("Hata!")
		flash.Store(&c.Controller)
		return
	}
	flash.Notice("wget -N --no-check-certificate http://allstats.in/static/scripts/allstats-install.sh  && bash allstats-install.sh "+ server.Uuid)
	flash.Store(&c.Controller)
	return
}

func (c *PanelController) ServerShow() {
	id := c.GetString(":id")
	c.Data["Count"] = c.Count
	c.Data["Title"] = "Server "+id
	c.TplName = "panel/show.html"
}

func (c *PanelController) Account() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	User := &models.User{Id: c.Userinfo.Id}
	if err := orm.NewOrm().Read(User); err == nil {
		c.Data["Name"] = User.Name
		c.Data["Id"] = User.Id
		c.Data["Email"] = User.Email
		c.Data["Password"] = User.Password
		c.Data["CreatedDate"] = User.Created
	}
	c.Data["Title"] = "User Profile"
	c.TplName = "panel/account.html"
}

func (c *PanelController) ChangePassword() {

}

func (c *PanelController) Api() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
	}
	c.TplName = "panel/api.html"
	flash := beego.NewFlash()
	o := orm.NewOrm().QueryTable("Api").Filter("User", c.Userinfo)
	var api []orm.Params
	_, err := o.Values(&api, "Key", "User__Name", "Active", "Created")
	if err != nil {
		beego.Error(err)
		flash.Success("Not Found Api Key!")
		flash.Store(&c.Controller)
		return
	}
	c.Data["api"] = api
	c.Data["Title"] = "API"
}
