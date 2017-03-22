package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/niyanchun/k8s-middleware/models"
	"net/http"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	"encoding/json"
)

// Operations about service
type ServiceController struct {
	BaseController
	namespace string
}

func (s *ServiceController) Prepare() {
	namespace := s.GetString("namespace")
	method := s.Ctx.Input.Method()
	if method != http.MethodPost && method != http.MethodPut {
		s.CheckEmpty(namespace, "namespace")
	}

	s.namespace = namespace
}

// TODO: 增加根据labels过滤
// @Title List
// @Description list all services
// @Param namespace query string true "namespace"
// @router / [get]
func (s *ServiceController) List() {
	svc_list, err := models.Client.ListServices(s.namespace)
	if err != nil {
		logs.Error("list service error, %s", err.Error())
		s.CustomAbort(http.StatusInternalServerError, "list service error")
	}

	s.Data["json"] = svc_list
	s.ServeJSON()
}


// @Title Get
// @Description get service details
// @Param namespace query string true "namespace"
// @Param service_name path string true "service name"
// @router /:service_name [get]
func (s *ServiceController) Get() {
	service_name := s.GetString(":service_name")
	s.CheckEmpty(service_name, "service_name")

	svc_detail, err := models.Client.GetService(s.namespace, service_name)
	if err != nil {
		logs.Error("get service error, %v", err)
		s.CustomAbort(http.StatusInternalServerError, "get service error, " + err.Error())
	}

	s.Data["json"] = svc_detail
	s.ServeJSON()
}


// @Title Post
// @Description create a service
// @Param service body models.ServiceCopy true "service struct params"
// @router / [post]
func (s *ServiceController) Post() {
	var svc v1.Service
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &svc)
	if err != nil {
		logs.Error("unmarshal body error, %s", err.Error())
		s.CustomAbort(http.StatusInternalServerError, "unmarshal body error, " + err.Error() )
	}

	ret, err := models.Client.CreateService(svc.Namespace, &svc)
	s.CheckError(err, "create service", http.StatusInternalServerError)

	s.Data["json"] = ret
	s.ServeJSON()
}

// @Description  update a service
// @Param service body models.ServiceCopy true "service struct params"
// @router / [put]
func (s *ServiceController) Put() {
	var svc v1.Service
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &svc)
	if err != nil {
		logs.Error("unmarshal body error, %s", err.Error())
		s.CustomAbort(http.StatusInternalServerError, "unmarshal body error, " + err.Error() )
	}

	ret, err := models.Client.UpdateService(svc.Namespace, &svc)
	s.CheckError(err, "update serivce error", http.StatusInternalServerError)

	s.Data["json"] = ret
	s.ServeJSON()
}

// @Title Delete
// @Description delete a service
// @Param namespace query string true "namespace"
// @Param svc_name path string true "service name"
// @router /:svc_name [delete]
func (s *ServiceController) Delete() {
	svc_name := s.GetString(":svc_name")
	s.CheckEmpty(svc_name, "svc_name")

	err := models.Client.DeleteService(s.namespace, svc_name)
	if err != nil {
		logs.Error("delete service error, %s", err.Error())
		s.CustomAbort(http.StatusInternalServerError, "delete service error, " + err.Error())
	}
}