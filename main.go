package main

import (
	// "fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"log"
	"macaron/controller"
	// "macaron/models"
	"macaron/modules/initConf"
)

var conf *goconfig.ConfigFile

// init Database  &  init Config file
func init() {
	// models.RegisterDB()
	conf = initConf.InitConf()
	// fmt.Println(conf)
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
	m.Get("/file", FileHandler)
	m.Post("/fileup", FileUpdateHandler)

	m.Run(conf.MustInt("baseconfig", "port"))
}

func FileHandler(ctx *macaron.Context, log *log.Logger) {
	ctx.HTML(200, "fileup")
}

func FileUpdateHandler(ctx *macaron.Context, log *log.Logger) {
	ctx.Redirect("/", 301)
}
