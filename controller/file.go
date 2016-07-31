package controller

import (
	"fmt"
	// "github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"os"
	"path"
	"weixin_dayin/models"
)

var newuser *models.User

func FileHandler(ctx *macaron.Context, log *log.Logger) {
	// sid := ctx.GetCookie("WxPaySession")
	// Sess, err := SessionStorage.Start(ctx)
	// if err != nil {
	// 	log.Println(err)
	// // }
	// if sid == Sess.ID() {
	ctx.HTML(200, "index")
	// } else {
	// 	ctx.Redirect("/", 301)
	// }
}

func UploadHandler(ctx *macaron.Context) string {
	// sid := ctx.GetCookie("WxPaySession")
	// Sess, err := SessionStorage.Start(ctx)
	// if err != nil {
	// 	log.Println(err)
	// 	return "<h1>你还没有登录哦！</h1> <br> <a href=\"" + "http://wxpay.jaxiu.cn" + "\">点这里跳回主页</a>"
	// }
	// if sid != Sess.ID() {
	// 	ctx.Redirect("/", 301)
	// }

	// fmt.Println(uf.TextUpload.Filename)
	fmt.Println("test")
	_, fh, err := ctx.GetFile("file")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	var attachment string
	if fh != nil {
		//上传文件
		attachmentFilename := fh.Filename
		// filepath,err := PrivideDir(user.OpenId)
		attachment = attachmentFilename
		// ctx.Info(attachment)
		fmt.Println(attachment, path.Join("attachment", attachment))
		err := ctx.SaveToFile("file", path.Join("attachment", attachment)) //可以使用相对路径
		fmt.Println(err)
		// filename : tmp.go
		// attachement/tmp.go
		if err != nil {
			err.Error()
		}

	}

	// ... 您可以在这里读取上传的文件内容

	return "thanks"
}

// func GetUser(openid string, wid string) (err error) {
// 	newuser, err = models.GetUser(openid, wid)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func PrivideUserDir(username string) (err error) {
	filepath := "attachment/" + username
	err = os.Mkdir(filepath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func PrivideDir(username string, dataTime string) (filepath string, err error) {
	filepath = "attachment/" + username + "/" + dataTime
	err = os.Mkdir(filepath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
