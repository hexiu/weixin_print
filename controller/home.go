package controller

import (
	"gopkg.in/macaron.v1"
	"log"
)

func HomeHandler(ctx *macaron.Context, logger *log.Logger) {
	ctx.HTML(200, "index")
}

func FileHandler(ctx *macaron.Context, log *log.Logger) {
	ctx.HTML(200, "fileup")
}

func FileUpdateHandler(ctx *macaron.Context, log *log.Logger) {
	ctx.Redirect("/", 301)
}
