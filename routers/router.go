// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/niyanchun/k8s-middleware/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/node",
			beego.NSInclude(
				&controllers.NodeController{},
			),
		),
		beego.NSNamespace("/namespace",
			beego.NSInclude(
				&controllers.NamespaceController{},
			),
		),
		beego.NSNamespace("/rc",
			beego.NSInclude(
				&controllers.RCController{},
			),
		),
		beego.NSNamespace("/pod",
			beego.NSInclude(
				&controllers.PodController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
