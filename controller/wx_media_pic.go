package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
)

func MediaPicHandler(ctx *core.Context) {
	fmt.Println(ctx.MixedMsg)
}
