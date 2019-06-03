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
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	topics, err := models.GetAllTopic(false, false) //取出所有文章信息
	if err != nil {
		beego.Error(err)
	}

	this.Data["Topics"] = topics

}

func (this *TopicController) Add() {

	if !checkAccount(this.Ctx) { //管理员权限检查
		this.Redirect("/login", 302)
		return
	}
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

//更新页面初始化
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

//更新操作
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

func (this *TopicController) ViewTopic() {
	id := this.Ctx.Input.Param(":id")
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "viewTopic.html"
	this.Data["IsTopic"] = true

	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
	}
	log.Println(tid)
	this.Data["Tid"] = tid
	this.Data["Title"] = topic.Title
	this.Data["Content"] = topic.Content
	this.Data["Created"] = topic.Created
	this.Data["Views"] = topic.Views
	this.Data["Author"] = beego.AppConfig.String("userName")
	this.Data["Admin"] = checkAccount(this.Ctx)
	comments := make([]*models.Comments, 0)

	comments, err = models.GetComment(tid)
	if err != nil {
		beego.Error(err)
	}
	for i, v := range comments {
		v.Floor = i + 1
	}
	this.Data["Comments"] = comments

	return
}
