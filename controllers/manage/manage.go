package manage

import (
	"allstats/controllers"
	"allstats/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"

	"allstats/lib"
	"strconv"
	"strings"
)

type ManageController struct {
	controllers.PanelController
	server models.Server
	statics models.ServerStatics
}

func (c *ManageController) NestPrepare() {
	if !c.IsLogin {
		if c.Ctx.Input.URI() != c.URLFor("LoginController.Login") {
			c.Ctx.SetCookie("LoginReturnUrl", c.Ctx.Input.URI(), 0, "/")
		}
		c.Redirect(c.URLFor("LoginController.Login"), 302)
	}
	id, _ := c.GetInt64(":id")
	_, err := orm.NewOrm().QueryTable("Server").Filter("Id", id).All(&c.server)
	if err != nil{
		beego.Error("Server not found error!")
		return
	}

}

func (c *ManageController) Show() {
	c.TplName = "panel/manage/show.html"
	flash := beego.NewFlash()

	o := orm.NewOrm().QueryTable("ServerStatics").Filter("Server__id", c.server.Id)
	var Statics []models.ServerStatics
	_, err := o.All(&Statics)
	if err != nil || Statics == nil{
		beego.Error(err)
		flash.Warning("Statics not found!")
		flash.Store(&c.Controller)
		return
	}
	if len(Statics) == 0 {
		beego.Error(err)
		flash.Warning("Statics not found!")
		flash.Store(&c.Controller)
		return
	}
	var data1 []int
	var data2 []int
	var data3 []float64
	var data4 []float64
	var data5 []float64
	var data6 []int
	var data7 []int
	var data8 []int
	var data9 []int
	var data10 []int
	var data12 []string

	for i,v := range Statics {
		if i < 20 {
			data1 = append(data1, conv(v.Rx))
			data2 = append(data2, conv(v.Tx))
			data3 = append(data3, strtoint(v.PingAs))
			data4 = append(data4, strtoint(v.PingUs))
			data5 = append(data5, strtoint(v.PingEu))
			data6 = append(data6, conv(v.Load))
			data7 = append(data7, conv(v.LoadCpu))
			data8 = append(data8, conv(v.LoadIO))
			data9 = append(data9, conv(v.RamUsage))
			data10 = append(data10, conv(v.SwapUsage))
		}
	}

	temp := strings.Split(Statics[0].ProcessesArray, ";")
	for _,v := range temp {
		data12 = append(data12, v)
	}
	year, month, day :=c.server.Created.Date()
	c.Data["Load"] = Statics[0].Load
	c.Data["OS"] = Statics[0].OsName
	c.Data["Kernel"] = Statics[0].OsKernel
	c.Data["CpuModel"] = Statics[0].CpuName
	c.Data["RamTotal"] = conv(Statics[0].RamTotal)
	c.Data["DiskTotal"] = conv(Statics[0].DiskTotal)
	c.Data["RamUsage"] = conv(Statics[0].RamUsage)
	c.Data["DiskUsage"] = conv(Statics[0].DiskUsage)
	c.Data["Data1"] = data1
	c.Data["Data2"] = data2
	c.Data["Data3"] = data3
	c.Data["Data4"] = data4
	c.Data["Data5"] = data5
	c.Data["Data6"] = data6
	c.Data["Data7"] = data7
	c.Data["Data8"] = data8
	c.Data["Data9"] = data9
	c.Data["Data10"] = data10
	c.Data["Count"] = c.Count
	c.Data["s"] = data12
	c.Data["Title"] = c.server.Name
	c.Data["ipv4"] = c.server.Ipv4
	c.Data["Status"] = c.server.Status
	c.Data["ipv6"] = c.server.Ipv6
	c.Data["Mac"] = c.server.Mac
	c.Data["Created"] = fmt.Sprintf("%d-%s-%d", year,month.String(),day)
}

func conv(val string) int {
	value , _ := strconv.Atoi(val)
	b,_ := lib.Convert(uint64(value))
	return int(b)
}

func strtoint(val string) float64 {
	value , _ := strconv.ParseFloat(val, 64)
	return value
}