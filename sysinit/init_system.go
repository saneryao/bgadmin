package sysinit

// 后台系统的配置初始化（此接口在main中的router后调用）
func Init() {
	initLog()
	initCache()
	initCaptcha()
	initLocale()
	initSMTP()
	initDb()
	initSession()
}
