package sysinit

import (
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

// 后台系统的配置初始化（此接口在main中的router后调用）
func init() {
	initLog()
	initCache()
	initCaptcha()
	initLocale()
	initSMTP()
	initDb()
	initSession()
}
