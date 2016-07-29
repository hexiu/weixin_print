package initConf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func InitConf() (conf *goconfig.ConfigFile) {
	conf, err := goconfig.LoadConfigFile("conf/app.conf")
	if err != nil {
		fmt.Println(err)
	}
	return
}
