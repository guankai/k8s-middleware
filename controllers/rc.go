package controllers

import (
	"net/http"

	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/niyanchun/k8s-middleware/models"
)

// Operations about Replication Controllers
type RCController struct {
	BaseController
	namespace string
}

func (rc *RCController) Prepare() {
	namespace := rc.GetString("namespace")

	method := rc.Ctx.Input.Method()
	if method != http.MethodPost && method != http.MethodPut {
		rc.CheckEmpty(namespace, "namespace")
	}

	logs.Debug("namespace: %s", namespace)
	rc.namespace = namespace
}

// TODO: 支持使用labels过滤查询到的RC
// @Title Get
// @Description List all Replication Controller in specified namespace
// @Param namespace query string true "namespace of the rc resources"
// @router / [get]
func (rc *RCController) List() {
	rc_list, err := models.Client.ListReplicationControllers(rc.namespace)
	if err != nil {
		logs.Error("list rc error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	rc.Data["json"] = rc_list
	rc.ServeJSON()
}

// @Title Get
// @Description Get Replication Controller details
// @Param namespace query string true "namespace of the rc resources"
// @Param rc_name path string true "replicationController name"
// @router /:rc_name [get]
func (rc *RCController) Get() {
	rc_name := rc.GetString(":rc_name")
	rc.CheckEmpty(rc_name, "rc_name")

	rc_detail, err := models.Client.GetReplicationController(rc.namespace, rc_name)
	if err != nil {
		logs.Error("get RC error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "get RC error")
	}

	rc.Data["json"] = rc_detail
	rc.ServeJSON()
}

// @Title Post
// @Description create Replication Controller
// @Param rc_params body models.ReplicationControllerCopy true "required params for RC"
// @router / [post]
func (rc *RCController) Post() {
	var err error
	var rc_inst v1.ReplicationController

	err = json.Unmarshal(rc.Ctx.Input.RequestBody, &rc_inst)
	if err != nil {
		logs.Error("unmarshal rc params error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "unmarshal rc params error")
	}

	ret, err := models.Client.CreateReplicationController(rc_inst.Namespace, &rc_inst)
	if err != nil {
		logs.Error("create RC error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "create RC error")
	}

	rc.Data["json"] = ret
	rc.ServeJSON()
}

// @Title Put
// @Description update Replication Controller
// @Param rc_params body models.ReplicationControllerCopy true "required params for RC"
// @router / [put]
func (rc *RCController) Put() {
	var err error
	var rc_inst v1.ReplicationController

	err = json.Unmarshal(rc.Ctx.Input.RequestBody, &rc_inst)
	if err != nil {
		logs.Error("unmarshal rc params error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "unmarshal rc params error")
	}

	ret, err := models.Client.UpdateReplicationController(rc_inst.Namespace, &rc_inst)
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
	rc_name := rc.GetString(":rc_name")
	rc.CheckEmpty(rc_name, "rc_name")

	err := models.Client.DeleteReplicationController(rc.namespace, rc_name)
	if err != nil {
		logs.Error("delete RC error, %v", err)
		rc.CustomAbort(http.StatusInternalServerError, "delete RC error")
	}
}
