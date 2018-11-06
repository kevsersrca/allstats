package main

import (
	_ "allstats/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPSCertFile = "/etc/letsencrypt/live/allstats.in/cert.pem"
	beego.BConfig.Listen.HTTPSKeyFile = "/etc/letsencrypt/live/allstats.in/privkey.pem"
	// beego.BConfig.Listen.HTTPSAddr = ""
	// beego.BConfig.Listen.HTTPSPort = 443
	// beego.BConfig.Listen.EnableHTTPS = true
	beego.Run()
}

