package controllers

import (
	"github.com/niyanchun/k8s-middleware/models"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// Operations about Nodes
type NodeController struct {
	BaseController
}


// @Title List
// @Description List all nodes
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

// @Title Get
// @Description get node information by name
// @Param name path string true "node name"
// @router /:name [get]
func (n *NodeController) Get() {
	name := n.GetString(":name")
	n.CheckEmpty(name, "name")

	logs.Debug("node name: %s", name)
	node, err:= models.Client.GetNode(name)
	if err != nil {
		logs.Error("get node error %v", err)
		n.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	n.Data["json"] = node
	n.ServeJSON()
}

