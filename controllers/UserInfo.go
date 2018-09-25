package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
	"strconv"
	"CMSTest/models"
)

type UserInfoController struct {
	beego.Controller
}

//记录分页信息
type RecordPagingData struct {
	PageIndex  int
	PageSize   int
	Name       string
	Remark     string
	TotalCount int
}

func (this *UserInfoController) Index() {
	this.TplName = "UserInfo/Index.html"
}

func (this *UserInfoController) GetUserInfo() {
	pageIndex, _ := this.GetInt("page") //当前页码 easyui表单提交过来的数据
	pageSize, _ := this.GetInt("rows")  //每页展示的条数
	name := this.GetString("name")
	remark := this.GetString("remark")
	recordPaging := RecordPagingData{PageIndex: pageIndex, PageSize: pageSize, Name: name, Remark: remark}
	userData := recordPaging.SearchUserData()
	this.Data["json"] = map[string]interface{}{"rows": &userData, "total": recordPaging.TotalCount}
	this.ServeJSON()

}
func (this *RecordPagingData) SearchUserData() (*[]models.UserInfo) {
	newOrm := orm.NewOrm()
	queryTable := newOrm.QueryTable("user_info")
	if this.Name != "" {
		queryTable = queryTable.Filter("user_name__icontains", this.Name)
	}
	if this.Remark != "" {
		queryTable = queryTable.Filter("remark__icontains", this.Remark)
	}
	queryTable = queryTable.Filter("del_flag", 0)
	count, _ := queryTable.Count()
	this.TotalCount = int(count)
	start := (this.PageIndex - 1) * this.PageSize
	userInfo := new([]models.UserInfo)
	queryTable.OrderBy("Id").Limit(this.PageSize, start).All(userInfo)
	return userInfo
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

func (this *UserInfoController) ShowSetUserRole() {
	userId, _ := this.GetInt("userId") //得到被修改角色信息的人员ID
	newOrm := orm.NewOrm()
	var userInfo models.UserInfo
	newOrm.QueryTable("user_info").Filter("id", userId).One(&userInfo)
	newOrm.LoadRelated(&userInfo, "Roles")
	var allRoles []models.RoleInfo
	newOrm.QueryTable("role_info").Filter("del_flag", 0).All(&allRoles)
	this.Data["allRoles"] = allRoles //全部权限
	this.Data["userInfo"] = userInfo //选择的权限
	this.TplName = "UserInfo/ShowSetUserRole.html"
}

//有两种方式可以获取form的值
//roleInfo.RoleName=this.GetString("roleName")
//roleInfo.Remark=this.GetString("roleRemark")
//直接通过getString取值
//另外一种如下:
func (this *UserInfoController) SetUserRole() {
	allKeys := this.Ctx.Request.PostForm
	beego.Info("复选框信息", allKeys)
	var list []int
	for key, _ := range allKeys {
		if strings.Contains(key, "cba_") { //只有被选中的复选框才会提交给客户端~
			id := strings.Replace(key, "cba_", "", -1)
			roleID, _ := strconv.Atoi(id)
			list = append(list, roleID)
		}
	}
	userId, _ := this.GetInt("userId")
	beego.Info("隐藏域选择的用户名", userId)
	userInfo := models.UserInfo{}
	newOrm := orm.NewOrm()
	newOrm.QueryTable("user_info").Filter("id", userId).One(&userInfo)
	newOrm.LoadRelated(&userInfo, "Roles")
	queryM2M := newOrm.QueryM2M(&userInfo, "Roles")
	for _, role := range userInfo.Roles {
		queryM2M.Remove(role)
	}
	roleInfo := models.RoleInfo{}
	for i := 0; i < len(list); i++ {
		newOrm.QueryTable("role_info").Filter("id", list[i]).One(&roleInfo)
		queryM2M.Add(roleInfo)
	}
	this.Data["json"] = map[string]interface{}{"flag": "ok"}
	this.ServeJSON()
}

/**
展示用户所拥有的权限和能否使用的权限
 */
func (this *UserInfoController) ShowSetUserAction() {
	userId, _ := this.GetInt("userId")
	newOrm := orm.NewOrm()
	userInfo := models.UserInfo{}
	newOrm.QueryTable("user_info").Filter("id", userId).One(&userInfo)
	//查询用户已经有的权限
	userActions := []models.UserAction{}
	newOrm.QueryTable("user_action").Filter("users_id", userId).All(&userActions)
	//查询所有权限
	actionInfos := []models.ActionInfo{}
	newOrm.QueryTable("action_info").Filter("del_flag", 0).All(&actionInfos)
	this.Data["userActions"] = userActions //用户已有权限集合
	this.Data["actionInfos"] = actionInfos //系统的全部权限
	this.Data["userInfo"] = userInfo
	this.TplName = "UserInfo/ShowSetUserAction.html"
}

func (this *UserInfoController) DeleteUserAction() {
	//删除权限
}
func (this *UserInfoController) SetUserAction() {
	//设置权限
}
