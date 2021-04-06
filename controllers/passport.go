package controllers

import (
	"github.com/beego/beego/v2/client/cache"
	captchas "github.com/beego/beego/v2/server/web/captcha"
	"github.com/hwkkevin/beego_demo/models"
	"fmt"
	"github.com/hwkkevin/beego_demo/util"
	"strings"
)

var captcha *captchas.Captcha

type UserData struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Captcha  string `form:"captcha"`
	CaptchaId  string `form:"captcha_id"`
}

type PassportController struct{
	BaseController
}
//init函数初始化captcha
func init() {
	//验证码功能
	//使用Beego缓存存储验证器数据
	store := cache.NewMemoryCache()
	//创建验证码
	captcha = captchas.NewWithFilter("/captcha",store)
	//设置验证码长度
	captcha.ChallengeNums = 4
	//设置验证码模板高度
	captcha.StdHeight = 38
	//设置验证码模板高度 
	captcha.StdWidth = 90
}

func (b *PassportController) Login()  {
	var users UserData
	fmt.Printf("<<<<<<<<<存储是session数据%d>>>>>>>>",b.GetSession("user"))
	if b.Ctx.Request.Method == "POST"{
		if !captcha.VerifyReq(b.Ctx.Request) {
			b.History("验证码错误","")
		}
		if err := b.ParseForm(&users); err != nil{
			b.History("表单数据为空","")
		}

		userModel := models.User{Username:users.Username} 
		if err := b.o.Read(&userModel,"Username"); err !=nil{
			b.History("账号不存在","")
		}
		fmt.Println("用户输入密码%d,数据库查询的密码%d",util.Md5(users.Password),strings.Trim(userModel.Password," "))
		if util.Md5(users.Password) != strings.Trim(userModel.Password," "){
			b.History("密码错误","")
		}
		userModel.LastIp = b.getClientIp()
		userModel.LoginCount = userModel.LoginCount + 1
		if num,err := b.o.Update(&userModel); err !=nil{
			fmt.Printf("返回影响的行数%d",num)
			b.History("更新失败","")
		}else{
			b.SetSession("user",userModel)
			fmt.Printf("<<<<<<<<<<<我存储的session数据:%d>>>>>>>>>>",b.GetSession("user"))
			b.History("登录成功","/")
		}

	}
	b.GetTemplate("login")
}


func (b *PassportController) Logout()  {
	b.DelSession("user")
	b.History("退出成功","/login")
}