package main

import (
	_ "one_Blog/models"
	model "one_Blog/models"
	_ "one_Blog/routers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func init() {
	model.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.SetStaticPath("/topic/static", "static")
	beego.Run(":8658")

}
