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
	beego.Router("/topic/del/:id", &controllers.TopicController{}, "get:DelTopic")        //
	beego.Router("/topic/update/:id", &controllers.TopicController{}, "get:UpTopic")      //
	beego.Router("/topic/update/:id", &controllers.TopicController{}, "post:Update")      //
	beego.Router("/topic/view/:id", &controllers.TopicController{}, "get:ViewTopic")      //
	beego.Router("/comment/add/:id", &controllers.CommentController{}, "post:AddComment") //
	beego.Router("/comment/del/:id", &controllers.CommentController{}, "post:DelComment") //
}
