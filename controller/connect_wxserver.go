package controller

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"strconv"
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

	//================================================================

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

	//=======================================================================

	if ok, err := conf.GetValue("Session", "Provider"); err == nil {
		provider = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "ProviderConfig"); err == nil {
		providerConfig = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "CookieName"); err == nil {
		cookieName = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "CookiePath"); err == nil {
		cookiePath = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "Gclifetime"); err == nil {
		log.Println(err)
		rel, err := strconv.Atoi(ok)
		if err != nil {
			log.Println(err)
		}
		gclifetime = int64(rel)
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "Maxlifetime"); err == nil {
		rel, err := strconv.Atoi(ok)
		if err != nil {
			log.Println(err)
		}
		maxlifetime = int64(rel)
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "Secure"); err == nil {
		if ok == "true" {
			secure = true
		} else {
			secure = false
		}
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "CookieLifeTime"); err == nil {
		rel, err := strconv.Atoi(ok)
		if err != nil {
			log.Println(err)
		}
		cookieLifeTime = rel
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "Domain"); err == nil {
		domain = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "IDLength"); err == nil {
		rel, err := strconv.Atoi(ok)
		if err != nil {
			log.Println(err)
		}
		iDLength = rel
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Session", "Section"); err == nil {
		section = ok
	} else {
		log.Println(err)
	}

	if ok, err := conf.GetValue("Server", "WebSiteUrl"); err == nil {
		webSiteUrl = ok
	} else {
		log.Println(err)
	}

	//=======================================================================

	if ok, err := conf.GetValue("Oauth2", "Oauth2Url"); err == nil {
		oauth2Url = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("Oauth2", "Oauth2RedirectURI"); err == nil {
		Oauth2RedirectURI = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("Oauth2", "Oauth2Scope"); err == nil {
		Oauth2Scope = ok
	} else {
		log.Println(err)
	}

	//=====================================================================

	if ok, err := conf.GetValue("WeiXinMsg", "SubScribeMsgBack"); err == nil {
		SubScribeMsgBack = ok
	} else {
		log.Println(err)
	}

}
