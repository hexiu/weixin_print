package controller

import (
	"fmt"
	// "github.com/Unknwon/goconfig"
	"github.com/chanxuehong/wechat.v2/mp/core"
	// "github.com/chanxuehong/wechat.v2/mp/menu"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	// "net/http"
	// "strconv"
	"time"
	"weixin_dayin/models"
	// "weixin_connect/modules/initConf"
)

func AddMediaInfo(ctx *core.Context) (err error) {
	newmediafile := new(models.FileInfo)
	mediamsg := ctx.MixedMsg
	fmt.Println(mediamsg.FromUserName, ctx.MixedMsg, "          ", ctx)
	getuser, err := models.GetUser(mediamsg.FromUserName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(getuser)
	if mediamsg.MsgType == "image" {
		newmediafile.Wid = mediamsg.ToUserName
		newmediafile.OpenId = mediamsg.FromUserName
		newmediafile.FileWherePath = "wx"
		newmediafile.Fee = 1.00
		newmediafile.FilePayInfo = false
		newmediafile.FileUploadDate = time.Now().String()[0:16]
		newmediafile.FileUrl = mediamsg.PicURL
		newmediafile.FileUploadTime = mediamsg.CreateTime
		newmediafile.MediaId = mediamsg.MediaId
		newmediafile.MsgId = mediamsg.MsgId
		newmediafile.Flag = 0
		newmediafile.FileType = "image"
		newmediafile.PrintNum = 1
		// newmediafile.Id = getuser.Id
	}
	err = models.AddFileInfo(newmediafile)
	if err != nil {
		log.Println("Controller Menu_Media Handler AddImageIndo Error : ", err)
	}
	return nil
}
