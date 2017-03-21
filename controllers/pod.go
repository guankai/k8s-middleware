package controllers

import (
	"github.com/niyanchun/k8s-middleware/models"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// Operations about pod
type PodController struct {
	BaseController
	namespace string
}

func (p *PodController) Prepare() {
	namespace := p.GetString("namespace")
	p.CheckEmpty(namespace, "namespace")

	logs.Debug("namespace: %s", namespace)

	p.namespace = namespace
}


// TODO: 支持使用labels过滤查询到的Pod
// @Title Get
// @Description list all pods
// @Param namespace query string true "namespace"
// @router / [get]
func (p *PodController) List() {
	pod_list, err := models.Client.ListPods(p.namespace)
	if err != nil {
		logs.Error("list pods error: %s", err.Error())
		p.CustomAbort(http.StatusInternalServerError, "list pod error")
	}

	p.Data["json"] = pod_list
	p.ServeJSON()
}


// @Title Get
// @Description get pod details
// @Param namespace query string true "namespace"
// @Param pod_name path string true "pod name"
// @router /:pod_name [get]
func (p *PodController) Get() {
	pod_name := p.GetString(":pod_name")
	p.CheckEmpty(pod_name, "pod_name")
	logs.Debug("pod name: %s", pod_name)
	logs.Debug("namespace: %s", p.namespace)
	pod_detail, err := models.Client.GetPod(p.namespace, pod_name)
	if err != nil {
		logs.Error("get pod error, %s", err.Error())
		p.CustomAbort(http.StatusInternalServerError, "get pod error")
	}

	p.Data["json"] = pod_detail
	p.ServeJSON()
}


// TODO: UpdatePod()