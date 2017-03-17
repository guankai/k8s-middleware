package controllers

import (
	"github.com/astaxie/beego"
	"github.com/niyanchun/k8s-client/models"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// Operations about Namespace
type NamespaceController struct {
	beego.Controller
}

// @Title Get
// @Description get all namespaces
// @Success 200
// @router / [get]
func (ns *NamespaceController) Get() {
	namespaces, err := models.Client.ListNamespaces()
	if err != nil {
		logs.Error("list namespaces error, %v", err)
		ns.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	ns.Data["json"] = namespaces
	ns.ServeJSON()
}

// @Title Post
// @Description create a namespace
// @Param namespace formData string true "namespace name"
// @Success 200
// @router / [post]
func (ns *NamespaceController) Post() {
	name := ns.GetString("namespace")
	if len(name) == 0 {
		logs.Error("namespace is empty")
		ns.CustomAbort(http.StatusBadRequest, "namespace is empty")
	}

	namespace, err := models.Client.CreateNamespace(name)
	if err != nil {
		logs.Error("create namespace error, %v", err)
		ns.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	ns.Data["json"] = namespace
	ns.ServeJSON()
}

// @Title Delete
// @Description delete a namespace
// @Param namespace path string true "the namespace you want to delete"
// @router /:namespace [delete]
func (ns *NamespaceController) Delete() {
	name := ns.GetString(":namespace")
	if len(name) == 0 {
		logs.Error("namespace is empty")
		ns.CustomAbort(http.StatusBadRequest, "namespace is empty")
	}

	err := models.Client.DeleteNamespace(name)
	if err != nil {
		logs.Error("delete namespace error, %v", err)
		ns.CustomAbort(http.StatusInternalServerError, err.Error())
	}
}