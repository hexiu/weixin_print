package controller

import (
	"fmt"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"weixin_dayin/models"
)

var filetype string = "doc"
var printnum int = 1

// var newuser *models.User
var DelFile string = "true"

func FileHandler(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	fmt.Println("this is test. ")
	sid := ctx.GetCookie(cookieName)
	ctx.Data["IsFileUpload"] = true

	if sess.ID() == sid {
		ctx.HTML(200, "index")
	} else {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
}

func UploadHandler(ctx *macaron.Context, sess session.Store, log *log.Logger) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	filetype = ctx.Req.FormValue("filetype")
	printNum := ctx.Req.FormValue("printnum")
	if len(printNum) != 0 {
		printnum, err = strconv.Atoi(printNum)
		if err != nil {
			log.Println(err)
		}
	}

	_, fh, err := ctx.GetFile("file")
	if err != nil {
		log.Println(err)
	}

	var attachment string

	if fh != nil {
		//上传文件
		attachmentFilename := fh.Filename
		attachment = attachmentFilename
		getuser, err := models.GetUser(fmt.Sprintf("%v", sess.Get("openid")))
		if err != nil {
			log.Println(err)
		}
		filelists, _ := models.GetNotPrintFileInfo(getuser.OpenId)
		CurFileNum := len(filelists)
		if CurFileNum > 20 {
			errinfo = "您未付款的打印文件过多，已超过预设值，请对文件付款或者删除之后再次上传，谢谢您的配合"
			gotourl = WebSiteUrl + "/filelist"
			ctx.Redirect("/errorinfo", 301)
		}
		fmt.Println(attachment, path.Join(getuser.FileSavePath, attachment))
		err = ctx.SaveToFile("file", path.Join(getuser.FileSavePath, attachment)) //可以使用相对路径
		fmt.Println(err)
		if err != nil {
			err.Error()
		}
		fileinfo := new(models.FileInfo)

		fileinfo.OpenId = userinfo.OpenId
		fileinfo.FileName = attachmentFilename
		fileinfo.FilePayInfo = false
		fileinfo.FileWherePath = "local"
		fileinfo.FileUrl = webSiteUrl + "/" + getuser.OpenId + "/" + attachmentFilename
		fileinfo.FileType = filetype
		fileinfo.PrintNum = printnum
		fileinfo.FileUploadTime = time.Now().Unix()
		fileinfo.Flag = 0
		// fileinfo.Uid = getuser.Id //
		fileinfo.FileUploadDate = time.Now().String()[0:16]
		if filetype == "image" {
			fileinfo.Fee = 1.00
		} else {
			// fileinfo.Fee =
		}
		err = models.AddFileInfo(fileinfo)
		if err != nil {
			log.Println("Controller FileUpload Handler AddFileInfo Error : ", err)
		}

		getuser.UploadFileNum = getuser.UploadFileNum + 1
		err = models.UpdateUserInfo(getuser)
		if err == nil {
			log.Println(err)
		}

		ctx.Redirect("/file", 301)
	}
}

func ShowAllFileInfo(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	ctx.Data["IsAllFileInfo"] = true
	fmt.Println("User Flag:", sess.Get("openid"))
	getuser, err := models.GetUser(fmt.Sprintf("%v", sess.Get("openid")))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("UserInfo:", getuser)
	fileinfolist, err := models.GetAllFileInfo(getuser.OpenId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fileinfo:", fileinfolist)
	ctx.Data["FileInfoList"] = fileinfolist
	ctx.HTML(200, "allfilelist")
}

func DelFileHandler(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}

	op := ctx.Req.FormValue("op")
	id := ctx.Req.FormValue("id")
	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	if op == "del" {
		fileinfo, err := models.GetFileInfo(Id)
		if err != nil {
			log.Println(err)
		}
		filepath := "attachment" + fileinfo.OpenId + fileinfo.FileName
		if DelFile == "true" {
			err = os.Remove(filepath)
			if err != nil {
				log.Println(err)
			}
		}
		fileinfo.Flag = 1

		models.UpdateFileInfo(fileinfo)
		ctx.Redirect("/allfileinfo", 301)
	}
	errinfo = "文件删除失败"
	ctx.Redirect("/errorinfo", 301)
}
