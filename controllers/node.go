package controllers

import (
	"github.com/astaxie/beego"
	"github.com/niyanchun/k8s-client/models"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// Operations about Nodes
type NodeController struct {
	beego.Controller
}

// @Title Get
// @Description get node information by name
// @Param name path string true "node name"
// @Success 200
// @router /:name [get]
func (n *NodeController) Get() {
	name := n.GetString(":name")
	if len(name) == 0 {
		logs.Error("name is empty")
		n.CustomAbort(http.StatusBadRequest, "name is empty")
	}

	logs.Debug("node name: %s", name)

	node, err:= models.Client.GetNode(name)
	if err != nil {
		logs.Error("get node error %v", err)
		n.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	n.Data["json"] = node
	n.ServeJSON()
}


// @Title List
// @Description List all nodes
// @Success 200
// @router / [get]
func (n *NodeController) ListNodes() {
	nodes, err := models.Client.ListNodes()
	if err != nil {
		logs.Error("List nodes error %v", err)
		n.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	logs.Debug("nodes: %v", nodes)

	n.Data["json"] = nodes
	n.ServeJSON()
}