package routers

import (
	"CMSTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/Admin/UserInfo/Index", &controllers.UserInfoController{}, "get:Index")
	beego.Router("/Admin/UserInfo/GetUserInfo", &controllers.UserInfoController{}, "post:GetUserInfo")
	beego.Router("/Admin/UserInfo/AddUser", &controllers.UserInfoController{}, "post:AddUser")
	beego.Router("/Admin/UserInfo/DeleteUser", &controllers.UserInfoController{}, "post:DeleteUser")
	beego.Router("/Admin/UserInfo/GetSingleUserInfo", &controllers.UserInfoController{}, "get:GetSingleUserInfo")
	beego.Router("/Admin/UserInfo/EditUserInfo", &controllers.UserInfoController{}, "post:EditUserInfo")
	beego.Router("/Admin/UserInfo/ShowSetUserRole", &controllers.UserInfoController{}, "get:ShowSetUserRole")
	beego.Router("/Admin/UserInfo/SetUserRole", &controllers.UserInfoController{}, "post:SetUserRole")
	//------------------------------角色管理----------------------------------
	beego.Router("/Admin/RoleInfo/Index", &controllers.RoleInfoController{}, "get:Index")
	beego.Router("/Admin/RoleInfo/ShowAddRole", &controllers.RoleInfoController{}, "get:ShowAddRole")
	beego.Router("/Admin/RoleInfo/AddRole", &controllers.RoleInfoController{}, "post:AddRole")
	beego.Router("/Admin/RoleInfo/GetRoleInfo", &controllers.RoleInfoController{}, "post:GetRoleInfo")
	beego.Router("/Admin/RoleInfo/ShowRoleAction", &controllers.RoleInfoController{}, "get:ShowRoleAction")
	beego.Router("/Admin/RoleInfo/SetRoleAction", &controllers.RoleInfoController{}, "post:SetRoleAction")

	//---------------------------权限管理-------------------------------------
	beego.Router("/Admin/ActionInfo/Index", &controllers.ActionInfoController{}, "get:Index")
	beego.Router("/Admin/ActionInfo/FileUp", &controllers.ActionInfoController{}, "post:FileUp")
	beego.Router("/Admin/ActionInfo/AddAction", &controllers.ActionInfoController{}, "post:AddAction")
	beego.Router("/Admin/ActionInfo/GetActionInfo", &controllers.ActionInfoController{}, "post:GetActionInfo")
}
