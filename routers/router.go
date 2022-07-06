package routers

import (
    "base_api/controllers"
    beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{});
    beego.Router("/token", &controllers.AuthController{});
}
