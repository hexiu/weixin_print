package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	// "log"
	"weixin_dayin/controller"
	"weixin_dayin/models"
	// "github.com/go-macaron/binding"
	// "mime/multipart"
	// "path"
	"weixin_dayin/modules/initConf"
)

const (
	Port int = 8080
)

var (
	port int = Port
)

var conf *goconfig.ConfigFile

// init Database  &  init Config file
func init() {
	err := models.RegisterDB()
	if err != nil {
		fmt.Println("Error : ", err)
	}
	conf = initConf.InitConf()
	if err != nil {
		fmt.Println("Load Config File Error! \t", err)
	}

	if ok := conf.MustInt("Server", "ListenPort"); ok != 0 {
		port = ok
	}

}

func main() {
	m := macaron.Classic()
	//Register middle key
	m.Use(macaron.Renderer())
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(session.Sessioner())
	// Router info
	m.Get("/", controller.HomeHandler)
	m.Get("/file", controller.FileHandler)
	m.Post("/fileup", controller.UploadHandler)

	m.Run(port)
}
