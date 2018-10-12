package utils

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

/* 功能：通过html请求信息中获取语言配置，并设置国际化语言
 * 参数：ctrl（控制器）
 * 返回值：redirect（true表示语言进行了切换需要跳转页面，false表示没有切换不需要跳转）
 *                 lang （设置的网页语言）
 */
func SetLangVer(ctrl *beego.Controller) (redirect bool, lang string) {
	hasCookie := false

	// 1. 从URL中获取语言配置
	lang = ctrl.Input().Get("lang")

	// 2. 从cookies中获取语言配置
	if len(lang) == 0 {
		lang = ctrl.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		redirect = true
	}

	// 3. 若有人故意修改，请再次检查
	if !i18n.IsExist(lang) {
		lang = ""
		redirect = false
		hasCookie = false
	}

	// 4. 从html请求头'Accept-Language'中获取语言配置
	if len(lang) == 0 {
		al := ctrl.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // 只比较前5个字符
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 5. 没有检测到语言配置，默认使用“中文简体”
	if len(lang) == 0 {
		lang = "zh-CN"
		redirect = false
	}

	// 6. 保存语言配置到cookies
	if !hasCookie {
		ctrl.Ctx.SetCookie("lang", lang, 1<<31-1, "/")
	}

	// 设置国际化语言属性（之后的控制器和模板文件通过此属性进行国际化）
	// ctrl.Lang = lang
	// ctrl.Data["Lang"] = ctrl.Lang
	ctrl.Data["Lang"] = lang

	return
}
