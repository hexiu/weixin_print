package models

import (
// "fmt"
// "time"
)

type PayInfo struct {
	Id         int64
	Uid        int64
	Fid        int64
	Wid        string
	OpenId     string `xorm:"index"`
	TotalPay   int64
	PayTime    string `xorm:"index"`
	CreateTime int64  `xorm:"index"`
	// TimeStart     string //" 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则"
	// TimeExpire    string //"订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则"
	// NotifyURL     string // "接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。"
	TradeType     string // "取值如下：JSAPI，NATIVE，APP，详细说明见参数规定"
	Attach        string //"附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据"
	FeeType       string // "符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型"
	PrintOk       bool
	PayOk         bool
	TransactionId string
	OutTradeNo    string
	BankType      string
	CashFee       int64

	DeviceInfo  string
	IsSubscribe string
	MchId       string
}

func AddPayInfo(payinfo *PayInfo) (err error) {
	connectDB()
	_, err = engine.Insert(payinfo)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func JudePayInfo(payinfo *PayInfo) (has bool, err error) {
	connectDB()
	has, err = engine.Get(payinfo)
	// mt.Println(user, "judge user")
	if err != nil {
		return false, err
	}
	// fmt.Println(user, "judge user")
	defer engine.Close()
	return has, nil

}
