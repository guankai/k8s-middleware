package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:AppController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:AppController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:app_name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:namespace`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NodeController"],
		beego.ControllerComments{
			Method: "ListNodes",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:PodController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:PodController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:PodController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:PodController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:pod_name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:rc_name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:RCController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:rc_name`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:service_name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:svc_name`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-middleware/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
