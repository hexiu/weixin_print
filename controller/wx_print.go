package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	// "net/http"
	// "strconv"
	// "gopkg.in/macaron.v1"
	// "strings"
	"os"
	"time"
	"weixin_dayin/models"
	// "weixin_dayin/modules/NetSendPrintMsg"
)

// not use in 2016-09-01
func printStartHandler(ctx *core.Context, event *menu.ClickEvent) {
	respMsg := "*创昕小印* 已经接受到打印请求，请您耐心等待，谢谢您的配合，祝您使用愉快  "
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, respMsg)
	ctx.RawResponse(resp)
	fileinfos, err := models.GetPayNotPrintFileInfo(event.FromUserName, true)
	if err != nil {
		log.Println(err)
	}
	length := len(fileinfos)
	getuser, err := models.GetUser(event.FromUserName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Print Handler user : ", getuser)
	printCode := getuser.PrintCode
	if length == 0 {
		clientQuit(printCode)
	}
	for i := 0; i < length; i++ {
		msg := new(TxMsg)
		msg.FileType = fileinfos[i].FileType
		msg.MsgInfo = "print_start"
		msg.MsgType = "fileinfo"
		msg.OpenId = fileinfos[i].OpenId
		msg.MsgURL = fileinfos[i].FileUrl
		msg.PrintCode = printCode
		msg.PrintNum = int64(length)
		msg.Time = time.Now().Unix()
		// mutex.Lock()
		TxChan <- *msg
		// mutex.Unlock()

		fileinfos[i].FilePrintTime = time.Now().Unix()
		fileinfos[i].FilePrintDate = time.Now().String()[:16]
		if DelFile == "true" && fileinfos[i].FileWherePath == "local" {
			filepath := "attachment" + "/" + fileinfos[i].OpenId + "/" + fileinfos[i].FileName
			err = os.Remove(filepath)
			if err != nil {
				log.Println(err)
			}
			fileinfos[i].Flag = 2
		} else {
			fileinfos[i].Flag = 1
		}
		err := models.UpdateFileInfo(fileinfos[i])
		if err != nil {
			log.Println("Print Handler Update FileInfo PrintInfo Error : ", err)
		}

		fmt.Println(msg)
	}
	clientQuit(printCode)
	fmt.Println("msg send ok ")
}

func clientQuit(printCode string) {
	msg := new(TxMsg)
	msg.FileType = "NULL"
	msg.MsgURL = "NULL"
	msg.OpenId = "NULL"
	msg.PrintCode = printCode
	msg.Time = time.Now().Unix()
	msg.MsgType = "connect"
	msg.PrintNum = 0
	msg.MsgInfo = "quit"
	// mutex.Lock()
	TxChan <- *msg
	// mutex.Unlock()
}

func printCodeHandler(ctx *core.Context, event *menu.ClickEvent) {
	fileinfos, err := models.GetPayNotPrintFileInfo(event.FromUserName, true)
	if err != nil {
		log.Println(err)
	}
	length := len(fileinfos)

	respMsg := fmt.Sprintln("*创昕小印* 已经知道您所在的打印机位置啦，您有%d个文件已付款，将会被打印，请点击菜单“扫描确认”，开始打印   ", length)

	printCode := ctx.MixedMsg.ScanCodeInfo.ScanResult
	fmt.Println("Printcode:", printCode)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, respMsg)
	msg := new(TxMsg)
	msg.FileType = "null"
	msg.MsgInfo = "print_code"
	msg.MsgType = "msg"
	msg.OpenId = event.FromUserName
	msg.MsgURL = "null"
	msg.PrintCode = printCode
	msg.PrintNum = int64(length)
	msg.Time = time.Now().Unix()
	// mutex.Lock()
	TxChan <- *msg
	// mutex.Unlock()
	fmt.Println("msg:", msg)
	ctx.RawResponse(resp)
	getuser, err := models.GetUser(event.FromUserName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(getuser)
	getuser.PrintCode = printCode
	fmt.Println("PrintCode Handler user : ", getuser)
	models.UpdateUserInfo(getuser)

	printCode = getuser.PrintCode
	if length == 0 {
		clientQuit(printCode)
	}
	for i := 0; i < length; i++ {
		msg := new(TxMsg)
		msg.FileType = fileinfos[i].FileType
		msg.MsgInfo = "print_start"
		msg.MsgType = "fileinfo"
		msg.OpenId = fileinfos[i].OpenId
		msg.MsgURL = fileinfos[i].FileUrl
		msg.PrintCode = printCode
		msg.PrintNum = int64(length)
		msg.Time = time.Now().Unix()
		// mutex.Lock()
		TxChan <- *msg
		// mutex.Unlock()

		fileinfos[i].FilePrintTime = time.Now().Unix()
		fileinfos[i].FilePrintDate = time.Now().String()[:16]
		if DelFile == "true" && fileinfos[i].FileWherePath == "local" {
			filepath := "attachment" + "/" + fileinfos[i].OpenId + "/" + fileinfos[i].FileName
			err = os.Remove(filepath)
			if err != nil {
				log.Println(err)
			}
			fileinfos[i].Flag = 2
		} else {
			fileinfos[i].Flag = 1
		}
		err := models.UpdateFileInfo(fileinfos[i])
		if err != nil {
			log.Println("Print Handler Update FileInfo PrintInfo Error : ", err)
		}

		fmt.Println(msg)
	}
	clientQuit(printCode)
}

func printOkHandler(ctx *core.Context, event *menu.ClickEvent) {
	respMsg := "*创昕小印* 正在为您获取您已付款的打印文件，准备打印中……，让您久等了   "
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, respMsg)
	ctx.RawResponse(resp)
}
