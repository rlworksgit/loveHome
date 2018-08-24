package routers

import (
	"loveHome/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//地域信息
	beego.Router("/api/v1.0/areas", &controllers.AreaControllers{}, "get:GetAreaInfo")
	//请求session
	beego.Router("/api/v1.0/session", &controllers.SessionControllers{}, "get:GetSessionInfo;delete:DeleteSession")
	//房屋首页信息
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexControllers{}, "get:GetHouseIndex")
	//注册
	beego.Router("/api/v1.0/users", &controllers.UserControllers{}, "post:Reg")
	//登陆
	beego.Router("/api/v1.0/sessions", &controllers.UserControllers{}, "post:Login")
	//登入请求
	//beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:Login")

}
