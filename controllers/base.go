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

func (bc *BaseController) CheckError(err error, errMsg string, errCode int) {
	if err != nil {
		s := errMsg + "," + err.Error()
		logs.Error(s)
		bc.CustomAbort(errCode, s)
	}
}