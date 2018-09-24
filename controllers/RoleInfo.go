package controllers

import (
	"github.com/astaxie/beego"
	"CMSTest/models"
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

type RoleInfoController struct {
	beego.Controller
}

func (this *RoleInfoController) Index() {
	this.TplName = "RoleInfo/Index.html"
}

func (this *RoleInfoController) ShowAddRole() {
	this.TplName = "RoleInfo/ShowAddRole.html"
}

func (this *RoleInfoController) AddRole() {
	var roleInfo = models.RoleInfo{}
	roleInfo.RoleName = this.GetString("roleName")
	roleInfo.Remark = this.GetString("roleRemark")
	roleInfo.DelFlag = 0
	roleInfo.AddDate = time.Now()
	roleInfo.ModifDate = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(&roleInfo)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "no"}
	}
	this.ServeJSON()
}

//获取角色信息
func (this *RoleInfoController) GetRoleInfo() {
	pageIndex, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("rows")
	start := (pageIndex - 1) * pageSize
	o := orm.NewOrm()
	var roles []models.RoleInfo
	o.QueryTable("role_info").Filter("del_flag", 0).OrderBy("Id").Limit(pageSize, start).All(&roles)
	count, _ := o.QueryTable("role_info").Filter("del_flag", 0).Count()
	this.Data["json"] = map[string]interface{}{"rows": roles, "total": count}
	this.ServeJSON()
}

func (this *RoleInfoController) ShowRoleAction() {
	roleId := this.GetString("roleId")
	newOrm := orm.NewOrm()
	roleInfo := models.RoleInfo{}
	newOrm.QueryTable("role_info").Filter("id", roleId).One(&roleInfo)
	newOrm.LoadRelated(&roleInfo, "Actions")
	var actionInfoAll []models.ActionInfo
	newOrm.QueryTable("action_info").Filter("del_flag", 0).All(&actionInfoAll)
	this.Data["chcekAction"] = roleInfo.Actions
	this.Data["allAction"] = actionInfoAll
	this.Data["roleId"] = roleInfo.Id
	this.TplName = "RoleInfo/ShowSetRoleAction.html"
}

func (this *RoleInfoController) SetRoleAction() {
	roleId, _ := this.GetInt("roleId")
	allKeys := this.Ctx.Request.PostForm
	var listId []int
	for key, _ := range allKeys {
		if strings.Contains(key, "cba_") {
			id := strings.Replace(key, "cba_", "", -1)
			idInt, _ := strconv.Atoi(id)
			listId = append(listId, idInt)
		}
	}
	newOrm := orm.NewOrm()
	roleInfo := models.RoleInfo{}
	newOrm.QueryTable("role_info").Filter("id", roleId).One(&roleInfo)
	newOrm.LoadRelated(&roleInfo, "Actions")
	queryM2M := newOrm.QueryM2M(&roleInfo, "Actions")
	for _, action := range roleInfo.Actions {
		queryM2M.Remove(action)
	}
	info := models.ActionInfo{}
	for i := 0; i < len(listId); i++ {
		newOrm.QueryTable("action_info").Filter("id", listId[i]).One(&info)
		queryM2M.Add(info)
	}
	this.Data["json"] = map[string]interface{}{"flag": "ok"}
	this.ServeJSON()
}
