package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/niyanchun/k8s-middleware/models"
	"io/ioutil"
)

// Operations about Apps
type AppController struct {
	BaseController
	namespace string
}

func (app *AppController) Prepare() {
	namespace := app.GetString("namespace")

	method := app.Ctx.Input.Method()
	if method != http.MethodPost && method != http.MethodPut {
		app.CheckEmpty(namespace, "namespace")
	}

	logs.Debug("namespace: %s", namespace)
	app.namespace = namespace
}

// @Description list user's all apps
// @Param namespace query string true "namespace"
// @router / [get]
func (app *AppController) List() {
	svc_list, err := models.Client.ListServices(app.namespace)
	app.CheckError(err, "[app.List] list service error", http.StatusInternalServerError)

	var apps []models.AppInfo
	var a models.AppInfo
	for _, svc := range svc_list.Items {
		a.Name = svc.Name

		err, status := models.GetPodsStatus(svc.Namespace, svc.Spec.Selector)
		if err != nil {
			s := fmt.Sprintf("get service %s's status error, %s", svc.Name, err.Error())
			logs.Error(s)
			a.Status = models.AppStatusUnkonwn
		} else {
			a.Status = status
		}

		apps = append(apps, a)
	}

	app.Data["json"] = apps
	app.ServeJSON()
}

// @Description get app detail: service + RC + Pod
// @Param namespace query string true "namespace"
// @Param app_name path string true "app name"
// @router /:app_name [get]
func (app *AppController) Get() {
	app_name := app.GetString(":app_name")
	app.CheckEmpty(app_name, "app_name")

	logs.Debug("app_name: %s", app_name)

	var app_details models.AppDetails

	service, err := models.Client.GetService(app.namespace, app_name)
	app.CheckError(err, "get service error", http.StatusInternalServerError)
	app_details.Service = *service

	rc, err := models.Client.GetReplicationController(app.namespace, app_name)
	app.CheckError(err, "get rc error", http.StatusInternalServerError)
	app_details.ReplicationController = *rc

	labels := service.Spec.Selector
	logs.Debug("labels: %v", labels)
	pods, err := models.Client.ListPodsWithLabel(app.namespace, labels)
	app.CheckError(err, "get pod error", http.StatusInternalServerError)

	pod_num := len(pods.Items)
	if pod_num == 1 {
		app_details.Pod = pods.Items[0]
	} else if pod_num == 0 {
		// no pod is running, just set App status as "Stopped"
		app_details.Pod.Status.Phase = models.AppStatusStopped
	} else {
		s := fmt.Sprintf("pods with label %v is not unique", labels)
		logs.Error(s)
		app.CustomAbort(http.StatusInternalServerError, s)
	}

	app.Data["json"] = app_details
	app.ServeJSON()
}

// @Description toggle app status(start, stop). 0: stopped, 1: running
// @Param namespace query string true "namespace"
// @Param app_name path string true "app name"
// @router /:app_name [post]
func (app *AppController) Toggle() {
	app_name := app.GetString(":app_name")
	app.CheckEmpty(app_name, "app_name")

	logs.Debug("app: %s, ns: %s", app_name, app.namespace)
	svc, err := models.Client.GetService(app.namespace, app_name)
	app.CheckError(err, "get service error", http.StatusInternalServerError)

	labels := svc.Spec.Selector

	rcs, err := models.Client.ListReplicationControllersWithLabel(app.namespace, labels)
	app.CheckError(err, fmt.Sprintf("list rc with label %s error", labels), http.StatusInternalServerError)

	num := len(rcs.Items)
	if num != 0 {

	}

	rc := rcs.Items[0]
	replicas := *rc.Spec.Replicas

	if replicas != 0 {
		*rc.Spec.Replicas = 0
	} else {
		*rc.Spec.Replicas = 1
	}

	ret, err := models.Client.UpdateReplicationController(app.namespace, &rc)
	app.CheckError(err, "update rc error", http.StatusInternalServerError)

	app.Data["json"] = ret.Spec.Replicas
	app.ServeJSON()
}

// @Description delete an App
// @Param namespace query string true "namespace"
// @Param app_name path string true "app name"
// @router / [delete]
func (app *AppController) Delete() {
	app_name := app.GetString(":app_name")
	app.CheckEmpty(app_name, "app_name")

	err1 := models.Client.DeleteReplicationController(app.namespace, app_name)
	if err1 != nil {
		logs.Error("delete RC error: " + err1.Error())
	}

	err2 := models.Client.DeleteService(app.namespace, app_name)
	if err2 != nil {
		logs.Error("delete service error:" + err2.Error())
	}

	if err1 != nil || err2 != nil {
		app.CustomAbort(http.StatusInternalServerError, "delete failed")
	}
}

// @Description get app log
// @Param namespace query string true "namespace"
// @Param app_name path string true "app name"
// @router /logs/:app_name [get]
func (app *AppController) Log() {
	type PodLogs struct {
		PodName string `json:"pod_name"`
		Log     string `json:"log"`
	}

	app_name := app.GetString(":app_name")
	app.CheckEmpty(app_name, "app_name")

	svc, err := models.Client.GetService(app.namespace, app_name)
	app.CheckError(err, "get service error", http.StatusInternalServerError)

	labels := svc.Spec.Selector
	pods, err := models.Client.ListPodsWithLabel(app.namespace, labels)

	var pod_logs []PodLogs
	for _, pod := range pods.Items {
		var log PodLogs
		log.PodName = pod.Name
		url := models.Client.Url + "/api/v1/namespaces/" + app.namespace + "/pods/" + pod.Name + "/log"
		resp, err := http.Get(url)
		app.CheckError(err, "get pod log error", http.StatusInternalServerError)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		log.Log = string(body)
		pod_logs = append(pod_logs, log)
	}

	app.Data["json"] = pod_logs
	app.ServeJSON()
}
