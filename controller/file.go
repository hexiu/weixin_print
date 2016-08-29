package controller

import (
	"fmt"
	// "github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"weixin_dayin/models"
)

type UploadForm struct {
	Filetype string                `form:"filetype"`
	Printnum string                `form:"printnum"`
	File     *multipart.FileHeader `form:"file"`
}

var filetype string = "doc"
var printnum int = 1

// var newuser *models.User
var DelFile string = "true"

func FileHandler(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	fmt.Println("this is test. ")
	sid := ctx.GetCookie(cookieName)
	ctx.Data["IsFileUpload"] = true

	if sess.ID() == sid {
		ctx.Data["Title"] = "打印文件上传"
		ctx.HTML(200, "index")
	} else {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
}

// func UploadHandler(ctx *macaron.Context, sess session.Store, log *log.Logger, uf UploadForm) {
func UploadHandler(ctx *macaron.Context, sess session.Store, log *log.Logger, uf UploadForm) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	fmt.Println(ctx.Req.Request, ctx.Req.PostForm)

	rel := ctx.Req.Request
	fmt.Println(rel)
	rel.ParseForm()
	test := rel.PostForm.Get("filetype")
	fmt.Println(test, rel.Form, rel)
	// filetype := ctx.Req.Form.Get("filetype")
	filetype := uf.Filetype
	printnum := 1
	// printNum := ctx.Req.Form.Get("printnum")
	printNum := uf.Printnum
	if len(printNum) != 0 {
		printnum, err = strconv.Atoi(printNum)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Form:", ctx.Req.Form)
	fmt.Println("filetype:", filetype, "printnum:", printNum, ctx.Req.Form)
	// fmt.Println("R           E                   L               :", cpt.VerifyReq(ctx.Req))
	// if !cpt.VerifyReq(ctx.Req) {
	// 	errinfo = "验证码错误！"
	// 	gotourl = webSiteUrl
	// 	ctx.Redirect(webSiteUrl+"/file", 301)
	// }
	fmt.Println("\n\n\nFilename:\n ", uf.File.Filename, "\n\n\n\n", uf)
	_, fh, err := ctx.GetFile("file")

	if err != nil {
		log.Println(err)
	}
	var attachment string

	if fh != nil {
		//上传文件
		attachmentFilename := fh.Filename
		attachment = attachmentFilename
		fileNameSplit := strings.Split(attachmentFilename, ".")
		length := len(fileNameSplit)
		fileReName := strconv.Itoa(int(time.Now().Unix())) + "." + fileNameSplit[length-1]

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
		fileinfo.FileReName = fileReName
		fileinfo.FilePayInfo = false
		fileinfo.FileWherePath = "local"
		fileinfo.FileUrl = webSiteUrl + "/" + getuser.OpenId + "/" + fileReName
		fileinfo.FileType = fmt.Sprintf("%v", filetype)
		fileinfo.PrintNum = printnum
		fileinfo.FileUploadTime = time.Now().Unix()
		fileinfo.Flag = 0
		fileinfo.OutTradeNo = strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(time.Now().UnixNano()))
		// fileinfo.Uid = getuser.Id //
		fileinfo.FileUploadDate = time.Now().String()[0:16]
		if filetype == "image" {
			fileinfo.Fee = 100
		} else {
			// fileinfo.Fee =
		}
		err = os.Rename("attachment/"+getuser.OpenId+"/"+attachmentFilename, "attachment/"+getuser.OpenId+"/"+fileReName)
		if err != nil {
			log.Println("Renamefile Error : ", err)
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
		// return "thanks"
	}
}

func ShowAllFileInfo(ctx *macaron.Context, log *log.Logger, sess session.Store) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	ctx.Data["WxPayUrl"] = WebSiteUrl + "/wxpay"
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

func WxPayFileHandler(ctx *macaron.Context, sess session.Store, log *log.Logger) {
	sid := ctx.GetCookie(cookieName)
	if sid != sess.ID() {
		errinfo = "你还没有登录哦！"
		gotourl = webSiteUrl
		ctx.Redirect("/errorinfo", 301)
	}
	ctx.Data["Title"] = "已打印文件信息"
	fmt.Printf("%v", sess.Get("openid"))
	fileinfolist, err := models.GetPrintFileInfo(fmt.Sprintf("%v", sess.Get("openid")))
	if err != nil {
		log.Println("Get WxPayFileInfo Error : ", err)
		errinfo = " 发现错误信息，错误内容已经提交至后台 。 "
		gotourl = WebSiteUrl + "/page1"
		ctx.Redirect("/errorinfo", 301)
		return
	}
	fmt.Println(fileinfolist)
	// length := len(filelistinfo)
	ctx.Data["FileInfoList"] = fileinfolist
	ctx.HTML(200, "payok")
}
