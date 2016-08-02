package controller

import (
	"gopkg.in/macaron.v1"
	// "log"
)

var errinfo string
var gotourl string

func ErrHandler(ctx *macaron.Context) {
	ctx.Data["ErrInfo"] = errinfo
	ctx.Data["GoToUrl"] = gotourl
	ctx.HTML(200, "error")
}
