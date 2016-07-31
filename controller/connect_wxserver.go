package controller

import (
	// "encoding/json"
	"fmt"
	// "github.com/chanxuehong/wechat.v2/mp/base"
	// "github.com/chanxuehong/wechat.v2/mp/core"
	// "github.com/chanxuehong/wechat.v2/mp/menu"
	// "io/ioutil"
	"github.com/Unknwon/goconfig"
	// "net/http"
	// "weixin_dayin/modules/initConf"

	"log"
	// "strconv"
	// "os"
	// "path/filepath"
)

var (
	// 微信对接相关变量
	WxAppId         string = ""
	WxAppSecret     string = ""
	WxOriId         string = ""
	WxToken         string
	WxEncodedAESKey string = ""
)

// var (
// 	msgClient   *core.Client
// 	tokenServer core.AccessTokenServer
// )
var conf *goconfig.ConfigFile

func initconf() {
	if ok, err := conf.GetValue("WeiXin", "WxAppId"); err == nil {
		WxAppId = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxAppSecret"); err == nil {
		WxAppSecret = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxOriId"); err == nil {
		WxOriId = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxToken"); err == nil {
		if ok == "" {
			log.Println("WeiXin Config Error : ", "WxToken can not null")
			return
		} else {
			WxToken = ok
			fmt.Println(WxToken)
		}

	} else {
		log.Println(err)
		return
	}
	if ok, err := conf.GetValue("WeiXin", "WxEncodedAESKey"); err == nil {
		WxEncodedAESKey = ok
	} else {
		log.Println(err)
	}

}
