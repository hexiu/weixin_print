package controller

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	// "strconv"
	// "time"
)

func TxHandler(ctx *macaron.Context, log *log.Logger) string {
	msg := new(TxMsg)

	msg.FileType = ctx.Req.FormValue("FileType")
	msg.MsgURL = ctx.Req.FormValue("MsgURL")
	msg.OpenId = ctx.Req.FormValue("OpenId")
	msg.PrintCode = ctx.Req.FormValue("PrintCode")
	// msg.Time = ctx.Req.FormValue("Time")
	msg.MsgType = ctx.Req.FormValue("MsgType")
	// msg.PrintNum = ctx.Req.FormValue("PrintNum")
	msg.MsgInfo = ctx.Req.FormValue("MsgInfo")

	if ctx.Req.FormValue("MsgInfo") == "connect" {
		msg := new(TxMsg)
		msg.FileType = "NULL"
		msg.MsgURL = "NULL"
		msg.OpenId = "NULL"
		msg.PrintCode = "Print00001"
		// msg.Time = strconv.Itoa(int(time.Now().Unix()))
		msg.MsgType = "connect"
		// msg.PrintNum = "0"
		msg.MsgInfo = "connect"
	}
	if ctx.Req.FormValue("MsgInfo") == "quit" {
		msg := new(TxMsg)
		msg.FileType = "NULL"
		msg.MsgURL = "NULL"
		msg.OpenId = "NULL"
		// msg.PrintCode = printCode
		// msg.Time = strconv.Itoa(int(time.Now().Unix()))
		msg.MsgType = "connect"
		// msg.PrintNum = "0"
		msg.MsgInfo = "quit"
		return ""
	}
	return fmt.Sprintln(ctx.Req.PostForm)
}
