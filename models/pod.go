package models

import (
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"
	"github.com/astaxie/beego/logs"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	"fmt"
	"github.com/pkg/errors"
)

const (
	PodStatusStopped v1.PodPhase =  "Stopped"  // 已停止     v1.PodSucceeded
)


func GetPodsStatus(namespace string, labels client.Labels) (error, v1.PodPhase) {
	pods, err := Client.ListPodsWithLabel(namespace, labels)
	if err != nil {
		logs.Error("GetPodsStatus error, %s", err.Error())
		return err, v1.PodUnknown
	}

	pod_num := len(pods.Items)
	if pod_num ==  1 {
		return nil, pods.Items[0].Status.Phase
	} else if pod_num == 0 {
		return nil, PodStatusStopped
	}

	s := fmt.Sprintf("pod with label %v is not unique", labels)
	logs.Error(s)
	return errors.New(s), v1.PodUnknown
}