// @APIVersion 1.0.0
// @Title Innovation Platform Kubernetes Middleware API
// @Description just as the title describes, no more
// @Contact ych.ni@hnair.com
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
		beego.NSNamespace("/app",
			beego.NSInclude(
				&controllers.AppController{},
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
		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.ServiceController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
