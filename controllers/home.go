package controllers

import (
	"github.com/astaxie/beego"
)

//首页的控制器

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true

	isLogin := checkAccount(this.Ctx)
	if isLogin {
		this.Data["IsLogin"] = true
	}

}
