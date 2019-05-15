package controllers

import (
	"allstats/lib"
	"allstats/models"
	"github.com/astaxie/beego/orm"
)

type ApiController struct {
	BaseController
}


func (c *ApiController) Agent() {
	auth := c.GetString("token")
	data := c.GetString("data")
	if auth == "" {
		c.Data["json"] = lib.JsonData(false, "Auth required!")
		c.ServeJSON()
		return
	}
	task := new(models.Task)
	task.Task = data
	task.Ongoing = true
	if _, err := orm.NewOrm().Insert(task); err != nil {
		c.Data["json"] = lib.JsonData(false, "Task not saved!", err.Error())
		c.ServeJSON()
		return
	}
	c.Data["json"] = lib.JsonData(true, "Success")
	c.ServeJSON()
	return
}
func (c *ApiController) Create() {

}

func (c *ApiController) Status() {

}

func (c *ApiController) Cancel() {

}