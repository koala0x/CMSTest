package routers

import (
	"itcastCms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/Admin/UserInfo/Index",&controllers.UserInfoController{},"get:Index")
    beego.Router("/Admin/UserInfo/GetUserInfo",&controllers.UserInfoController{},"post:GetUserInfo")
    beego.Router("/Admin/UserInfo/AddUser",&controllers.UserInfoController{},"post:AddUser")
    beego.Router("/Admin/UserInfo/DeleteUser",&controllers.UserInfoController{},"post:DeleteUser")
    beego.Router("/Admin/UserInfo/GetSingleUserInfo",&controllers.UserInfoController{},"get:GetSingleUserInfo")
    beego.Router("/Admin/UserInfo/EditUserInfo",&controllers.UserInfoController{},"post:EditUserInfo")
}
