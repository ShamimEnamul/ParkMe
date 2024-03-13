package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingLotController"],
        beego.ControllerComments{
            Method: "GetDetails",
            Router: "/details/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:ParkingSlotController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get1",
            Router: "/dd",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:UserController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "GetVehicleDetails",
            Router: "/details",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "Park",
            Router: "/park",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"] = append(beego.GlobalControllerRouter["ParkMe/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "UnPark",
            Router: "/unpark/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
