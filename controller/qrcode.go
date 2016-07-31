package controller

import (
	"fmt"
	"github.com/chanxuehong/wechat.v2/mp/qrcode"
)

func ProvideQrcode() {
	z, err := qrcode.CreateTempQrcode(Client, 553, 60)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(z)
	// z.Ticket = "haha"
	z.PermQrcode.URL = "http://blog.jaxiu.cn"
	fmt.Println(z)
	z.URL = "http://blog.jaxiu.cn"
	rel := qrcode.QrcodePicURL(z.Ticket)
	fmt.Println(rel)
}
