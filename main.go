package main

import (
	_ "CMSTest/routers"
	"github.com/astaxie/beego"
	_ "CMSTest/models"
	"github.com/astaxie/beego/logs"
	"CMSTest/models"
)

//beego.Emergency("this is emergency")
//beego.Alert("this is alert") 1
//beego.Critical("this is critical") 2
//beego.Error("this is error") 3
//beego.Warning("this is warning") 4
//beego.Notice("this is notice") 5
//beego.Informational("this is informational") 6
//beego.Debug("this is debug") 7

func init() {
	beego.SetLogFuncCall(true)
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterFile, `{"filename":"CMSTest.log","level":3}`)
}

func CheckUserRole(userInfo models.UserInfo, roleId int) bool {
	check := false
	for i := 0; i < len(userInfo.Roles); i++ {
		if userInfo.Roles[i].Id == roleId {
			check = true
			break
		}
	}
	return check
}

func ShowActionInfo(info []*models.ActionInfo, roleId int) (b bool) {
	b = false
	for i := 0; i < len(info); i++ {
		if info[i].Id == roleId {
			b = true
			break
		}
	}
	return
}

func main() {
	beego.AddFuncMap("checkUserRole", CheckUserRole) //模板函数的定义
	beego.AddFuncMap("showActionInfo", ShowActionInfo)
	beego.Run()
}
