package controller

import (
	"fmt"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	// "os"
	"path"
	"time"
	"weixin_dayin/models"
)

var filetype string = "doc"

var newuser *models.User

func FileHandler(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	fmt.Println("this is test. ")
	sid := ctx.GetCookie(cookieName)

	if sess.ID() == sid {
		ctx.HTML(200, "index")
	} else {
		ctx.Redirect("/", 301)
	}
}

func UploadHandler(ctx *macaron.Context, sess session.Store, log *log.Logger) string {
	sid := ctx.GetCookie(cookieName)

	if sid != sess.ID() {
		return "<h1>你还没有登录哦！</h1> <br> <a href=\"" + "http://wxpay.jaxiu.cn" + "\">点这里跳回主页</a>"
	}
	_, fh, err := ctx.GetFile("file")
	if err != nil {
		return err.Error()
	}

	var attachment string

	if fh != nil {
		//上传文件
		attachmentFilename := fh.Filename
		attachment = attachmentFilename
		getuser, err := models.GetUser(userinfo.OpenId)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(attachment, path.Join(getuser.FileSavePath, attachment))
		err = ctx.SaveToFile("file", path.Join(getuser.FileSavePath, attachment)) //可以使用相对路径
		fmt.Println(err)
		if err != nil {
			err.Error()
		}
		getuser, err = models.GetUser(userinfo.OpenId)
		if err != nil {
			log.Println(err)
		}
		fileinfo := new(models.FileInfo)
		fileinfo.Wid = getuser.Wid
		fileinfo.OpenId = userinfo.OpenId
		fileinfo.FileName = attachmentFilename
		fileinfo.FilePayInfo = false
		fileinfo.FileWherePath = "local"
		fileinfo.FileUrl = webSiteUrl + getuser.FileSavePath + attachmentFilename
		fileinfo.FileType = filetype
		fileinfo.FileUploadTime = time.Now().Unix()
		fileinfo.FileUploadDate = string(time.Now().Year()) + string(time.Now().Month()) + string(time.Now().Day())
		if filetype == "image" {
			fileinfo.Fee = 1.00
		} else {
			// fileinfo.Fee =

		}
		return "update ok"
	}

	// ... 您可以在这里读取上传的文件内容

	return "thanks"
}
