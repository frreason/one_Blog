package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//登录页的控制器

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	isExit := (this.Input().Get("exit") == "true")
	log.Println("handler")
	if isExit {
		log.Println("exit")
		this.Ctx.SetCookie("userName", "", -1, "/") //删除cookies
		this.Ctx.SetCookie("pwd", "", -1, "/")      //同上
		this.Redirect("/", 301)
	} else {
		this.TplName = "login.html"
	}
}

func (this *LoginController) Post() {
	userName := this.Input().Get("userName")
	pwd := this.Input().Get("pwd")
	autoLogin := (this.Input().Get("autoLogin") == "on")

	if userName == beego.AppConfig.String("userName") && pwd == beego.AppConfig.String("pwd") {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("userName", userName, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool { //新版不能写beego.Context
	ck, err := ctx.Request.Cookie("userName")
	if err != nil {
		return false
	}
	userName := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	if userName == beego.AppConfig.String("userName") && pwd == beego.AppConfig.String("pwd") {
		return true
	}
	return false
}
