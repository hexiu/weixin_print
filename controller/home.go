package controller

import (
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
)

func HomeHandler(ctx *macaron.Context, logger *log.Logger, sess session.Store) {
	ctx.Data["TITLE"] = "维宝宝打印系统"
	log.Println(sess.ID())
	ctx.Data["WebSiteName"] = "维宝宝打印系统"
	ctx.HTML(200, "home")
}
