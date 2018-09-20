package controllers

import (
	"github.com/astaxie/beego"
	"itcastCms/models"
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

type UserInfoController struct {
	beego.Controller
}

func (this *UserInfoController) Index() {
	this.TplName = "UserInfo/Index.html"
}

func (this *UserInfoController) GetUserInfo() {
	pageIndex, _ := this.GetInt("page") //当前页码 easyui表单提交过来的数据
	pageSize, _ := this.GetInt("rows")  //每页展示的条数
	users := new([]models.UserInfo)
	//var users []models.UserInfo
	newOrm := orm.NewOrm()
	newOrm.QueryTable("user_info").Filter("del_flag", 0).OrderBy("id").Limit(pageSize, (pageIndex-1)*pageSize).All(users)
	count, _ := newOrm.QueryTable("user_info").Filter("del_flag", 0).Filter("del_flag", 0).Count()
	this.Data["json"] = map[string]interface{}{"rows": users, "total": count}
	this.ServeJSON()
}

func (this *UserInfoController) AddUser() {
	userInfo := new(models.UserInfo)
	(*userInfo).UserName = this.GetString("userName")
	(*userInfo).UserPwd = this.GetString("userPwd")
	(*userInfo).Remark = this.GetString("userRemark")
	(*userInfo).AddDate = time.Now()
	(*userInfo).ModifDate = time.Now()
	(*userInfo).DelFlag = 0
	newOrm := orm.NewOrm()
	_, e := newOrm.Insert(userInfo) //这里一定是地址值
	if e == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "no"}
	}
	this.ServeJSON()
}
func (this *UserInfoController) DeleteUser() {
	ids := this.GetString("strId")
	strIds := strings.Split(ids, ",")
	orm := orm.NewOrm()
	info := models.UserInfo{}
	for i := 0; i < len(strIds); i++ {
		id, _ := strconv.Atoi(strIds[i])
		info.Id = id
		err := orm.Read(&info)
		if err == nil {
			info.DelFlag = 1
		}
		_, e := orm.Update(&info)
		//_, e := orm.Delete(&info)
		if e == nil {
			this.Data["json"] = map[string]interface{}{"flag": "ok"}
		} else {
			this.Data["json"] = map[string]interface{}{"flag": "fail"}
		}
		this.ServeJSON()
	}
}

func (this *UserInfoController) GetSingleUserInfo() {
	editId, _ := this.GetInt("editId")
	newOrm := orm.NewOrm()
	info := new(models.UserInfo)
	(*info).Id = editId
	err := newOrm.Read(info) //等会儿加上日志
	if err != nil {
		beego.Info("要加入日志系统", err)
		this.Data["json"] = map[string]interface{}{"flag": "fail", "msg": err}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "ok", "data": &info, "msg": "数据获取成功"}
	}
	this.ServeJSON()
}

func (this *UserInfoController) EditUserInfo() {
	userName := this.GetString("editUserName")
	userPwd := this.GetString("editUserPwd")
	userRemark := this.GetString("editUserRemark")
	editId, _ := this.GetInt("Id")
	userInfo := new(models.UserInfo)
	(*userInfo).Id = editId
	newOrm := orm.NewOrm()
	err := newOrm.Read(userInfo)
	if err != nil {
		beego.Info("报错了", err)
		return
	}
	(*userInfo).UserName = userName
	(*userInfo).Remark = userRemark
	(*userInfo).UserPwd = userPwd
	(*userInfo).ModifDate = time.Now()
	_, err2 := newOrm.Update(userInfo)
	if err2 == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok", "msg": "修改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "fail", "msg": "修改失败"}
	}
	this.ServeJSON()
}
