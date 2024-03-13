package handler

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func PanicRecovery(ctx web.Controller) {
	if r := recover(); r != nil {
		logs.Error("%v in %v", r, ctx)
		// logs.GetLogger("ORM").Println("this is a message of orm")

		ctx.Data["json"] = "Panic occured"
		ctx.ServeJSON()
		return
	}
}
