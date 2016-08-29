package controller

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/mch/pay"
	mpcore "github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"hash"
	"log"
	"math/rand"
	// "net"
	"net/http"
	// "os"
	rand1 "github.com/chanxuehong/rand"
	// testapi "github.com/chanxuehong/wechat.v2/internal/debug/api"
	"github.com/chanxuehong/wechat.v2/json"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
	// mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	// "errors"
	// "encoding/xml"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"
	"weixin_dayin/models"
)

const (
	apiKey = "cHuaNgXinYOUxianZeRen2016GSWxpAy"
	mchId  = "1373270202"
)

var wxpayHandlerFunc core.HandlerFunc = wxPayRelHandler
var wxpayHandler core.Handler = wxpayHandlerFunc

var tokenServer mpcore.AccessTokenServer
var wxpayServer *core.Server
var wxpayClient *core.Client
var req *pay.UnifiedOrderRequest = new(pay.UnifiedOrderRequest)
var wxpay core.HandlerChain = make(core.HandlerChain, 1)

// func WxpayHandler(ctx *macaron.Context) {
// 	var w http.ResponseWriter
// 	w = ctx.Resp
// 	fmt.Println("...")
// 	wxpayServer.ServeHTTP(w, ctx.Req.Request, nil)
// }

// func wxpayhandler(ctx *core.Context) {

// }

func WxPayHandler(ctx *macaron.Context, sess session.Store) {
	sid := ctx.GetCookie(cookieName)
	ctx.Data["IsWxPay"] = true
	if sess.ID() != sid {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		// ctx.Redirect("/errorinfo", 301)
	}

	op := ctx.Req.FormValue("op")
	id := ctx.Req.FormValue("id")

	if op == "pay" {
		fid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println(err)
			return
		}

		fileinfo, err := models.GetFileInfo(fid)
		if err != nil {
			log.Println(err)
			errinfo = "文件未找到，请重新上传文件，请点击下面的链接进行跳转。"
			gotourl = WebSiteUrl + "/file"
			ctx.Redirect("/errorinfo", 301)
		}
		ctx.Data["Title"] = "打印文件支付"
		// req.OpenId = sess.Get("openid")
		req.OpenId = fmt.Sprintf("%v", sess.Get("openid"))
		req.DeviceInfo = "WEB"
		req.NonceStr = fmt.Sprintf("%f", rand.Int31n(int32(time.Now().UnixNano()%1000000000)))
		req.Body = "打印文件支付-文件付款"
		req.Attach = "创昕云打印文件支付信息"
		req.OutTradeNo = fileinfo.OutTradeNo
		req.FeeType = "CNY"
		req.TotalFee = 1
		remoteAddr := strings.Split(ctx.Req.RemoteAddr, ":")
		req.SpbillCreateIP = remoteAddr[0]
		req.TimeStart = core.FormatTime(time.Now())
		req.TimeExpire = core.FormatTime(time.Now().AddDate(0, 0, 1))
		req.GoodsTag = fileinfo.FileType
		req.NotifyURL = WebSiteUrl + "/wxpayrel"
		req.TradeType = "JSAPI"
		req.ProductId = strconv.Itoa(int(fileinfo.Id))
		req.LimitPay = "no_credit"

		// fmt.Println("req::", req)
		// fmt.Println(wxpayClient)
		resp, _, m2, err := UnifiedOrder2(wxpayClient, req)
		// fmt.Println(resp)
		// fmt.Println(wxpayClient)
		if err != nil {
			log.Println(err)
		}

		accessToken, err := tokenServer.Token()
		if err != nil {
			log.Println("AccessToken:", err)
		}

		// fmt.Println("AccessToken :", accessToken)
		url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + url.QueryEscape(accessToken) + "&type=jsapi"
		// fmt.Println("URL:", url)
		clt := http.DefaultClient
		httpResp, err := clt.Get(url)
		if err != nil {
			log.Println(accessToken, err)
		}
		defer httpResp.Body.Close()
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			log.Println(err)
		}
		// fmt.Println("Body:", body, string(body))

		if httpResp.StatusCode != http.StatusOK {
			err = fmt.Errorf("http.Status: %s", httpResp.Status)
			return
		}

		var result map[string]interface{}

		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Println(err)
		}
		// fmt.Println(result)
		if err != nil && err != io.EOF {
			// errors.Add([]string{}, ERR_DESERIALIZATION, err.Error())
			log.Println(err)
		}
		// fmt.Println("Result:", result, httpResp.Header)

		timeStamp := strconv.Itoa(int(time.Now().Unix()))
		jsapiTicket := fmt.Sprintf("%v", result["ticket"])

		ctx.Data["WxConfigSign"] = jssdk.WXConfigSign(jsapiTicket, m2["nonce_str"], timeStamp, WebSiteUrl+"/wxpay/wxpay/wxpay")
		ctx.Data["AppId"] = WxAppId
		ctx.Data["NonceStr"] = m2["nonce_str"]
		// ctx.Data["NonceStr"] = req.NonceStr
		ctx.Data["Pg"] = "prepay_id=" + resp.PrepayId
		ctx.Data["SignType"] = "MD5"

		// paySign := core.JsapiSign(resp.AppId, m2["nonce_str"], strconv.Itoa(int(time.Now().Unix())), resp.PrepayId, "MD5", apiKey)
		// timeStamp = strconv.Itoa(int(time.Now().Unix()))
		ctx.Data["TimeStamp"] = timeStamp
		// "<![CDATA["++"]]>"
		paySign := JsapiSign(WxAppId, timeStamp, m2["nonce_str"], "prepay_id="+resp.PrepayId, "MD5", apiKey)
		// paySign := core.Sign(m1, apiKey, nil)

		ctx.Data["PaySign"] = paySign

		// fmt.Println(req.NonceStr, resp.TradeType, paySign)
		// ctx.Redirect("/allfileinfo", 301)
		ctx.HTML(200, "wxpay")
	} else {
		errinfo = "Error op!"
		gotourl = WebSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}

}

