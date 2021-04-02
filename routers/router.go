package routers

import (
	"github.com/hwkkevin/beego_demo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)
 
func init() {
    beego.Router("/login", &controllers.PassportController{},"*:Login")
	beego.Router("/logout", &controllers.PassportController{},"*:Logout")
	beego.Router("/", &controllers.IndexController{},"GET:Index")
	beego.Router("/create", &controllers.IndexController{},"*:Create")
}
