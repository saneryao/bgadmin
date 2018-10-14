package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

/* 功能：初始化语言国际化设置
 * 参数：空
 * 返回值：空
 */
func initLocale() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
}
