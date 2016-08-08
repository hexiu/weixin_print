package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/mch/pay"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"math/rand"
	// "net"
	"net/http"
	// "os"
	"crypto/md5"
	rand1 "github.com/chanxuehong/rand"
	"strconv"
	"strings"
	"time"
)

const (
	apiKey = "cHuaNgXinYOUxianZeRen2016GSWxpAy"
	mchId  = "1373270202"
)

var wxpayServer *core.Server
var wxpayClient *core.Client = core.NewClient(WxAppId, mchId, apiKey, nil)
var req *pay.UnifiedOrderRequest = new(pay.UnifiedOrderRequest)

func WxpayHandler(ctx *macaron.Context) {
	var w http.ResponseWriter
	w = ctx.Resp
	fmt.Println("...")
	wxpayServer.ServeHTTP(w, ctx.Req.Request, nil)
}

func wxpayhandler(ctx *core.Context) {

}

func WxPayHandler(ctx *macaron.Context, sess session.Store) {
	sid := ctx.GetCookie(cookieName)
	ctx.Data["IsWxPay"] = true
	if sess.ID() != sid {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	ctx.Data["Title"] = "打印文件支付"
	req.OpenId = fmt.Sprintf("%v", sess.Get("openid"))
	req.DeviceInfo = "WEB"
	req.NonceStr = fmt.Sprintf("%f", rand.Float32())
	req.Body = "打印文件支付-文件付款"
	req.Attach = "test"
	req.OutTradeNo = strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(time.Now().UnixNano()))
	req.FeeType = "CNY"
	req.TotalFee = 1
	remoteAddr := strings.Split(ctx.Req.RemoteAddr, ":")
	req.SpbillCreateIP = remoteAddr[0]
	req.TimeStart = core.FormatTime(time.Now())
	req.TimeExpire = core.FormatTime(time.Now().AddDate(0, 0, 1))
	req.GoodsTag = ""
	req.NotifyURL = WebSiteUrl + "/wxpayrel"
	req.TradeType = "JSAPI"
	req.ProductId = ""
	req.LimitPay = "no_credit"

	fmt.Println("req::", req)
	fmt.Println(wxpayClient)
	resp, err := UnifiedOrder2(wxpayClient, req)
	fmt.Println(resp)
	fmt.Println(wxpayClient)
	if err != nil {
		log.Println(err)
	}
	ctx.Data["AppId"] = resp.AppId
	ctx.Data["TimeStamp"] = strconv.Itoa(int(time.Now().Unix()))
	ctx.Data["NonceStr"] = req.NonceStr
	ctx.Data["Pg"] = resp.PrepayId
	ctx.Data["SignType"] = resp.TradeType
	ctx.Data["PaySign"] = core.JsapiSign(resp.AppId, req.NonceStr, strconv.Itoa(int(time.Now().Unix())), resp.PrepayId, "MD5", apiKey)
	ctx.HTML(200, "wxpay")
}

func WxPayRelHandler(ctx *macaron.Context, log *log.Logger) {
	body, err := ctx.Req.Body().String()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ctx.Req.Header, "\n", body)
}

func UnifiedOrder2(clt *core.Client, req *pay.UnifiedOrderRequest) (resp *pay.UnifiedOrderResponse, err error) {
	m1 := make(map[string]string, 20)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	if req.DeviceInfo != "" {
		m1["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = string(rand1.NewHex())
	}
	m1["body"] = req.Body
	if req.Detail != "" {
		m1["detail"] = req.Detail
	}
	if req.Attach != "" {
		m1["attach"] = req.Attach
	}
	m1["out_trade_no"] = req.OutTradeNo
	if req.FeeType != "" {
		m1["fee_type"] = req.FeeType
	}
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["spbill_create_ip"] = req.SpbillCreateIP
	if req.TimeStart != "" {
		m1["time_start"] = req.TimeStart
	}
	if req.TimeExpire != "" {
		m1["time_expire"] = req.TimeExpire
	}
	if req.GoodsTag != "" {
		m1["goods_tag"] = req.GoodsTag
	}
	m1["notify_url"] = req.NotifyURL
	m1["trade_type"] = req.TradeType
	if req.ProductId != "" {
		m1["product_id"] = req.ProductId
	}
	if req.LimitPay != "" {
		m1["limit_pay"] = req.LimitPay
	}
	if req.OpenId != "" {
		m1["openid"] = req.OpenId
	}
	m1["sign"] = core.Sign(m1, clt.ApiKey(), md5.New)

	fmt.Println(m1)
	m2, err := UnifiedOrder(clt, m1)
	fmt.Println(m2, "This is resp.", err)
	if err != nil {
		return
	}
	fmt.Println(m2, "This is resp.")

	// 判断业务状态
	resultCode, ok := m2["result_code"]
	if !ok {
		err = core.ErrNotFoundResultCode
		return
	}
	if resultCode != core.ResultCodeSuccess {
		err = &core.BizError{
			ResultCode:  resultCode,
			ErrCode:     m2["err_code"],
			ErrCodeDesc: m2["err_code_des"],
		}
		return
	}

	resp = &pay.UnifiedOrderResponse{
		AppId: m2["appid"],
		MchId: m2["mch_id"],

		TradeType:  m2["trade_type"],
		PrepayId:   m2["prepay_id"],
		DeviceInfo: m2["device_info"],
		CodeURL:    m2["code_url"],
		MWebURL:    m2["mweb_url"],
	}
	return
}

func UnifiedOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/unifiedorder", req)
}

// func test() {
// 	client := http.DefaultClient
// 	resp, err := client.Get(webSiteUrl)
// 	body, err := resp.Body
// 	resp.Request.Body
// }
