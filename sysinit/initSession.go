package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var SessionMgr *session.Manager

/* 功能：初始化session设置
 * 参数：空
 * 返回值：空
 */
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
