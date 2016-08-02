package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"os"
	"strconv"
	"weixin_dayin/controller"
	"weixin_dayin/models"
	"weixin_dayin/modules/NetSendPrintMsg"
	"weixin_dayin/modules/initConf"
)

// 定义默认的端口
const (
	Port int = 8080
)

// 初始化默认端口
var (
	port int = Port
)

// 定义全局配置信息接口
var conf *goconfig.ConfigFile

const (
	// 提供器的名称，默认为 "memory"
	Provider = "memory"
	// 提供器的配置，根据提供器而不同
	ProviderConfig = ""
	// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
	CookieName = "MacaronSession"
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

// init Database  &  init Config file
func init() {
	err := models.RegisterDB()
	if err != nil {
		fmt.Println("Error : ", err)
	}
	conf, err = initConf.InitConf()
	initconf()
	go NetSendPrintMsg.StartServer()

}

func main() {
	m := macaron.Classic()

	//Register middle key
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner(session.Options{
		Provider: provider,
		// 提供器的配置，根据提供器而不同
		// ProviderConfig: providerConfig,
		// 用于存放会话 ID 的 Cookie 名称，默认为 "MacaronSession"
		CookieName: cookieName,
		// Cookie 储存路径，默认为 "/"
		CookiePath: cookiePath,
		// GC 执行时间间隔，默认为 3600 秒
		Gclifetime: gclifetime,
		// 最大生存时间，默认和 GC 执行时间间隔相同
		Maxlifetime: maxlifetime,
		// 仅限使用 HTTPS，默认为 false
		Secure: secure,
		// Cookie 生存时间，默认为 0 秒
		CookieLifeTime: cookieLifeTime,
		// Cookie 储存域名，默认为空
		// Domain: domain,
		// 会话 ID 长度，默认为 16 位
		IDLength: iDLength,
		// 配置分区名称，默认为 "session"
		Section: section,
	}))

	// fmt.Println(session.Options{
	// 	Provider: "memory",
	// 	// 提供器的配置，根据提供器而不同
	// 	ProviderConfig: providerConfig,
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
	// 	// Domain: domain,
	// 	// 会话 ID 长度，默认为 16 位
	// 	IDLength: iDLength,
	// 	// 配置分区名称，默认为 "session"
	// 	Section: section})

	// Router info
	m.Get("/page1", controller.Page1Handler)
	m.Get("/page2", controller.GetWxInfoHandler)
	m.Get("/", controller.HomeHandler)
	m.Get("/file", controller.FileHandler)
	m.Post("/fileup", controller.UploadHandler)
	m.Get("/wx_callback", controller.WxCallbackHandler)
	m.Post("/wx_callback", controller.WxCallbackHandler)

	err := os.Mkdir("attachment", os.ModePerm)
	if err != nil {
		log.Println("Create Directory Error : ", err)
	}

	m.Run(port)
}

// 初始化主机的配置文件
func initconf() {

	// 获取配置文件配置的服务监听端口
	if ok := conf.MustInt("Server", "ListenPort"); ok != 0 {
		port = ok
	}

	// Session模块
	// 获取Session相关配置 详情查看这里：https://go-macaron.com/docs/middlewares/session
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
}
