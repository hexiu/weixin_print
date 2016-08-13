package controller

import (
	// "encoding/json"
	"fmt"
	// "github.com/chanxuehong/rand"
	// "github.com/chanxuehong/sid"
	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	// "github.com/chanxuehong/wechat.v2/mp/user"
	"github.com/chanxuehong/wechat.v2/oauth2"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"io"
	"log"
	"net/http"
	"net/url"
	// "time"
	// "weixin_dayin/models"
	// "weixin_dayin/modules/initConf"
)

const (
	oauth2RedirectURI = "http://wxpay.jaxiu.cn/page2" // oauth2 认证通过跳转页面
	oauth2Scope       = "snsapi_userinfo"             // oauth2 指定获取用户信息的等级
	WebSiteUrl        = "http://wxpay.jaxiu.cn"
	Oauth2Url         = "http://wxpay.jaxiu.cn"
)

var (
	// SessionStorage *session.Manager
	// SessionMange   *session.MemProvider
	// sess           session.Store                                                   // 全局Session管理器
	userinfo       *mpoauth2.UserInfo //用户oauth2认证登录之后，抓取的用户信息
	oauth2Endpoint oauth2.Endpoint    //oauth2客户端认证所需的信息
)

var (
	webSiteUrl        string = WebSiteUrl
	oauth2Url         string = Oauth2Url
	Oauth2RedirectURI        = oauth2RedirectURI // oauth2 认证通过跳转页面
	Oauth2Scope              = oauth2Scope       // oauth2 指定获取用户信息的等级
)

// 建立必要的 session, 然后跳转到授权页面
func Page1Handler(ctx *macaron.Context, sess session.Store) {
	// fmt.Println(ctx.GetCookie(cookieName), cookieName, "This is Page1 is Cookie")

	if err != nil {
		log.Println(err)
	}
	// fmt.Println(ctx.Req.Cookies(), sess.Count())
	ctx.SetCookie(CookieName, sess.ID())
	state := sess.ID()
	// fmt.Println(ctx.Req.Cookies(), sess.Count())

	AuthCodeURL := mpoauth2.AuthCodeURL(WxAppId, oauth2RedirectURI, oauth2Scope, state)
	log.Println("AuthCodeURL:", AuthCodeURL)

	ctx.Redirect(AuthCodeURL, http.StatusFound)
}

// func PortalPCHandler(ctx *macaron.Context) {
// 	sess, err := SessionStorage.Start(ctx)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	state := sess.ID()
// 	AuthCodeURL := mpoauth2.AuthCodeURL(WxAppId, oauth2RedirectURI, oauth2Scope, state)
// 	log.Println("AuthCodeURL:", AuthCodeURL)
// 	ctx.Redirect(AuthCodeURL, http.StatusFound)
// }

// 授权后回调页面
func GetWxInfoHandler(ctx *macaron.Context, sess session.Store) {
	fmt.Println(ctx.GetCookie(cookieName), cookieName, "This is Page2 is Cookie")

	// fmt.Println("SessionStorage", SessionStorage, SessionStorage.Count())
	log.Println(ctx.Req.RequestURI)
	// cookie, err := ctx.Req.Cookie(CookieName)
	// fmt.Println("cookie:", cookieName, cookie)
	// fmt.Println(ctx.Req.Cookies(), sess.Count())
	// fmt.Println(cookieName, " : ", sess.Get(cookieName))
	// fmt.Println("SessionStorage", sess.Count(), sess.ID())

	// state = sess.ID()

	// fmt.Println("Page2:", sess)
	sessionid := ctx.GetCookie(CookieName)
	// fmt.Println(sessionid, sess.Count())

	queryValues, err := url.ParseQuery(ctx.Req.URL.RawQuery)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err, "QueryValues")
		return
	}
	code := queryValues.Get("code")
	if code == "" {
		log.Println("用户禁止授权")
		return
	}
	queryState := queryValues.Get("state")
	if queryState == "" {
		log.Println("state 参数为空")
		return
	}
	if sessionid != queryState {
		str := fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", sessionid, queryState)
		io.WriteString(ctx.Resp, str)
		log.Println(str)
		return
	}
	// fmt.Println(queryState, sessionid)
	oauth2Endpoint := mpoauth2.NewEndpoint(WxAppId, WxAppSecret)
	oauth2Client := &oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err)
		return
	}
	log.Printf("token: %+v\r\n", token)

	userinfo, err = mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err)
		return
	}
	// json.NewEncoder(ctx.Resp).Encode(userinfo)
	log.Printf("userinfo: %+v\r\n", userinfo)
	if !ExistUser(userinfo.OpenId) {
		UserAddFromWebHandler()
	}
	sess.Set("openid", userinfo.OpenId)
	ctx.Redirect("/file", 301)
	return
}

//
// func UserAddHandler(openid string) (err error) {
// 	// newuser := models.NewUser()

// 	ok, err := JudgeUser(openid, "")
// 	if ok == false && err != nil {
// 		return err
// 	}
// 	if ok == true {
// 		return nil
// 	} else {
// 		newuser := new(models.User)
// 		newuser.OpenId = userinfo.OpenId
// 		newuser.CreateTime = time.Now().Unix()
// 		newuser.Flag = 0
// 		newuser.UpdateTime = time.Now().Unix()
// 		newuser.PrintFileNum = 0
// 		newuser.NotPrintFile = 0
// 		newuser.UploadFileNum = 0
// 		newuser.TotalConsumption = 0
// 		userinfoUpdate, err := userUpdateFromWeiXin(newuser.OpenId, "zh_CN")

// 		if err != nil {
// 			log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
// 		}
// 		newuser.Nickname = userinfoUpdate.Nickname
// 		newuser.Sex = userinfoUpdate.Sex
// 		newuser.Country = userinfoUpdate.Country
// 		newuser.City = userinfoUpdate.City
// 		newuser.Language = userinfoUpdate.Language

// 		err = models.AddUser(newuser)
// 		if err != nil {
// 			log.Println("Controller UserHander AddUser Error : ", err)
// 		}
// 		log.Println(newuser)
// 	}
// 	return nil
// }

// // 网站用户添加模块
// func userUpdateFromWeiXin(openId string, lang string) (userinfoUpdate *user.UserInfo, err error) {
// 	userinfoUpdate = new(user.UserInfo)
// 	userinfoUpdate, err = user.Get(Client, openId, lang)
// 	if err != nil {
// 		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
// 		return nil, err
// 	}
// 	return userinfoUpdate, nil
// }

// func JudgeUser(openid string, wid string) (has bool, err error) {
// 	newuser := &models.User{
// 		OpenId: openid,
// 		Wid:    wid,
// 	}
// 	has, err = models.JudgeUser(newuser)
// 	if err != nil && has == false {
// 		return false, err
// 	}
// 	return true, nil
// }
