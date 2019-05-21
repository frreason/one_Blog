package controllers

import (
	"one_Blog/models"

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
	topics, err := models.GetAllTopic()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

}
