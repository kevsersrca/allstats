package controllers

import (
	"allstats/lib"
	"allstats/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"encoding/base64"
)

type ApiController struct {
	BaseController
}

type test_struct struct {
	Test string
}

func (c *ApiController) Agent() {
	auth := c.GetString("token")
	if auth == "" {
		c.Data["json"] = lib.JsonData(false, "Auth Required!", auth)
		c.ServeJSON()
		return
	}
	data := c.GetStrings("data")
	if len(data) == 0 {
		c.Data["json"] = lib.JsonData(false, "Data Not Found!", data)
		c.ServeJSON()
		return
	}

	var statics models.ServerStatics

	datas := strings.Split(data[0], " ")
	statics.Version, _ = base64.StdEncoding.DecodeString(datas[0])
	statics.Uptime, _ = base64.StdEncoding.DecodeString(datas[1])
	statics.Sessions, _ = base64.StdEncoding.DecodeString(datas[2])
	statics.Processes, _ = base64.StdEncoding.DecodeString(datas[3])
	statics.ProcessesArray, _ = base64.StdEncoding.DecodeString(datas[4])
	statics.FileHandles, _ = base64.StdEncoding.DecodeString(datas[5])
	statics.FileHandlesLimit, _ = base64.StdEncoding.DecodeString(datas[6])
	statics.OsKernel, _ = base64.StdEncoding.DecodeString(datas[7])
	statics.OsName, _ = base64.StdEncoding.DecodeString(datas[8])
	statics.OsArch, _ = base64.StdEncoding.DecodeString(datas[9])
	statics.CpuName, _ = base64.StdEncoding.DecodeString(datas[10])
	statics.CpuCores, _ = base64.StdEncoding.DecodeString(datas[11])
	statics.CpuFreq, _ = base64.StdEncoding.DecodeString(datas[12])
	statics.RamTotal, _ = base64.StdEncoding.DecodeString(datas[13])
	statics.RamUsage, _ = base64.StdEncoding.DecodeString(datas[14])
	statics.SwapTotal, _ = base64.StdEncoding.DecodeString(datas[15])
	statics.SwapUsage, _ = base64.StdEncoding.DecodeString(datas[16])
	statics.DiskArray, _ = base64.StdEncoding.DecodeString(datas[17])
	statics.DiskTotal, _ = base64.StdEncoding.DecodeString(datas[18])
	statics.DiskUsage, _ = base64.StdEncoding.DecodeString(datas[19])
	statics.Connections, _ = base64.StdEncoding.DecodeString(datas[20])
	statics.Nic, _ = base64.StdEncoding.DecodeString(datas[21])
	statics.IPV4, _ = base64.StdEncoding.DecodeString(datas[22])
	statics.IPV6, _ = base64.StdEncoding.DecodeString(datas[23])
	statics.Rx, _ = base64.StdEncoding.DecodeString(datas[24])
	statics.Tx, _ = base64.StdEncoding.DecodeString(datas[25])
	statics.RxGap, _ = base64.StdEncoding.DecodeString(datas[26])
	statics.TxGap, _ = base64.StdEncoding.DecodeString(datas[27])
	statics.Load, _ = base64.StdEncoding.DecodeString(datas[28])
	statics.LoadCpu, _ = base64.StdEncoding.DecodeString(datas[29])
	statics.LoadIO, _ = base64.StdEncoding.DecodeString(datas[30])
	statics.PingEu, _ = base64.StdEncoding.DecodeString(datas[31])
	statics.PingUs, _ = base64.StdEncoding.DecodeString(datas[32])
	statics.PingAs, _ = base64.StdEncoding.DecodeString(datas[33])

	orm.NewOrm().Insert(&statics)
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
