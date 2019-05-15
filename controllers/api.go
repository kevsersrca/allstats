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
	auth := c.GetString("post_data")
	if auth == "" {
		c.Data["json"] = lib.JsonData(false, "Auth required!", auth)
		c.ServeJSON()
		return
	}

	var task models.Task
	task.Task = auth
	task.Ongoing = true
	_, err := orm.NewOrm().Insert(&task)
	if err != nil {
		c.Data["json"] = lib.JsonData(false, "Task not saved!", err)
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