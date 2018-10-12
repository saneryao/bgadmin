package sysinit

import (
	"github.com/astaxie/beego/logs"
)

/* 功能：初始化日志设置
 * 参数：空
 * 返回值：空
 */
func initLog() {
	logs.SetLogger(logs.AdapterConsole)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}
