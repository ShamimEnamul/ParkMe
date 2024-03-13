// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ParkMe/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/parking-lot",
			beego.NSInclude(
				&controllers.ParkingLotController{},
			),
		),
		beego.NSNamespace("/parking-slot",
			beego.NSInclude(
				&controllers.ParkingSlotController{},
			),
		),
		/*	beego.NSNamespace("/reservation",
			beego.NSInclude(
				&controllers.ReservationController{},
			),
		),*/
		beego.NSNamespace("/vehicle",
			beego.NSInclude(
				&controllers.VehicleController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
