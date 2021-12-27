package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Post("/",func(ctx *context.Context){
    	ctx.Output.Body([]byte("hello world"))
})
}

