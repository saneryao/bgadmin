package sysinit

/* 功能：后台系统的初始化配置（此接口在main中的router后调用）
 * 参数：空
 * 返回值：空
 */
func init() {
	initLog()
	initCache()
	initCaptcha()
	initLocale()
	initDb()
	initEmail()
}
