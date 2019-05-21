package controllers

import (
	"log"
	"one_Blog/models"
	"strconv"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic.html"

	topics, err := models.GetAllTopic()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

}

func (this *TopicController) Add() {

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topicAdd.html"
}

func (this *TopicController) DelTopic() {

	id := this.Ctx.Input.Param(":id")
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// log.Println(tid)
	err = models.DeleteTopic(tid)

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
	return
}

func (this *TopicController) AddTopic() {

	if !checkAccount(this.Ctx) { //管理员权限检查
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	err := models.AddTopic(title, content, category)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}

func (this *TopicController) UpTopic() {

	if !checkAccount(this.Ctx) { //管理员权限检查
		this.Redirect("/login", 302)
		return
	}
	id := this.Ctx.Input.Param(":id")
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	this.TplName = "topicUpdate.html"

	//先取出原来的内容
	var oneTopic *models.Topic
	oneTopic, err = models.GetTopic(tid)

	if err != nil {
		beego.Error(err)
	}
	this.Data["Id"] = id
	this.Data["Category"] = oneTopic.Category
	this.Data["Title"] = oneTopic.Title
	this.Data["Content"] = oneTopic.Content
	return
}

func (this *TopicController) Update() {

	if !checkAccount(this.Ctx) { //管理员权限检查
		this.Redirect("/login", 302)
		return
	}
	id := this.Ctx.Input.Param(":id")
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	newTitle := this.Input().Get("title")
	newContent := this.Input().Get("content")
	newCategory := this.Input().Get("category")
	err = models.UpdateTopic(tid, newTitle, newContent, newCategory)
	if err != nil {
		beego.Error(err)
	}
	log.Println("okokokok")
	this.Redirect("/topic", 302)
	return

}
