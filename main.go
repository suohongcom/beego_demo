package main

import (
	_ "github.com/hwkkevin/beego_demo/routers"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/hwkkevin/beego_demo/models"
)
func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

