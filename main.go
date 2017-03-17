package main

import (
	_ "github.com/niyanchun/k8s-client/routers"

	"github.com/astaxie/beego"

	_ "github.com/niyanchun/k8s-client/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
