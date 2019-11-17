package main

import (
	"github.com/astaxie/beego"
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/sysinit"

	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/saneryao/bgadmin/routers"
)

func main() {
	sysinit.Init()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
