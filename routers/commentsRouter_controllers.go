package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NamespaceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:namespace`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NodeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NodeController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:NodeController"],
		beego.ControllerComments{
			Method: "ListNodes",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/niyanchun/k8s-client/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
