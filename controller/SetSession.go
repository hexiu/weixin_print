package controller

import (
	// "encoding/json"
	"fmt"
	// "github.com/chanxuehong/rand"
	// "github.com/chanxuehong/sid"
	// mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/mp/core"
	// "github.com/chanxuehong/wechat.v2/oauth2"
	// "github.com/go-macaron/session"
	// "gopkg.in/macaron.v1"
	// "io"
	"log"
	"net/http"
	// "net/url"
	// "time"
	"strconv"
	"weixin_dayin/modules/initConf"
	// "weixin_dayin/models"
)

const (
	// 提供器的名称，默认为 "memory"
	Provider = "memory"
	// 提供器的配置，根据提供器而不同
	ProviderConfig = ""
	// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
	CookieName = "WxPaySession"
	// Cookie 储存路径，默认为 "/"
	CookiePath = "/"
	// GC 执行时间间隔，默认为 3600 秒
	Gclifetime int64 = 3600
	// 最大生存时间，默认和 GC 执行时间间隔相同
	Maxlifetime int64 = 3600
	// 仅限使用 HTTPS，默认为 false
	Secure = false
	// Cookie 生存时间，默认为 0 秒
	CookieLifeTime = 0
	// Cookie 储存域名，默认为空
	Domain = ""
	// 会话 ID 长度，默认为 16 位
	IDLength = 16
	// 配置分区名称，默认为 "session"
	Section = "session"
)

var (
	// 提供器的名称，默认为 "memory"
	provider = "memory"
	// 提供器的配置，根据提供器而不同
	providerConfig = ProviderConfig
	// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
	cookieName = CookieName
	// Cookie 储存路径，默认为 "/"
	cookiePath = CookiePath
	// GC 执行时间间隔，默认为 3600 秒
	gclifetime int64 = Gclifetime
	// 最大生存时间，默认和 GC 执行时间间隔相同
	maxlifetime int64 = Maxlifetime
	// 仅限使用 HTTPS，默认为 false
	secure = Secure
	// Cookie 生存时间，默认为 0 秒
	cookieLifeTime = CookieLifeTime
	// Cookie 储存域名，默认为空
	domain = Domain
	// 会话 ID 长度，默认为 16 位
	iDLength = IDLength
	// 配置分区名称，默认为 "session"
	section = Section
)

var (
	err error
)

func init() {
	conf, err = initConf.InitConf()
	if err != nil {
		log.Println(err)
	}
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

	initconf()

	TokenServer = core.NewDefaultAccessTokenServer(WxAppId, WxAppSecret, nil)
	Client = core.NewClient(TokenServer, http.DefaultClient)
	fmt.Println(WxAppSecret, WxAppId, TokenServer)

	// SessionStorage, err = session.NewManager("memory", session.Options{})
	// fmt.Println("Test:", SessionStorage, err)
	// if err != nil {
	// 	log.Println(err)
	// }

	// SessionStorage, err = session.NewManager("memory", session.Options{
	// 	Provider: "memory",
	// 	// 提供器的配置，根据提供器而不同
	// 	// ProviderConfig: providerConfig,
	// 	// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
	// 	CookieName: cookieName,
	// 	// Cookie 储存路径，默认为 "/"
	// 	CookiePath: cookiePath,
	// 	// GC 执行时间间隔，默认为 3600 秒
	// 	Gclifetime: gclifetime,
	// 	// 最大生存时间，默认和 GC 执行时间间隔相同
	// 	Maxlifetime: maxlifetime,
	// 	// 仅限使用 HTTPS，默认为 false
	// 	Secure: secure,
	// 	// Cookie 生存时间，默认为 0 秒
	// 	CookieLifeTime: cookieLifeTime,
	// 	// Cookie 储存域名，默认为空
	// 	Domain: domain,
	// 	// 会话 ID 长度，默认为 16 位
	// 	IDLength: iDLength,
	// 	// 配置分区名称，默认为 "session"
	// 	Section: section,
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println("SessionStorage :", SessionStorage, SessionStorage.Count())
	// go SessionStorage.GC()
	// SessionMange = SessionStorage
	// sess.GC()
}

func test() {
	fmt.Println("...")
}
