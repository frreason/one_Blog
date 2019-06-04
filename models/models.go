package models

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"

	"github.com/Unknwon/com"

	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME       = "data/web_blog.db"
	_SQLITE_DRIVER = "sqlite3"
)

type Category struct { //文章分类
	Id      int64
	Title   string
	Total   int64     `orm:"index"`
	Created time.Time `orm:"index"`
	Views   int64     `orm:"index"`
	Updated time.Time `orm:"index"`
}
type Topic struct { //每一个文章 Id对应Comments的Id
	Id                int64
	Title             string
	Content           string `orm:"size(5000)"`
	Author            string
	Views             int64     `orm:"index"`
	Created           time.Time `orm:"index"`
	Updated           time.Time `orm:"index"`
	CommentLastUserId int64
	CommentCount      int64 `orm:"index"`
	CommentTime       time.Time
	Category          string
	label             string
}

type Author struct { //作者信息 就是本人。。。
	Id    int64 `orm:"index"`
	Name  string
	Email string
}

type Comments struct { //评论信息表 Tid是外码
	Id      int64
	Tid     int64     //对应topic里的id
	Created time.Time `orm:"index"`
	Content string    `orm:"size(5000)"`
	Writer  string
	Floor   int
}

func RegisterDB() {
	log.Println(path.Dir(_DB_NAME))
	if !com.IsExist(_DB_NAME) {
		log.Println(path.Dir(_DB_NAME))
		os.Mkdir(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	//先注册数据表 RegisterModel
	//然后注册数据库驱动 RegisterDriver
	//再然后注册数据库
	orm.RegisterModel(new(Category), new(Topic), new(Comments))
	orm.RegisterDriver(_SQLITE_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE_DRIVER, _DB_NAME, 10) //设置数据库连接参数  "root:061365404@"+_DB_NAME+"?charset=utf8"
	orm.DefaultTimeLoc = time.UTC
}

func AddCategory(titleName string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("Category")

	now := time.Now()
	timeAdd, err := time.ParseDuration("+8h")
	realNow := now.Add(timeAdd)
	oneCategory := &Category{
		Title:   titleName,
		Created: realNow,
		Updated: realNow,
	}

	err = qs.Filter("Title", titleName).One(oneCategory)
	if err == nil {
		return nil //nil 说明正常
	}
	_, err = o.Insert(oneCategory)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategory(totalMax bool) ([]*Category, error) {

	o := orm.NewOrm()
	qs := o.QueryTable("Category")

	categories := make([]*Category, 0)
	var err error
	if totalMax {
		_, err = qs.OrderBy("-total").All(&categories)
	} else {
		_, err = qs.All(&categories)
	}
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func DeleteCategory(id int64) error {

	o := orm.NewOrm()

	oneCategory := &Category{
		Id: id,
	}
	_, err := o.Delete(oneCategory)
	if err != nil {
		return err
	}
	return nil
}

func AddTopic(title, content, category string) error {
	o := orm.NewOrm()

	now := time.Now()
	timeAdd, err := time.ParseDuration("+8h") //因为伦敦时间慢8小时
	realNow := now.Add(timeAdd)

	oneTopic := &Topic{
		Title:       title,
		Content:     content,
		Created:     realNow,
		Updated:     realNow,
		Category:    category,
		CommentTime: realNow,
		Author:      beego.AppConfig.String("userName"),
	}
	_, err = o.Insert(oneTopic)
	if err != nil {
		return err
	}
	qs := o.QueryTable("Category")
	oneCategory := &Category{}
	err = qs.Filter("Title", category).One(oneCategory)

	if err != nil { //在category表中新增
		oneCategory.Title = category
		oneCategory.Total += 1
		oneCategory.Created = realNow
		oneCategory.Updated = realNow
		_, err = o.Insert(oneCategory)
		if err != nil {
			return err
		}
	} else {
		oneCategory.Total += 1
		oneCategory.Updated = realNow
		_, err = o.Update(oneCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetTopic(tid int64) (*Topic, error) {

	o := orm.NewOrm()
	qs := o.QueryTable("Topic")

	oneTopic := &Topic{}
	err := qs.Filter("Id", tid).One(oneTopic)
	if err != nil {
		return nil, err
	}
	oneTopic.Views++
	_, err = o.Update(oneTopic)
	if err != nil {
		return nil, err
	}

	return oneTopic, nil
}

func GetAllTopic(viewsMax, lastest bool) ([]*Topic, error) {

	o := orm.NewOrm()
	qs := o.QueryTable("Topic")

	topics := make([]*Topic, 0)

	var err error
	if viewsMax == true {
		_, err = qs.OrderBy("-Views").All(&topics)
	} else {
		if lastest {
			_, err = qs.OrderBy("Updated").All(&topics)
		} else {
			_, err = qs.All(&topics)
		}
	}
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func UpdateTopic(tid int64, title, content, category string) error { //tid是必须的，假如title和content变了的话我无从知道该更新到哪条记录里
	//并且需要和老的数据进行一些对比
	o := orm.NewOrm()
	qs := o.QueryTable("Topic")

	oneTopic := &Topic{}
	err := qs.Filter("Id", tid).One(oneTopic) //读出旧文章，好进行对比
	if err != nil {                           //说明没有这篇文章，那自然就不能进行更新
		return err
	}
	if oneTopic.Title == title && oneTopic.Content == content && oneTopic.Category == category { //无需改动
		return nil
	}

	now := time.Now()
	timeAdd, err := time.ParseDuration("+8h") //因为伦敦时间慢8小时
	realNow := now.Add(timeAdd)

	categoryQs := o.QueryTable("Category")
	oneCategory := &Category{}
	err = categoryQs.Filter("Title", category).One(oneCategory)
	if oneTopic.Category == category { //文章分类没有变
		oneCategory.Updated = realNow
		_, err = o.Update(oneCategory)
		if err != nil {
			return err
		}
	} else { //文章分类变化
		if err != nil { //说明原来的category表中没有该类别，则增加一条记录
			oneCategory.Created = realNow
			oneCategory.Title = category
			oneCategory.Total += 1
			oneCategory.Updated = realNow
			_, err = o.Insert(oneCategory)
			if err != nil {
				return err
			}
		} else {
			oneCategory.Total += 1
			oneCategory.Updated = realNow
			_, err = o.Update(oneCategory)
			if err != nil {
				return err
			}
		}
	}
	//应该让旧分类total-1
	err = categoryQs.Filter("Title", oneTopic.Category).One(oneCategory)

	if err != nil {
		return err
	}
	oneCategory.Total -= 1
	oneCategory.Updated = realNow
	_, err = o.Update(oneCategory)
	if err != nil {
		return err
	}
	oneTopic.Category = category
	oneTopic.Title = title
	oneTopic.Content = content
	//更新记录
	_, err = o.Update(oneTopic)
	if err != nil {
		return err
	}
	return nil
}

//还得把相应得评论删除
func DeleteTopic(tid int64) error {

	now := time.Now()
	timeAdd, err := time.ParseDuration("+8h") //因为伦敦时间慢8小时
	realNow := now.Add(timeAdd)

	o := orm.NewOrm()

	oneTopic := &Topic{
		Id: tid,
	}
	qs := o.QueryTable("Topic")
	err = qs.Filter("Id", tid).One(oneTopic)
	if err != nil {
		return err
	}
	oneCategory := &Category{}
	categoryQs := o.QueryTable("Category")
	err = categoryQs.Filter("Title", oneTopic.Category).One(oneCategory)
	if err != nil {
		return err
	}
	oneCategory.Total -= 1
	oneCategory.Updated = realNow
	_, err = o.Update(oneCategory)
	if err != nil {
		return err
	}
	_, err = o.Delete(oneTopic)
	if err != nil {
		return err
	}
	return nil
}

func AddComment(tid int64, nickName, comment string) error {

	now := time.Now()
	timeAdd, err := time.ParseDuration("+8h") //因为伦敦时间慢8小时
	realNow := now.Add(timeAdd)

	o := orm.NewOrm()
	qs := o.QueryTable("Topic")
	oneTopic := &Topic{
		Id: tid,
	}
	err = qs.Filter("Id", tid).One(oneTopic)

	if err != nil {
		return err
	}
	oneTopic.CommentCount += 1
	oneTopic.CommentTime = realNow //最新的评论时间
	_, err = o.Update(oneTopic)
	if err != nil {
		return err
	}

	oneComment := &Comments{
		Tid:     tid,
		Writer:  nickName,
		Content: comment,
		Created: realNow,
	}
	_, err = o.Insert(oneComment)
	if err != nil {
		return err
	}

	return nil
}

//读取的是对应的文章的评论
func GetComment(tid int64) ([]*Comments, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Comments")

	comments := make([]*Comments, 0)
	_, err := qs.Filter("Tid", tid).All(&comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetAllComment(lastest bool) ([]*Comments, error) {

	o := orm.NewOrm()
	qs := o.QueryTable("Comments")
	comments := make([]*Comments, 0)
	var err error

	if lastest {
		_, err = qs.OrderBy("Created").All(&comments)
	} else {
		_, err = qs.All(&comments)
	}

	if err != nil {
		return nil, err
	}
	return comments, err
}

//删除Comments中的一条评论以及相应的topic评论total-1
func DeleteComment(cid int64) error {

	o := orm.NewOrm()

	qs := o.QueryTable("Comments")

	oneComment := &Comments{
		Id: cid,
	}
	err := qs.Filter("Id", cid).One(oneComment)
	if err != nil {
		return err
	}

	tid := oneComment.Tid
	oneTopic := &Topic{
		Id: tid,
	}
	qs = o.QueryTable("Topic")
	err = qs.Filter("Id", tid).One(oneTopic)
	if err != nil {
		return err
	}
	oneTopic.CommentCount--
	_, err = o.Update(oneTopic)
	if err != nil {
		return err
	}

	_, err = o.Delete(oneComment)
	if err != nil {
		return err
	}

	return nil
}

func ListCategoryTop(cid int64) ([]*Topic, error) {

	o := orm.NewOrm()
	qs := o.QueryTable("Category")

	cate := &Category{
		Id: cid,
	}

	err := qs.Filter("Id", cid).One(cate)
	if err != nil {
		return nil, err
	}
	cateName := cate.Title
	qs = o.QueryTable("Topic")
	topics := make([]*Topic, 0)
	_, err = qs.Filter("Category", cateName).All(&topics)
	if err != nil {
		return nil, err
	}

	return topics, nil
}
