package main

import (
	_ "github.com/niyanchun/k8s-middleware/routers"

	"github.com/astaxie/beego"

	_ "github.com/niyanchun/k8s-middleware/models"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.SetLogFuncCall(true)
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/registrywebservice.log"}`)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
