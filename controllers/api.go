package controllers

import (
	"allstats/lib"
	"allstats/models"
	"github.com/astaxie/beego/orm"
)

type ApiController struct {
	BaseController
}

type test_struct struct {
	Test string
}

func (c *ApiController) Agent() {
	jsoninfo := c.GetStrings("data")
	if len(jsoninfo) == 0 {
		c.Data["json"] = lib.JsonData(false, "Null", jsoninfo)
		c.ServeJSON()
		return
	}
	var task models.Task
	task.Task = jsoninfo[0]
	task.Ongoing = true
	orm.NewOrm().Insert(&task)
	c.Data["json"] = lib.JsonData(true, "Success", jsoninfo)
	c.ServeJSON()
	return

}
func (c *ApiController) Create() {

}

func (c *ApiController) Status() {

}

func (c *ApiController) Cancel() {

}