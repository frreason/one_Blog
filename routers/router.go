package routers

import (
	"one_Blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/topic/add", &controllers.TopicController{}, "get:Add")
	beego.Router("/topic/add", &controllers.TopicController{}, "post:AddTopic")
	beego.Router("/topic/del/:id", &controllers.TopicController{}, "get:DelTopic")        //正则匹配
	beego.Router("/topic/update/:id", &controllers.TopicController{}, "get:UpTopic")      //正则匹配
	beego.Router("/topic/update/:id", &controllers.TopicController{}, "post:Update")      //正则匹配
	beego.Router("/topic/view/:id", &controllers.TopicController{}, "get:ViewTopic")      //正则匹配
	beego.Router("/comment/add/:id", &controllers.CommentController{}, "post:AddComment") //正则匹配
}
