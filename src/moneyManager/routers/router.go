package routers

import (
	"github.com/astaxie/beego"
	"moneyManager/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MyController{})
}
