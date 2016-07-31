package controller

import (
	"encoding/json"
	"fmt"
	// "github.com/chanxuehong/wechat.v2/mp/base"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"io/ioutil"
	"net/http"
	// "weixin_dayin/modules/initConf"

	// "log"
	"os"
	"path/filepath"
)

var (
	TokenServer core.AccessTokenServer
	Client      *core.Client
)

// var (
// 	msgClient   *core.Client
// 	tokenServer core.AccessTokenServer
// )

func init() {
	// tokenServer = core.NewDefaultAccessTokenServer(WxAppId, WxAppSecret, nil)
	// msgClient = core.NewClient(tokenServer, http.DefaultClient)
	// menuCreateHandler()

	TokenServer = core.NewDefaultAccessTokenServer(WxAppId, WxAppSecret, nil)
	Client = core.NewClient(TokenServer, http.DefaultClient)
	fmt.Println(WxAppSecret, WxAppId, TokenServer)
	// DeleteMenu()
	CreateMenu()
	// menuCreateHandler()
}

const (
	fileName = "menu.json"
)

type JSONConfig struct {
}

func (self *JSONConfig) Parse(fileName string) (*menu.Menu, error) {
	fmt.Println("fileName:", fileName)
	//绝对路径
	jsonpath, _ := filepath.Abs("conf/" + fileName)
	fmt.Println("jsonpath", jsonpath)
	file, err := os.Open(jsonpath)
	defer file.Close()
	if err != nil {
		errInfo := fmt.Sprintf("open %s occur error", fileName)
		panic(errInfo)
		return nil, err
	}
	datas, err := ioutil.ReadAll(file)
	if err != nil {
		errInfo := fmt.Sprintf("read %s occur error", fileName)
		panic(errInfo)
		return nil, err
	}
	menu := new(menu.Menu)
	err = json.Unmarshal(datas, menu)
	if err != nil {
		panic("json 反序列化失败!")
		return nil, err
	}
	return menu, nil
}

func ParseJson() {
	jsonConfig := &JSONConfig{}
	fileName := "menu.json"
	menu, _ := jsonConfig.Parse(fileName)
	for index, button := range menu.Buttons {
		fmt.Printf("button[%b] is %v \n", index, button)
	}
}

func CreateMenu() {
	jsonConfig := &JSONConfig{}
	menu1, err := jsonConfig.Parse(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(menu1)
	err = menu.Create(Client, menu1)

	if err == nil {
		fmt.Println("TestCreateMenu success")
	}

}
func DeleteMenu() {
	err := menu.Delete(Client)
	if err == nil {
		fmt.Println("TestDeleteMenu success")
	}
}

func MenuHandler() {
	CreateMenu()
}

//==============================================================
//$$$$$$$$$$$$$$$$$$$$$$$
//以下部分为程序生成菜单。
//$$$$$$$$$$$$$$$$$$$$$$$
//
// func CreateMenu() {
// 	var Menu_wx *menu.Menu
// 	Menu_wx = new(menu.Menu)
// 	var Btn_wx *menu.Button
// 	Btn_wx = new(menu.Button)

// 	Btn_wx.SetAsPicWeixinButton("test_Pic", "weixin")
// 	Menu_wx.Buttons = append(Menu_wx.Buttons, *Btn_wx)
// 	fmt.Println("Menu_wx:", Menu_wx)
// 	err := menu.Create(Client, Menu_wx)
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}

// }

// func DeleteMenu() {
// 	err := menu.Delete(Client)
// 	if err != nil {
// 		fmt.Println("DeleteMenu Error : ", err)
// 	}
// }

// func menuCreateHandler() {
// 	fmt.Println("********************************")
// 	// var Clt *core.Client
// 	var Btn1 *menu.Button
// 	var Btn2 *menu.Button
// 	var Btn3 *menu.Button

// 	var Menu_wx *menu.Menu

// 	Btn1 = new(menu.Button)
// 	Btn1.Name = "test1"

// 	// Btn1.SubButton = new(Btn1.SubButton)
// 	Btn1_1 := newButton()
// 	Btn1_1.SetAsLocationSelectButton("Location", "1_1")
// 	Btn1_2 := newButton()
// 	Btn1_2.SetAsPicPhotoOrAlbumButton("Pictice", "1_2")
// 	Btn1.SubButtons = append(Btn1.SubButtons, *Btn1_1)
// 	Btn1.SubButtons = append(Btn1.SubButtons, *Btn1_2)
// 	Btn2 = new(menu.Button)
// 	Btn2.Name = "test2"
// 	// Btn2.SubButton = new(Btn1.SubButton)
// 	Btn2_1 := newButton()
// 	Btn2_1.SetAsScanCodePushButton("ScanCode", "2_1")
// 	Btn2_2 := newButton()
// 	Btn2_2.SetAsScanCodeWaitMsgButton("CodeWait", "2_2")
// 	Btn2.SubButtons = append(Btn2.SubButtons, *Btn2_1)
// 	Btn2.SubButtons = append(Btn2.SubButtons, *Btn2_2)

// 	Btn3 = new(menu.Button)
// 	Btn3.Name = "test3"
// 	// Btn3.SubButton = new(Btn1.SubButton)
// 	Btn3_1 := newButton()
// 	Btn3_1.SetAsViewButton("view", "http://blog.jaxiu.cn")
// 	Btn3_2 := newButton()
// 	Btn3_2.SetAsPicWeixinButton("PicWeiXin", "3_2")
// 	Btn3.SubButtons = append(Btn3.SubButtons, *Btn3_1)
// 	Btn3.SubButtons = append(Btn3.SubButtons, *Btn3_2)

// 	Menu_wx = new(menu.Menu)
// 	Menu_wx.Buttons = make([]menu.Button, 0)
// 	Menu_wx.Buttons = append(Menu_wx.Buttons, *Btn1)
// 	Menu_wx.Buttons = append(Menu_wx.Buttons, *Btn2)
// 	Menu_wx.Buttons = append(Menu_wx.Buttons, *Btn3)
// 	fmt.Println(Menu_wx)

// 	var Btn *menu.Button
// 	Btn = new(menu.Button)
// 	Btn.SetAsPicSysPhotoButton("test1", "haha")
// 	var Menu_wxtest *menu.Menu
// 	Menu_wxtest = new(menu.Menu)
// 	Menu_wxtest.Buttons = append(Menu_wxtest.Buttons, *Btn)
// 	fmt.Println(Menu_wxtest)

// 	err = menu.Create(Client, Menu_wx)
// 	// err = menu.Create(msgClsient, Menu_wxtest)
// 	if err != nil {
// 		log.Println("Error : ", " Create Menu Error ! ", err)
// 	}
// 	// menu.Create(ctx., menu)
// }

// func newButton() (btn *menu.Button) {
// 	btn = new(menu.Button)
// 	return
// }
