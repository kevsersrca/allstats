package controllers

import "fmt"

type ApiController struct {
	BaseController
}

func (c *ApiController) Agent() {
	jsoninfo := c.GetString("data")
	if jsoninfo == "" {
		c.Ctx.WriteString("jsoninfo is empty")
		return
	}
	fmt.Println(jsoninfo)
}
func (c *ApiController) Create() {

}

func (c *ApiController) Status() {

}

func (c *ApiController) Cancel() {

}