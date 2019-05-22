package controllers

import (
	"one_Blog/models"
	"strconv"

	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) AddComment() {

	tid := this.Ctx.Input.Param(":id")
	realTid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	nickName := this.Input().Get("nickName")
	comment := this.Input().Get("comment")
	err = models.AddComment(realTid, nickName, comment)

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 302)
	return
}

func (this *CommentController) DelComment() {

}

func (this *CommentController) GetComment() {

}
