package main

import (
	"github.com/astaxie/beego"
	"github.com/saneryao/bgadmin/controllers"

	_ "github.com/saneryao/bgadmin/routers"
	_ "github.com/saneryao/bgadmin/sysinit"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
