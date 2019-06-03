package controllers

import (
	"one_Blog/models"
	"strconv"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

//分类的页面应该根据访客和管理员呈现不一样的界面
func (this *CategoryController) Get() {

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	op := this.Input().Get("op") //说明是哪个操作

	switch op {
	case "del":
		if !checkAccount(this.Ctx) { //管理员权限检查
			this.Redirect("/login", 302)
			return
		}
		cid, err := strconv.ParseInt(this.Input().Get("id"), 10, 64) //该分类的Id
		if len(this.Input().Get("id")) == 0 {
			break
		}
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteCategory(cid)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "add":
		if !checkAccount(this.Ctx) { //管理员权限检查
			this.Redirect("/login", 302)
			return
		}
		tName := this.Input().Get("categoryName")
		if len(tName) == 0 {
			break
		}
		err := models.AddCategory(tName)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}

	this.TplName = "category.html"

	var err error
	this.Data["Categories"], err = models.GetAllCategory(false)
	if err != nil {
		beego.Error(err)
	}

}
