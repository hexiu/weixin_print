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

func UserAddHandler(ctx *core.Context) {
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
	userinfo, err := userUpdateFromWeiXin(newuser.OpenId, "zh_CN")
	newuser.Nickname = userinfo.Nickname
	newuser.Sex = userinfo.Sex
	newuser.Country = userinfo.Country
	newuser.City = userinfo.City
	newuser.Language = userinfo.Language
	newuser.IsSubscriber = userinfo.IsSubscriber

	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
	}

	err = models.AddUser(newuser)
	if err != nil {
		log.Println("Controller UserHander AddUser Error : ", err)
	}
	log.Println(newuser)

}

func userUpdateFromWeiXin(openId string, lang string) (userinfo *user.UserInfo, err error) {
	userinfo, err = user.Get(Client, openId, lang)
	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
		return nil, err
	}
	return userinfo, nil
}

func GetUser(openid string, wid string) (err error) {
	getuser, err := models.GetUser(openid, wid)
	if err != nil {
		return err
	}
	fmt.Println(getuser)
	return nil
}
