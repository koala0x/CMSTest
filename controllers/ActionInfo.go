package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"time"
	"os"
	"github.com/astaxie/beego/orm"
	"CMSTest/models"
)

type ActionInfoController struct {
	beego.Controller
}

func (this *ActionInfoController) Index() {
	this.TplName = "ActionInfo/Index.html"
}

func (this *ActionInfoController) FileUp() {
	file, header, e := this.GetFile("fileUp")
	defer file.Close()
	if e != nil {
		this.Data["json"] = map[string]interface{}{"flag": "no", "msg": "文件上传失败!!"}
		beego.Error("文件上传失败", e)
	} else {
		fileExit := path.Ext(header.Filename)
		if fileExit == ".jpg" || fileExit == ".png" {
			if header.Size < 50000 {
				dir := "./static/fileUp/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + "/"
				_, err1 := os.Stat(dir) //判断文件是否存在
				if err1 != nil {
					os.MkdirAll(dir, os.ModePerm)
				}
				newFileName := strconv.Itoa(time.Now().Year()) + time.Now().Month().String() + strconv.Itoa(time.Now().Day()) + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Nanosecond())
				//完整的上传路径
				fullDir := dir + newFileName + fileExit
				err2 := this.SaveToFile("fileUp", fullDir)
				if err2 == nil {
					this.Data["json"] = map[string]interface{}{"flag": "ok", "msg": fullDir}
				} else {
					this.Data["json"] = map[string]interface{}{"flag": "no", "msg": "文件保存失败!!!"}
					beego.Error("文件保存失败!!!", err2)
				}
			} else {
				this.Data["json"] = map[string]interface{}{"flag": "no", "msg": "文件太大!!!"}
			}
		} else {
			this.Data["json"] = map[string]interface{}{"flag": "no", "msg": "文件类型错误"}
		}
		this.ServeJSON()
	}
}
func (this *ActionInfoController) GetActionInfo() {
	pageIndex, _ := this.GetInt("page")
	pageSize, _ := this.GetInt("rows")
	start := (pageIndex - 1) * pageSize
	newOrm := orm.NewOrm()
	var actionInfo []models.ActionInfo
	newOrm.QueryTable("action_info").Filter("del_flag", 0).OrderBy("id").Limit(pageSize, start).All(&actionInfo)
	beego.Info("获取的信息", actionInfo)
	count, _ := newOrm.QueryTable("action_info").Filter("del_flag", 0).Count()
	this.Data["json"] = map[string]interface{}{"rows": actionInfo, "total": count}
	this.ServeJSON()
}

func (this *ActionInfoController) AddAction() {
	actionInfo := models.ActionInfo{}
	actionInfo.ActionTypeEnum, _ = this.GetInt("ActionTypeEnum")
	actionInfo.MenuIcon = this.GetString("MenuIcon")
	actionInfo.Url = this.GetString("Url")
	actionInfo.ActionInfoName = this.GetString("ActionInfoName")
	actionInfo.IconWidth = 0
	actionInfo.IconHeight = 0
	actionInfo.HttpMethod = this.GetString("HttpMethod")
	actionInfo.Remark = this.GetString("Remark")
	actionInfo.DelFlag = 0
	actionInfo.AddDate = time.Now()
	actionInfo.ModifDate = time.Now()
	//beego.Info("获取到的对象是什么样子的", actionInfo)
	newOrm := orm.NewOrm()
	_, e := newOrm.Insert(&actionInfo)
	if e == nil {
		this.Data["json"] = map[string]interface{}{"flag": "ok"}
	} else {
		this.Data["json"] = map[string]interface{}{"flag": "fail"}
	}
	this.ServeJSON()
}
