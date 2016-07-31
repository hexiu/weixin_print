package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
	// // "github.com/chanxuehong/wechat.v2/mp/menu"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/response"
)

func MediaPicHandler(ctx *core.Context) {
	fmt.Println(ctx.MixedMsg)
}
