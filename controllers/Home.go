package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) ShowIndex() {
	this.TplName = "Home/ShowIndex.html"
}

func (this *HomeController) Index() {
	this.TplName = "Home/Index.html"
}
