package models

import (
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/api/v1"
	"bitbucket.org/amdatulabs/amdatu-kubernetes-go/client"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
)

const (
	AppReplicas int = 1

	AppStatusPending v1.PodPhase = v1.PodPending
	AppStatusRunning v1.PodPhase = v1.PodRunning
	AppStatusStopped v1.PodPhase = PodStatusStopped
	AppStatusFailed  v1.PodPhase = v1.PodFailed
	AppStatusUnkonwn v1.PodPhase = v1.PodUnknown
)

type AppInfo struct {
	Name   string      `json:"name"`
	Status v1.PodPhase `json:"status"`
}

type AppList struct {
	Apps []AppInfo `json:"apps"`
}

type AppDetails struct {
	Service               v1.Service    `json:"service"`
	ReplicationController v1.ReplicationController  `json:"replication_controller"`
	Pod                   v1.Pod    `json:"pod"`
}

// App is stopped if RC's replicas is 0
func IsAppStoppped(namespace string, labels client.Labels) (error, bool) {
	rcs, err := Client.ListReplicationControllersWithLabel(namespace, labels)
	if err != nil {
		logs.Error("ListReplicationControllersWithLabel error, %s", err.Error())
		return err, true
	}

	if len(rcs.Items) != 1 {
		s := fmt.Sprintf("ReplicationControllers with label %v is not unique", labels)
		logs.Error(s)
		return errors.New(s), true
	}

	return nil, *(rcs.Items[0].Spec.Replicas) == int32(0)
}
