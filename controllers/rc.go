package controllers

import (
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/niyanchun/k8s-middleware/models"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	"encoding/json"
)

// Operations about Replication Controllers
type RCController struct {
	BaseController
}

// TODO: 支持使用labels过滤查询到的RC
// @Title Get
// @Description get Replication Controller in specified namespace
// @Param namespace path string true "namespace of the rc resources"
// @router /:namespace [get]
func (rc *RCController) Get() {
	namespace := rc.GetString(":namespace")
	rc.CheckEmpty(namespace, "namespace")

	rc_list, err := models.Client.ListReplicationControllers(namespace)
	if err != nil {
		logs.Error("list rc error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	rc.Data["json"] = rc_list
	rc.ServeJSON()
}


// @Title Post
// @Description create Replication Controller
// @Param rc_params body models.ReplicationControllerCopy true "required params for RC"
// @router / [post]
func (rc *RCController) Post() {
	var err error
	var	rc_inst v1.ReplicationController

	err = json.Unmarshal(rc.Ctx.Input.RequestBody, &rc_inst)
	if err != nil {
		logs.Error("unmarshal rc params error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "unmarshal rc params error")
	}

	ret , err := models.Client.CreateReplicationController(rc_inst.Namespace, &rc_inst)
	if err != nil {
		logs.Error("create RC error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "create RC error")
	}

	rc.Data["json"] = ret
	rc.ServeJSON()
}


// @Title Delete
// @Description delete ReplicationController
// @Param namespace query string true "namespace"
// @Param rc_name path string true "replication controller name"
// @router /:rc_name [delete]
func (rc *RCController) Delete() {
	namespace := rc.GetString("namespace")
	rc_name := rc.GetString(":rc_name")
	rc.CheckEmpty(namespace, "namespace")
	rc.CheckEmpty(rc_name, "rc_name")

	err := models.Client.DeleteReplicationController(namespace, rc_name)
	if err != nil {
		logs.Error("delete RC error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "delete RC error")
	}
}
