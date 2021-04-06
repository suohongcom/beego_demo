package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"fmt"
	"strconv"
)

type BaseController struct{
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName string
}

func (p *BaseController) Prepare()  {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	fmt.Println("我开始进入")
	p.o = orm.NewOrm()
	fmt.Printf("请求方法%d",strings.ToLower(p.actionName))
	if strings.ToLower(p.actionName) !=  "login"{
		if p.GetSession("user") == nil{
			p.History("未登录","/login")
		}
	}
}

func (p *BaseController) History(msg string ,url string) {
	if url == ""{
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1)</script>")
		p.StopRun()
	}else {
		p.Redirect(url,302)
	}
}

func (p *BaseController) getClientIp() string  {
	s := strings.Split(p.Ctx.Request.RemoteAddr , ":")
	return s[0]
}


func (p *BaseController) GetTemplate(name string)  {
	p.Data["Header"] = name
	p.Layout = "layout.tpl"
	p.TplName = name+".tpl"
}


func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}