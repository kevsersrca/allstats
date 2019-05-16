package controllers

import (
	"allstats/lib"
	"allstats/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/base64"
	"time"
)

type ApiController struct {
	BaseController
}

func Base64Convert(data string) string{
	temp , _:= base64.StdEncoding.DecodeString(data)
	return string(temp)
}

func (c *ApiController) Agent() {
	auth := c.GetString("token")
	if auth == "" {
		c.Data["json"] = lib.JsonData("ERROR", "Auth Required!")
		c.ServeJSON()
		return
	}
	data := c.GetStrings("data")
	if len(data) == 0 {
		c.Data["json"] = lib.JsonData("ERROR", "Data Not Found!")
		c.ServeJSON()
		return
	}

	var statics models.ServerStatics

	datas := strings.Split(data[0], " ")
	statics.Version = Base64Convert(datas[0])
	statics.Uptime = Base64Convert(datas[1])
	statics.Sessions = Base64Convert(datas[2])
	statics.Processes = Base64Convert(datas[3])
	statics.ProcessesArray = Base64Convert(datas[4])
	statics.FileHandles = Base64Convert(datas[5])
	statics.FileHandlesLimit = Base64Convert(datas[6])
	statics.OsKernel = Base64Convert(datas[7])
	statics.OsName = Base64Convert(datas[8])
	statics.OsArch = Base64Convert(datas[9])
	statics.CpuName = Base64Convert(datas[10])
	statics.CpuCores = Base64Convert(datas[11])
	statics.CpuFreq = Base64Convert(datas[12])
	statics.RamTotal = Base64Convert(datas[13])
	statics.RamUsage = Base64Convert(datas[14])
	statics.SwapTotal = Base64Convert(datas[15])
	statics.SwapUsage = Base64Convert(datas[16])
	statics.DiskArray = Base64Convert(datas[17])
	statics.DiskTotal = Base64Convert(datas[18])
	statics.DiskUsage = Base64Convert(datas[19])
	statics.Connections = Base64Convert(datas[20])
	statics.Nic = Base64Convert(datas[21])
	statics.IPV4 = Base64Convert(datas[22])
	statics.IPV6 = Base64Convert(datas[23])
	statics.Rx = Base64Convert(datas[24])
	statics.Tx = Base64Convert(datas[25])
	statics.RxGap = Base64Convert(datas[26])
	statics.TxGap = Base64Convert(datas[27])
	statics.Load = Base64Convert(datas[28])
	statics.LoadCpu = Base64Convert(datas[29])
	statics.LoadIO = Base64Convert(datas[30])
	statics.PingEu = Base64Convert(datas[31])
	statics.PingUs = Base64Convert(datas[32])
	statics.PingAs = Base64Convert(datas[33])

	orm.NewOrm().Insert(&statics)
	interval := time.Now().Format(time.RFC850)
	c.Data["json"] = lib.JsonData("ok", interval)
	c.ServeJSON()
	return
}
func (c *ApiController) Create() {

}

func (c *ApiController) Status() {

}

func (c *ApiController) Cancel() {

}
