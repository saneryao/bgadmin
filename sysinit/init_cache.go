package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var CacheRedis cache.Cache

/* 功能：初始化缓存（缓存数据通过redis进行）
 * 参数：空
 * 返回值：空
 */
func initCache() {
	host := beego.AppConfig.String("redis::host")
	port, _ := beego.AppConfig.Int("redis::port")
	pwd := beego.AppConfig.String("redis::password")
	var err error
	defer func() {
		if r := recover(); r != nil {
			CacheRedis = nil
		}
	}()
	strParams := fmt.Sprintf(`{"key":"seygockey","conn":"%s:%d","password":"%s"}`, host, port, pwd)
	CacheRedis, err = cache.NewCache("redis", strParams)
	if err != nil {
		logs.Error("Connect to the redis host(" + host + ") failed")
		logs.Error(err)
	}
}
