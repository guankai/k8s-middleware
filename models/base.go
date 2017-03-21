package models

import (
	"github.com/astaxie/beego"
	k8s_client "bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"
	"github.com/astaxie/beego/logs"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
)

var Client k8s_client.Client

func init() {
	k8s_url := beego.AppConfig.String("kubernetes_url")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	logs.Debug("Kubernetes server: %s, username: %s", k8s_url, username)
	Client = k8s_client.NewClient(k8s_url, username, password)
}

// these ***copy struct just for shut up the bee gendoc, no other use
type ReplicationControllerCopy struct {
	v1.ReplicationController
}

type ServiceCopy struct {
	v1.Service
}