package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"weixin_dayin/modules/initConf"
)

const (
	DataType = "mysql"
	Database = "weixin_print"
	Username = "root"
	Password = "axiu"
	Host     = "127.0.0.1"
	Port     = "3306"
)

var (
	datatype string
	database string
	username string
	password string
	host     string
	port     string
)

var engine *xorm.Engine

// Data Struct
type UserInfo struct {
	Id       int64
	Name     string
	Password string
	WeiXinID string
	Created  time.Time `xorm:"index"`
	Flag     int64
}

type FileInfo struct {
	Id         int64
	UserId     int64
	WeiXinID   string
	FileName   string
	ReFileName string
	OtherIndo  string
	Flag       int64
	Created    time.Time `xorm:"index"`
}

func init() {
	err := initconf()
	if err != nil {
		fmt.Println(err)
	}

}

func initconf() (err error) {
	conf := initConf.InitConf()

	// fmt.Println(conf)

	if ok, err := conf.GetValue("DataControl", "DataType"); err == nil {
		datatype = ok
	} else {
		datatype = DataType
	}

	if ok, err := conf.GetValue("DataControl", "DataBase"); err == nil {
		database = ok
	} else {
		database = Database
	}
	if ok, err := conf.GetValue("DataControl", "Username"); err == nil {
		username = ok
	} else {
		username = Username
	}
	if ok, err := conf.GetValue("DataControl", "Password"); err == nil {
		password = ok
	} else {
		password = Password
	}
	if ok, err := conf.GetValue("DataControl", "Host"); err == nil {
		host = ok
	} else {
		host = Host
	}
	if ok, err := conf.GetValue("DataControl", "Port"); err == nil {
		port = ok
	} else {
		port = Port
	}

	return nil
}

func connectDB() (err error) {
	engine, err = xorm.NewEngine(datatype, username+":"+password+"@tcp("+host+":"+port+")"+"/"+database+"?charset=utf8")
	if err != nil {
		return err
	}
	return nil

}

func RegisterDB() (err error) {
	err = connectDB()
	if err != nil {
		return err
	}

	fmt.Println(engine.Ping())

	if ok, _ := engine.IsTableExist("UserInfo"); !ok {
		engine.CreateTables(new(UserInfo))
	}

	if ok, _ := engine.IsTableExist("FileInfo"); !ok {
		engine.CreateTables(new(FileInfo))
	}

	defer engine.Close()
	return nil
}
