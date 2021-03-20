package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"strings"
)

// CheckPermission 检查已登录用户的权限，判断其是否可以继续访问
// 返回true表示有权限可以继续访问，返回false表示权限不够无法继续访问
func CheckPermission(ctrl *beego.Controller) bool {
	//ctrlName, actionName := ctrl.GetControllerAndAction()
	//ctrlLowerName := strings.ToLower(ctrlName)
	//actionLowerName := strings.ToLower(actionName)
	//if actionLowerName == "get" && ctrlLowerName == "pagecontroller" {
	//	return true
	//}

	userInfo := GetLoginUser(ctrl)
	if userInfo.ID == 0 {
		return false
	}

	if userInfo.State == 0 {
		return false
	}

	infoChanged := false
	if userInfo.Roles == nil {
		db := orm.NewOrm()
		if _, err := db.LoadRelated(&userInfo, "Roles"); err != nil {
			logs.Error("Query user(" + userInfo.Name + ")'s roles FAIL: " + err.Error())
			return false
		}
		infoChanged = true
	}
	if len(userInfo.Roles) <= 0 {
		return false
	}
	for k, role := range userInfo.Roles {
		if role.Menus == nil {
			db := orm.NewOrm()
			if _, err := db.LoadRelated(userInfo.Roles[k], "Menus"); err != nil {
				logs.Error("Query role(" + userInfo.Name + ")'s menus FAIL: " + err.Error())
				return false
			}
		}
		if role.Links == nil {
			db := orm.NewOrm()
			if _, err := db.LoadRelated(userInfo.Roles[k], "Links"); err != nil {
				logs.Error("Query role(" + userInfo.Name + ")'s links FAIL: " + err.Error())
				return false
			}
		}
		infoChanged = true
	}

	if infoChanged {
		SetLoginUser(ctrl, userInfo)
	}
	logs.Info(userInfo)
	var menus []models.Menu
	var links []models.Link
	for _, role := range userInfo.Roles {
		for _, menu := range role.Menus {
			if menu != nil {
				menus = append(menus, *menu)
			}
		}
		for _, link := range role.Links {
			if link != nil {
				links = append(links, *link)
			}
		}
	}
	if len(menus) <= 0 && len(links) <= 0 {
		return false
	}
	SetUserMenus(ctrl, menus)
	SetUserLinks(ctrl, links)

	url := ctrl.Ctx.Request.URL.Path
	for _, link := range links {
		if strings.Index(url, link.URL) == 0 {
			return true
		}
	}

	if url == "/" {
		return true
	}

	return false
}
