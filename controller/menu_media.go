package controller

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	"github.com/chanxuehong/wechat.v2/mp/core"
	// "github.com/chanxuehong/wechat.v2/mp/menu"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	// "net/http"
	// "strconv"
	"weixin_dayin/models"
	// "weixin_connect/modules/initConf"
)

func AddMediaInfo(ctx *core.Context) (err error) {
	newmediafile := new(models.FileInfo)
	mediamsg := ctx.MixedMsg
	if mediamsg.MsgType == "image" {
		newmediafile.Wid = mediamsg.ToUserName
		newmediafile.OpenId = mediamsg.FromUserName
		newmediafile.FileUrl = mediamsg.PicURL
		newmediafile.FileUploadTime = mediamsg.CreateTime
		newmediafile.MediaId = mediamsg.MediaId
		newmediafile.MsgId = mediamsg.MsgId
	}
	err = models.AddImageInfo(newmediafile)
	if err != nil {
		log.Println("Controller Menu_Media Handler AddImageIndo Error : ", err)
	}
	return nil
}
