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
	topics, err := models.GetAllTopic(false, true) //false表示不需要根据Views进行排序 后面的true表示根据更新的时间进行排序
	if err != nil {
		beego.Error(err)
	}

	categories, err := models.GetAllCategory(true)
	if err != nil {
		beego.Error(err)
	}
	if len(categories) > 5 {
		categories = categories[0:5]
	}

	viewsMaxTopic, err := models.GetAllTopic(true, false)
	if err != nil {
		beego.Error(err)
	}
	if len(viewsMaxTopic) > 5 {
		viewsMaxTopic = viewsMaxTopic[0:5]
	}

	lastestComments, err := models.GetAllComment(true)
	if err != nil {
		beego.Error(err)
	}
	if len(lastestComments) > 5 {
		lastestComments = lastestComments[0:5]
	}

	this.Data["Topics"] = topics
	this.Data["Categroies"] = categories
	this.Data["ViewsMaxTopic"] = viewsMaxTopic
	this.Data["LastestComments"] = lastestComments
}
