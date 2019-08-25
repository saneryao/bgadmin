package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

// SessionMgr 定义全局可用的session管理器
var SessionMgr *session.Manager

// initSession 初始化session设置
func initSession() {
	host := beego.AppConfig.String("redis::host")
	port, _ := beego.AppConfig.Int("redis::port")
	pwd := beego.AppConfig.String("redis::password")
	pool, _ := beego.AppConfig.Int("redis::pool")
	var err error
	defer func() {
		if r := recover(); r != nil {
			SessionMgr = nil
		}
	}()

	strParams := fmt.Sprintf("%s:%d,%d,%s", host, port, pool, pwd)
	sessionConfig := &session.ManagerConfig{
		CookieName:      "seygosid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  strParams,
	}
	SessionMgr, err = session.NewManager("redis", sessionConfig)
	if err != nil {
		logs.Error("Connect to the redis host(" + host + ") failed")
		logs.Error(err)
	} else {
		go SessionMgr.GC()
	}
}
