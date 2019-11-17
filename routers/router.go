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
	beego.Router("/api/v1/login", &pubapi.LoginAPI{}, "POST:Post")
	beego.Router("/api/v1/logout", &pubapi.LogoutAPI{}, "POST:Post")
	beego.Router("/api/v1/register", &pubapi.RegisterAPI{}, "POST:Post")
	beego.Router("/api/v1/findpwd", &pubapi.FindPwdAPI{}, "POST:Post")
	beego.Router("/api/v1/resetpwd", &pubapi.ResetPwdAPI{}, "POST:Post")

	beego.Router("/", &admview.PageController{}, "GET:Get")
	beego.Router("/api/v1/user/?:id:int", &admapi.UserAPI{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/role/?:id:int", &admapi.RoleAPI{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/menu/?:id:int", &admapi.MenuAPI{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/link/?:id:int", &admapi.LinkAPI{}, "GET:Get;POST:Post;PUT:Put;PATCH:Patch;DELETE:Delete")
	beego.Router("/api/v1/user_role/?:id:int", &admapi.UserRoleAPI{}, "GET:Get;POST:Post;DELETE:Delete")
	beego.Router("/api/v1/role_power/?:id:int", &admapi.RolePowerAPI{}, "GET:Get;POST:Post;DELETE:Delete")
	beego.Router("/api/v1/statistics/:entry([a-zA-Z]+)", &admapi.StatisticsAPI{}, "GET:Get")
}
