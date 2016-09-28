package routers

import (
	"thingscloud/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		/*beego.NSNamespace("/users",
			beego.NSRouter("/register",
				&controllers.UserController{},
				"post:Register"),
			beego.NSRouter("/login",
				&controllers.UserController{},
				"post:Login"),
			beego.NSRouter("/logout",
				&controllers.UserController{},
				"post:Logout"),
			beego.NSRouter("/passwd",
				&controllers.UserController{},
				"post:Passwd"),
			beego.NSRouter("/uploads",
				&controllers.UserController{},
				"post:Uploads"),
			beego.NSRouter("/downloads",
				&controllers.UserController{},
				"get:Downloads"),
		),*/
		beego.NSNamespace("/users",
			beego.NSRouter("/:id",
				&controllers.UserController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.UserController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/sessions",
			beego.NSRouter("/:id",
				&controllers.SessionController{},
				"delete:Delete"),
			beego.NSRouter("/",
				&controllers.SessionController{},
				"post:Post"),
		),
		beego.NSNamespace("/devices",
			beego.NSRouter("/:id",
				&controllers.DeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/gateways",
			beego.NSRouter("/:id",
				&controllers.GatewayController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.GatewayController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicetypes",
			beego.NSRouter("/:id",
				&controllers.DeviceTypeController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceTypeController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceattrs",
			beego.NSRouter("/:id",
				&controllers.DeviceAttrController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceAttrController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceattrinfos",
			beego.NSRouter("/:id",
				&controllers.DeviceAttrInfoController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceAttrInfoController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/userbindgateways",
			beego.NSRouter("/:id",
				&controllers.UserBindGatewayController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.UserBindGatewayController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/userbinddevices",
			beego.NSRouter("/:id",
				&controllers.UserBindDeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.UserBindDeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicectrls",
			beego.NSRouter("/:id",
				&controllers.DeviceCtrlController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceCtrlController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicereports",
			beego.NSRouter("/:id",
				&controllers.DeviceReportController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceReportController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicegroups",
			beego.NSRouter("/:id",
				&controllers.DeviceGroupController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceGroupController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicegroupbinddevices",
			beego.NSRouter("/:id",
				&controllers.DeviceGroupBindDeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceGroupBindDeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/rooms",
			beego.NSRouter("/:id",
				&controllers.RoomController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.RoomController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/roombinddevices",
			beego.NSRouter("/:id",
				&controllers.RoomBindDeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.RoomBindDeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/logicentitys",
			beego.NSRouter("/:id",
				&controllers.LogicEntityController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.LogicEntityController{},
				"get:GetAll;post:Post"),
		),
		/*beego.NSNamespace("/mqttservers",
			beego.NSRouter("/:id",
				&controllers.MQTTServerController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.MQTTServerController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceconnectstates",
			beego.NSRouter("/:id",
				&controllers.DeviceConnetcStateController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceConnectStateController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicedrivers",
			beego.NSRouter("/:id",
				&controllers.DeviceDriverController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceDriverController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/driverbinddevices",
			beego.NSRouter("/:id",
				&controllers.DriverBindDeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DriverBindDeviceController{},
				"get:GetAll;post:Post"),
		),*/
		beego.NSNamespace("/roles",
			beego.NSRouter("/:id",
				&controllers.RoleController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.RoleController{},
				"get:GetAll;post:Post"),
		),
	)
	beego.AddNamespace(ns)
}