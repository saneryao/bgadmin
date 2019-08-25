package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"net/url"
)

// initDb 初始化数据库连接
func initDb() {
	var err error
	runmode := beego.AppConfig.String("runmode")
	dbtype := beego.AppConfig.String("db::type")
	host := beego.AppConfig.String("db::host")
	var port int
	name := beego.AppConfig.String("db::name")
	user := beego.AppConfig.String("db::user")
	pwd := beego.AppConfig.String("db::password")
	prefix := beego.AppConfig.String("db::prefix")
	charset := beego.AppConfig.String("db::charset")
	timeloc := url.QueryEscape(beego.AppConfig.String("db::timeloc"))
	var maxIdle, maxConn int
	port, err = beego.AppConfig.Int("db::port")
	if err != nil {
		logs.Error(err)
	}
	maxIdle, err = beego.AppConfig.Int("db::maxIdle")
	if err != nil {
		logs.Error(err)
	}
	maxConn, err = beego.AppConfig.Int("db::maxConn")
	if err != nil {
		logs.Error(err)
	}

	// 若没有给出相应值，则使用默认值
	if dbtype == "" {
		dbtype = "mysql"
	}
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 3306
	}
	if name == "" {
		name = "saneryao.com"
	}
	if charset == "" {
		charset = "urf8"
	}
	if timeloc == "" {
		timeloc = "Asia%2FShanghai"
	}
	if maxIdle == 0 {
		maxIdle = 3
	}
	if maxConn == 0 {
		maxConn = 10
	}

	// 设置调试模式
	orm.Debug = (runmode == "dev")

	// mysql / sqlite3 / postgres 默认已经注册过，所以可以无需设置
	// orm.RegisterDriver("mysql", orm.DRMySQL)

	// ORM 必须注册一个别名为 default 的数据库，作为默认使用。
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s", user, pwd, host, port, name, charset, timeloc)
	err = orm.RegisterDataBase("default", dbtype, dns, maxIdle, maxConn)
	if err != nil {
		logs.Error(err, ":", dns)
	}

	orm.RegisterModelWithPrefix(prefix, new(models.User))
	orm.RegisterModelWithPrefix(prefix, new(models.Profile))
	orm.RegisterModelWithPrefix(prefix, new(models.Role))
	orm.RegisterModelWithPrefix(prefix, new(models.Menu))
	orm.RegisterModelWithPrefix(prefix, new(models.Link))
	orm.RegisterModelWithPrefix(prefix, new(models.UserRole))
	orm.RegisterModelWithPrefix(prefix, new(models.RoleMenu))
	orm.RegisterModelWithPrefix(prefix, new(models.RoleLink))
	orm.RegisterModelWithPrefix(prefix, new(models.Code))
	orm.RunSyncdb("default", false, true)
}
