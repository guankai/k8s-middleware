package main

import (
	_ "github.com/niyanchun/k8s-middleware/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	_ "github.com/niyanchun/k8s-middleware/models"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	if beego.BConfig.RunMode == "prod" {
		beego.SetLogFuncCall(true)
		logs.SetLogger(logs.AdapterFile, `{"filename":"/log/k8s-middleware.log"}`)
	}

	beego.Run()
}
