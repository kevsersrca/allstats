package controllers

import (
	"allstats/lib"
	"allstats/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type ApiController struct {
	BaseController
}

type test_struct struct {
	Test string
}

func (c *ApiController) Agent() {

	decoder := json.NewDecoder(c.Ctx.Input.Context.Request.Body)

	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		c.Data["json"] = lib.JsonData(false, "Decode error", err.Error())
		c.ServeJSON()
		return
	}
	var task models.Task
	task.Task = t.Test
	task.Ongoing = true
	orm.NewOrm().Insert(&task)
	c.Data["json"] = lib.JsonData(true, "Success", t.Test)
	c.ServeJSON()
	return

}
func (c *ApiController) Create() {

}

func (c *ApiController) Status() {

}

func (c *ApiController) Cancel() {

}