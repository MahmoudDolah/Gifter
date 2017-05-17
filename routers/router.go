package routers

import (
	"github.com/MahmoudDolah/Gifter/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
