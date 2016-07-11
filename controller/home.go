package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func HomeHandler(ctx *macaron.Context, logger *log.Logger) {
	ctx.Data["TITLE"] = "维宝宝打印系统"
	ctx.Data["WebSiteName"] = "维宝宝打印系统"
	ctx.HTML(200, "index")
}
