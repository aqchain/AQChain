package routers

import (
	"AQChain/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
}
