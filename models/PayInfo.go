package models

import (
	"fmt"
	"time"
)

type PayInfo struct {
	Zid           int64 `xorm:"index autoincr"`
	Wid           string
	OpenId        string `xorm:"index"`
	PrintFile     string
	PrintFiletype string
	PrintFileUrl  string
	PrintFiletime string
	PayMoney      float64
	PayTime       int64 `xorm:"index"`
	CreateTime    int64 `xorm:"index"`
	PrintOk       bool
	PayOk         bool
}

func test1() {
	fmt.Println("...", time.Now())
}
