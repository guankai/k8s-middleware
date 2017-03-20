package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

// check if param is empty
func (bc *BaseController) CheckEmpty(param, paramToLog string) {
	if len(param) == 0 {
		s := fmt.Sprintf("%s is empty", paramToLog)
		logs.Error(s)
		bc.CustomAbort(http.StatusBadRequest, s)
	}
}