func WxPayRelHandler(ctx *macaron.Context, log *log.Logger) {
	var w http.ResponseWriter
	w = ctx.Resp
	wxpayServer.ServeHTTP(w, ctx.Req.Request, nil)
}

func wxPayRelHandler(ctx *core.Context) {

	// fmt.Println("Ctx", ctx)
	reqmap := ctx.Msg
	payinfo := new(models.PayInfo)
	payinfo.Wid = reqmap["appid"]
	payinfo.Attach = reqmap["attach"]
	payinfo.BankType = reqmap["bank_type"]
	cashfee := reqmap["cash_fee"]
	payinfo.CashFee, err = strconv.ParseInt(cashfee, 10, 64)
	if err != nil {
		log.Println("PayFileInfo CashFee Error:", err)
	}

	payinfo.DeviceInfo = reqmap["device_info"]
	payinfo.FeeType = reqmap["fee_type"]
	payinfo.IsSubscribe = reqmap["is_subscribe"]
	payinfo.MchId = reqmap["mch_id"]
	payinfo.OpenId = reqmap["openid"]
	payinfo.OutTradeNo = reqmap["out_trade_no"]
	payinfo.TransactionId = reqmap["transaction_id"]
	payinfo.PayTime = reqmap["time_end"]
	payinfo.PayOk = true
	payinfo.PrintOk = false
	totalfee := reqmap["total_fee"]
	payinfo.TotalPay, err = strconv.ParseInt(totalfee, 10, 64)
	if err != nil {
		log.Println("PayFileInfo TotalFee Error : ", err)
	}
	payinfo.TradeType = reqmap["trade_type"]

	if reqmap["return_code"] == "SUCCESS" {
		m := make(map[string]string, 2)

		if reqmap["result_code"] == "SUCCESS" {
			m["return_code"] = "SUCCESS"
			m["return_msq"] = "OK"
			ctx.Response(m)
		} else {
			m["return_code"] = "FAIL"
			m["return_msq"] = "NO"
			ctx.Response(m)
			return
		}
	}

	// fmt.Println("PayInfo:", payinfo, "\nReqMAp : ", reqmap)
	jud, err := models.JudePayInfo(payinfo)
	if err != nil {
		log.Println("GetPayInfo Error : ", err)
	}
	if jud {
		log.Println("Exist PayInfo!")
		return
	}
	err = models.AddPayInfo(payinfo)
	if err != nil {
		log.Println("Add PayInfo Error : ", err)
		return
	}
	err = models.UpdatePayFileInfo(payinfo.OutTradeNo)
	if err != nil {
		log.Println("Pay UpdaTe Get FileInfo Error : ", err)
		return
	}
	if err != nil {
		log.Println("UpdatePayFileInfo Error : ", err)
	}
	getuser, err := models.GetUser(payinfo.OpenId)
	getuser.TotalConsumption = getuser.TotalConsumption + payinfo.CashFee
	if payinfo.IsSubscribe == "Y" {
		getuser.IsSubscriber = 1
	} else {
		getuser.IsSubscriber = 0
	}
	err = models.UpdateUserInfo(getuser)
	if err != nil {
		log.Println("{Pay Info UserInfo Update Error : ", err)
	}
}

func UnifiedOrder2(clt *core.Client, req *pay.UnifiedOrderRequest) (resp *pay.UnifiedOrderResponse, m1, m2 map[string]string, err error) {
	m1 = make(map[string]string, 20)
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

	// fmt.Println(m1)
	m2, err = UnifiedOrder(clt, m1)
	// fmt.Println(m2, "This is resp.", err)
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

// jssdk 支付签名, signType 只支持 "MD5", "SHA1", 传入其他的值会 panic.
func JsapiSign(appId, timeStamp, nonceStr, packageStr, signType string, apiKey string) string {
	var h hash.Hash
	switch signType {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	default:
		panic("unsupported signType")
	}

	bufw := bufio.NewWriterSize(h, 128)
	// appId
	// nonceStr
	// package
	// signType
	// timeStamp
	bufw.WriteString("appId=")
	bufw.WriteString(appId)
	bufw.WriteString("&nonceStr=")
	bufw.WriteString(nonceStr)
	bufw.WriteString("&package=")
	bufw.WriteString(packageStr)
	bufw.WriteString("&signType=")
	bufw.WriteString(signType)
	bufw.WriteString("&timeStamp=")
	bufw.WriteString(timeStamp)
	bufw.WriteString("&key=")
	bufw.WriteString(apiKey)

	bufw.Flush()
	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))

	return string(bytes.ToUpper(signature))
}
