package routers

import (
	"github.com/astaxie/beego"
	admapi "github.com/saneryao/bgadmin/api/v1/admin"
	pubapi "github.com/saneryao/bgadmin/api/v1/public"
	admview "github.com/saneryao/bgadmin/controllers/admin"
	pubview "github.com/saneryao/bgadmin/controllers/public"
)

func init() {
	beego.Router("/login", &pubview.LoginController{}, "GET:Get")
	beego.Router("/logout", &pubview.LogoutController{}, "GET:Get")
	beego.Router("/resetpwd", &pubview.ResetPwdController{}, "GET:Get")
	beego.Router("/active/:id:int/:code:string", &pubview.ActiveController{}, "GET:Get")
	beego.Router("/api/v1/login", &pubapi.LoginApi{}, "POST:Post")
	beego.Router("/api/v1/logout", &pubapi.LogoutApi{}, "POST:Post")
	beego.Router("/api/v1/register", &pubapi.RegisterApi{}, "POST:Post")
	beego.Router("/api/v1/findpwd", &pubapi.FindPwdApi{}, "POST:Post")
	beego.Router("/api/v1/resetpwd", &pubapi.ResetPwdApi{}, "POST:Post")

	beego.Router("/", &admview.PageController{}, "GET:Get")
	beego.Router("/api/v1/users/?:id:int", &admapi.UsersApi{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/roles/?:id:int", &admapi.RolesApi{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/menus/?:id:int", &admapi.MenusApi{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/statistics/:entry([a-zA-Z]+)", &admapi.StatisticsApi{}, "GET:Get")
}
