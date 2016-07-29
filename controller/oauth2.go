package controller

import (
	"encoding/json"
	"fmt"
	// "github.com/chanxuehong/rand"
	// "github.com/chanxuehong/sid"
	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/oauth2"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	wxAppId           = "wxa9e9c5e110b8665d"               // 填上自己的参数
	wxAppSecret       = "e9a7f50077fdcbf9b689e629678e244a" // 填上自己的参数
	oauth2RedirectURI = "http://wxpay.jaxiu.cn/page2"      // 填上自己的参数
	oauth2Scope       = "snsapi_userinfo"                  // 填上自己的参数
)

var (
	SessionStorage *session.Manager
	err            error
	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
)

func init() {
	SessionStorage, err = session.NewManager("memory", session.Options{})
	if err != nil {
		log.Println(err)
	}

}

// 建立必要的 session, 然后跳转到授权页面
func Page1Handler(ctx *macaron.Context) {
	fmt.Println("OK............1")
	// ctx.SetCookie("sid", "1")
	sess, err := SessionStorage.Start(ctx)
	state := sess.ID()

	fmt.Println(err)
	fmt.Println(sess, "Ok.......................2")

	// if sess, err := sessionStorage.Start(ctx); err != nil {
	// 	fmt.Println(sess, "Ok.......................2")

	// 	log.Println(err)
	// 	io.WriteString(ctx.Resp, err.Error())
	// } else {
	// 	fmt.Println(sess, "Ok.......................3")

	// }

	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
	log.Println("AuthCodeURL:", AuthCodeURL)
	fmt.Println("*****************************", AuthCodeURL)

	ctx.Redirect(AuthCodeURL, http.StatusFound)

}

// 授权后回调页面
func Page2Handler(ctx *macaron.Context) {
	log.Println(ctx.Req.RequestURI)
	cookie, _ := ctx.Req.Cookie("MacaronSession")
	fmt.Println("cookie:", cookie)
	sess, err := SessionStorage.Start(ctx)

	if err != nil {
		log.Println(err)
		io.WriteString(ctx.Resp, err.Error())
		return
	}

	session := sess.ID()

	// savedState := session.(string)
	// savedState := session.(string) // 一般是要序列化的, 这里保存在内存所以可以这么做

	queryValues, err := url.ParseQuery(ctx.Req.URL.RawQuery)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err)
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

	if session != queryState {
		str := fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", session, queryState)
		io.WriteString(ctx.Resp, str)
		log.Println(str)
		return
	}

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err)
		return
	}
	log.Printf("token: %+v\r\n", token)

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		io.WriteString(ctx.Resp, err.Error())
		log.Println(err)
		return
	}

	json.NewEncoder(ctx.Resp).Encode(userinfo)
	log.Printf("userinfo: %+v\r\n", userinfo)
	return
}

// // 建立必要的 session, 然后跳转到授权页面
// func Page1Handler(w http.ResponseWriter, r *http.Request) {
// 	sid := sid.New()
// 	state := string(rand.NewHex())

// 	if err := sessionStorage.Add(sid, state); err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	cookie := http.Cookie{
// 		Name:     "sid",
// 		Value:    sid,
// 		HttpOnly: true,
// 	}
// 	http.SetCookie(w, &cookie)

// 	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, oauth2Scope, state)
// 	log.Println("AuthCodeURL:", AuthCodeURL)

// 	http.Redirect(w, r, AuthCodeURL, http.StatusFound)
// }

// // 授权后回调页面
// func Page2Handler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.RequestURI)

// 	cookie, err := r.Cookie("sid")
// 	if err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	session, err := sessionStorage.Get(cookie.Value)
// 	if err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	savedState := session.(string) // 一般是要序列化的, 这里保存在内存所以可以这么做

// 	queryValues, err := url.ParseQuery(r.URL.RawQuery)
// 	if err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	code := queryValues.Get("code")
// 	if code == "" {
// 		log.Println("用户禁止授权")
// 		return
// 	}

// 	queryState := queryValues.Get("state")
// 	if queryState == "" {
// 		log.Println("state 参数为空")
// 		return
// 	}
// 	if savedState != queryState {
// 		str := fmt.Sprintf("state 不匹配, session 中的为 %q, url 传递过来的是 %q", savedState, queryState)
// 		io.WriteString(w, str)
// 		log.Println(str)
// 		return
// 	}

// 	oauth2Client := oauth2.Client{
// 		Endpoint: oauth2Endpoint,
// 	}
// 	token, err := oauth2Client.ExchangeToken(code)
// 	if err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}
// 	log.Printf("token: %+v\r\n", token)

// 	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
// 	if err != nil {
// 		io.WriteString(w, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(userinfo)
// 	log.Printf("userinfo: %+v\r\n", userinfo)
// 	return
// }

//
// //
