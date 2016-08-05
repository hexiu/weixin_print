package controller

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	"strconv"
	"time"
)

func TxHandler(ctx *macaron.Context, log *log.Logger) {
	fmt.Println(ctx.Req.Form)
	ctx.Data["FileType"] = "NULL"
	ctx.Data["MsgURL"] = "NULL"
	ctx.Data["OpenId"] = "NULL"
	ctx.Data["PrintCode"] = "Print00001"
	ctx.Data["Time"] = strconv.Itoa(int(time.Now().Unix()))
	ctx.Data["PrintNum"] = "0"
	ctx.Data["MsgInfo"] = "connect"
	ctx.Data["MsgType"] = "connect"
	fmt.Println(ctx.Resp)
}
