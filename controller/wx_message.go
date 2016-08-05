package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	"net/http"
	// "strings"
	// "strconv"
	"gopkg.in/macaron.v1"
	// "weixin_dayin/modules/initConf"
)

const (
	subScribeMsgBack string = "SubScribeMsgBack = 欢迎关注创昕云打印，我帮你便捷打印"
)

var (
	SubScribeMsgBack string = subScribeMsgBack
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func WxCallbackHandler(ctx *macaron.Context) {
	// MenuHandler()
	// CreateMenu()
	// menuCreateHandler()
	var w http.ResponseWriter
	w = ctx.Resp
	msgServer.ServeHTTP(w, ctx.Req.Request, nil)
}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	fmt.Println(msg)
	msg.Content = fmt.Sprintln(msg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	ProvideQrcode()
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	msg := ctx.MixedMsg
	msg.Content = fmt.Sprintln(msg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp)
	AddMediaInfo(ctx)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	if event.EventKey == "print_start" {
		printStartHandler(ctx, event)
	}
	fmt.Println("*********************************", event)

	//resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	//ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	// DeleteMenu()

}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	msg := ctx.MixedMsg
	if ctx.MixedMsg.EventType == "scancode_waitmsg" {
		event := menu.GetClickEvent(msg)
		info := event.EventKey
		if info == "print_code" {
			printCodeHandler(ctx, event)
		}
		if event.EventKey == "print_ok" {
			printOkHandler(ctx, event)
		}
	}
	if ctx.MixedMsg.EventType == "subscribe" {
		if !ExistUser(msg.FromUserName) {
			UserAddFromWeiXinHandler(ctx)
			event := menu.GetClickEvent(msg)
			resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, SubScribeMsgBack)
			ctx.RawResponse(resp)
		} else if UserHasFromWeb(msg.FromUserName) {
			UserUpdateFormWeiXin(ctx)
			event := menu.GetClickEvent(msg)
			resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, SubScribeMsgBack)
			ctx.RawResponse(resp)
		} else {
			ChageSubscribe(msg.FromUserName, 1)
		}
	}
	if ctx.MixedMsg.EventType == "unsubscribe" {
		ChageSubscribe(msg.FromUserName, 0)
	}
	ctx.NoneResponse()
}
