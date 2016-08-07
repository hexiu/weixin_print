package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/mch/pay"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
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
	req.OpenId = fmt.Sprintf("%v", sess.Get("openid"))
	req.DeviceInfo = "WEB"
	req.NonceStr = ""
	req.Body = "打印文件支付-文件付款"
	req.Attach = "test"
	req.OutTradeNo = strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(time.Now().UnixNano()))
	req.FeeType = "CNY"
	req.TotalFee = 1
	req.SpbillCreateIP = getClientIpAddress()
	req.TimeStart = core.FormatTime(time.Now())
	req.TimeExpire = core.FormatTime(time.Now().Add(1))
	req.GoodsTag = ""
	req.NotifyURL = WebSiteUrl + "/wxpayrel"
	req.TradeType = "JSAPI"
	req.ProductId = ""
	req.LimitPay = "no_credit"
	resp, err := pay.UnifiedOrder2(wxpayClient, req)
	if err != nil {
		log.Println(err)
	}
	ctx.Data["AppId"] = resp.AppId
	ctx.Data["TimeStamp"] = strconv.Itoa(int(time.Now().Unix()))
	ctx.Data["NonceStr"] = strconv.Itoa(rand.Intn(32))
	ctx.Data["Pg"] = resp.PrepayId
	ctx.Data["SignType"] = resp.TradeType
	ctx.Data["PaySign"] = core.JsapiSign(resp.AppId, strconv.Itoa(rand.Intn(32)), strconv.Itoa(int(time.Now().Unix())), resp.PrepayId, resp.TradeType, apiKey)
	ctx.HTML(200, "wxpay")
}

func getClientIpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	os.Exit(0)
	return "0.0.0.0"
}

func WxPayRelHandler(ctx *macaron.Context, log *log.Logger) {
	body, err := ctx.Req.Body().String()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ctx.Req.Header, "\n", body)
}
