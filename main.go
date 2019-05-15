package main

import (
	_ "allstats/routers"
	_ "allstats/inits"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPSCertFile ="cert.pem"
	beego.BConfig.Listen.HTTPSKeyFile = "key.pem"
	beego.BConfig.Listen.HTTPSAddr = ""
	beego.BConfig.Listen.HTTPSPort = 443
	beego.BConfig.Listen.EnableHTTPS = true
	beego.Run()
}
