package controller

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
)

var Title string = "维宝宝打印系统"
var WebSiteName string = "维宝宝打印系统"

func HomeHandler(ctx *macaron.Context, logger *log.Logger, sess session.Store) {
	ctx.Data["Title"] = "创昕云打印"
	log.Println(sess.ID())
	ctx.Data["WebSiteName"] = WebSiteName
	ctx.HTML(200, "home")
}
