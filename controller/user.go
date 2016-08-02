package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/user"

	"log"
	"time"
	"weixin_dayin/models"
	"weixin_dayin/modules/initConf"
)

const (
	Lang = "zh_CN"
)

var lang string = Lang

func init() {
	conf, err = initConf.InitConf()
	if err != nil {
		log.Println(err)
	}
	ok, err := conf.GetValue("UserInfo", "Lang")
	if err != nil {
		log.Println(err)
		return
	}
	lang = ok

}

func UserAddFromWeiXinHandler(ctx *core.Context) {
	// newuser := models.NewUser()
	newuser := new(models.User)
	msg := ctx.MixedMsg

	newuser.Wid = msg.ToUserName
	newuser.OpenId = msg.FromUserName
	newuser.CreateTime = msg.CreateTime
	newuser.Flag = 0
	newuser.UpdateTime = time.Now().Unix()
	newuser.PrintFileNum = 0
	newuser.NotPrintFile = 0
	newuser.UploadFileNum = 0
	newuser.TotalConsumption = 0

	userinfo, err := userUpdateFromWeiXin(newuser.OpenId, lang)

	newuser.Nickname = userinfo.Nickname
	newuser.Sex = userinfo.Sex
	newuser.Country = userinfo.Country
	newuser.City = userinfo.City
	newuser.Language = userinfo.Language
	newuser.IsSubscriber = userinfo.IsSubscriber
	newuser.Headimgurl = userinfo.HeadImageURL
	newuser.UnionId = userinfo.UnionId
	newuser.Province = userinfo.Province
	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
	}

	err = models.AddUser(newuser)
	if err != nil {
		log.Println("Controller UserHander AddUser Error : ", err)
	}
	// log.Println(newuser)

}

func UserAddFromWebHandler() {
	newuser := models.NewUser()
	newuser.OpenId = userinfo.OpenId
	newuser.CreateTime = time.Now().Unix()
	newuser.Flag = 0
	newuser.UpdateTime = time.Now().Unix()
	newuser.PrintFileNum = 0
	newuser.NotPrintFile = 0
	newuser.UploadFileNum = 0
	newuser.TotalConsumption = 0
	newuser.Nickname = userinfo.Nickname
	newuser.Sex = userinfo.Sex
	newuser.Country = userinfo.Country
	newuser.City = userinfo.City
	newuser.IsSubscriber = 0
	newuser.UnionId = userinfo.UnionId
	newuser.Province = userinfo.Province

	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
	}

	err = models.AddUser(newuser)
	if err != nil {
		log.Println("Controller UserHander AddUser Error : ", err)
	}
	// log.Println(newuser)

}

func userUpdateFromWeiXin(openId string, lang string) (userinfo *user.UserInfo, err error) {
	userinfo, err = user.Get(Client, openId, lang)
	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
		return nil, err
	}
	return userinfo, nil
}

func GetUser(openid string) (err error) {
	getuser, err := models.GetUser(openid)
	if err != nil {
		return err
	}
	fmt.Println(getuser)
	return nil
}

func ExistUser(openid string) bool {
	getuser := &models.User{
		OpenId: openid,
	}
	has, err := models.JudgeUser(getuser)
	if err != nil {
		log.Println(err)
	}
	return has
}

func UserHasFromWeb(openid string) bool {
	getuser, err := models.GetUser(openid)
	if err != nil {
		log.Println(err)
		return false
	}
	if len(getuser.Wid) == 0 {
		return true
	}
	return false
}

func UserUpdateFormWeiXin(ctx *core.Context) {
	msg := ctx.MixedMsg
	getuser, err := models.GetUser(msg.FromUserName)
	if err != nil {
		log.Println(err)
		return
	}
	getuser.NearUpdateFileTime = time.Now().Unix()
	getuser.CreateTime = msg.CreateTime
	getuser.Wid = msg.ToUserName
	getuser.IsSubscriber = 1
	err = models.UpdateUserInfo(getuser)
	if err != nil {
		log.Println(err)
	}
}

func ChageSubscribe(openid string, status int) {
	getuser, err := models.GetUser(openid)
	if err != nil {
		log.Println(err)
	}
	getuser.IsSubscriber = status
	err = models.UpdateUserInfo(getuser)
	if err != nil {
		log.Println(err)
	}
}
