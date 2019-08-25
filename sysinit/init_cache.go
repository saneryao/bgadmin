package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
)

// CacheRedis 定义全局可用缓存（使用redis缓存数据）
var CacheRedis cache.Cache

// initCache 初始化缓存（连接redis）
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
