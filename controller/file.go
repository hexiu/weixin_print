package controller

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
)

func HaHandler(ctx *macaron.Context, log *log.Logger) {
	ctx.HTML(200, "fileup")
}
