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

	return
}

func (this *TopicController) ListCategoryTopic() {
	id := this.Ctx.Input.Param(":id")
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	topics, err := models.ListCategoryTop(cid)
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "home.html"

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

	return
}
