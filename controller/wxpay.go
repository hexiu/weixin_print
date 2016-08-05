package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/mch/pay"

	"gopkg.in/macaron.v1"
	"net/http"
)

const (
	apiKey = "cHuaNgXinYOUxianZeRen2016GSWxpAy"
	mchId  = "1373270202"
)

var wxpayServer *core.Server
var wxpayClient *core.Client = core.NewClient(WxAppId, mchId, apiKey, nil)
var req *pay.UnifiedOrderRequest

func WxpayHandler(ctx *macaron.Context) {
	var w http.ResponseWriter
	w = ctx.Resp
	fmt.Println("...")
	wxpayServer.ServeHTTP(w, ctx.Req.Request, nil)
}

func WxPayHandler(ctx *core.Context) {
	req.DeviceInfo = "WEB"
	req.NonceStr = ""
	req.Body = "打印文件支付"
	req.Detail = "文件名称"
	req.Attach = "附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据"
	req.OutTradeNo = "商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号"
	req.FeeType = "符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型"
	req.TotalFee = 100 //"订单总金额，单位为分，详见支付金额"
	req.SpbillCreateIP = "APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。"
	req.TimeStart = " 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则"
	req.TimeExpire = "订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则"
	req.NotifyURL = "接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。"
	req.TradeType = "取值如下：JSAPI，NATIVE，APP，详细说明见参数规定"
	req.ProductId = "trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。"
	req.LimitPay = "no_credit--指定不能使用信用卡支付"
	req.OpenId = "rade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。"

	pay.UnifiedOrder2(wxpayClient, req)
}
