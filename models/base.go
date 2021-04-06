package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func Init()  {
	dbhost,_  := beego.AppConfig.String("dbhost")
	dbport,_  := beego.AppConfig.String("dbport")
	dbuser,_  := beego.AppConfig.String("dbuser")
	dbpasswd,_:= beego.AppConfig.String("dbpasswd")
	dbname,_  := beego.AppConfig.String("dbname")
	if dbhost=="" || dbuser ==""{
		fmt.Println("数据库配置错误")
	}
	dsn := dbuser + ":" +dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" +dbname + "?charset=utf8&loc=Asia%2FShanghai"
	// fmt.Println("dsn:%v",dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User),new(Money))
}

//返回带前缀的表名
func TableName(str string) string {
	dbprefix,_ := beego.AppConfig.String("dbprefix")
	return dbprefix + str
}
