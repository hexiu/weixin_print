package models

import (
	// "fmt"
	"log"
	// "time"
)

type User struct {
	Uid                int64   `xorm:"index" autoincr` //用户唯一标识自增ID，方便管理
	Wid                string  //微信公众号标识ID
	OpenId             string  `xorm:"index"` //用户对于公众号的唯一标识ID
	Username           string  //用户名（说明：姓名）
	Password           string  //用户密码（留用）
	Email              string  //用户邮箱
	Telnumber          string  //用户手机
	TotalConsumption   float64 //用户消费的总额度
	FileSavePath       string  //用户文件存放地址
	UploadFileNum      int64   //用户上传文件总数
	PrintFileNum       int64   //用户已打印文件数
	CreateTime         int64   `xorm:"index"` //用户创建时间
	UpdateTime         int64   `xorm:"index"` //用户信息更新时间
	NearUpdateFileTime int64   `xorm:"index"` //用户最近一次更新上传文件的时间
	NotPrintFile       int     //没有打印的文件
	Flag               int     //一个标识符（备用）
	Nickname           string  //用户微信昵称
	Sex                int     //用户微信性别
	Language           string  //用户微信所用语言
	City               string  //用户所在城市
	Province           string  //用户省市
	Country            string  //用户所在国家
	IsSubscriber       int     //用户是否关注了微信号
}

func (user *User) GetCreateTime() int64 {
	return user.CreateTime
}

func (user *User) GetUpdateTime() int64 {
	return user.UpdateTime
}

func (user *User) GetNearUpdateFileTime() int64 {
	return user.NearUpdateFileTime
}

func (user *User) GetFlag() int {
	return user.Flag
}

func (user *User) GetUid() int64 {
	return user.Uid
}

func (user *User) GetWid() string {
	return user.Wid
}

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) GetTelNumber() string {
	return user.Telnumber
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetTotalConsumption() float64 {
	return user.TotalConsumption
}

func (user *User) GetFileSavePath() string {
	return user.FileSavePath
}

func (user *User) GetUploadFileNum() int64 {
	return user.UploadFileNum
}

func (user *User) GetPrintFileNum() int64 {
	return user.PrintFileNum
}

func (user *User) GetNoPrintFile() int {
	return user.NotPrintFile
}

func NewUser() *User {
	user := new(User)
	return user
}

func AddUser(user *User) (err error) {
	connectDB()
	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func CheckUser(userid int64) (has bool, err error) {
	connectDB()
	// user := &User{Uid: userid}
	// has, err = engine.Get(user)
	if err != nil {
		return false, err
	}
	defer engine.Close()

	return has, nil
}

func ModifyUser() {

}

func JudgeUser(user *User) (has bool, err error) {
	connectDB()
	has, err = engine.Get(user)
	// mt.Println(user, "judge user")
	if err != nil {
		return false, err
	}
	// fmt.Println(user, "judge user")
	defer engine.Close()
	return has, nil
}

func GetUser(openid string, wid string) (user *User, err error) {
	connectDB()
	user = &User{
		Wid:    wid,
		OpenId: openid,
	}

	has, err := engine.Get(user)
	if has != true || err != nil {
		log.Println("Models Modules UserHandler GetUser Error : ", err)
		return nil, err
	}
	defer engine.Close()
	return user, nil
}

func UpdateUserInfo(user *User) (err error) {
	connectDB()
	_, err = engine.Update(user)
	if err != nil {
		return err
	}
	defer engine.Close()
	return nil
}

func (u *User) SetFlag(flag int) {
	u.Flag = flag
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetNearUpdateFileTime(curtime int64) {
	u.NearUpdateFileTime = curtime
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetTotalConsumption(totalConsumption float64) {
	u.TotalConsumption = totalConsumption
}

func (u *User) SetPrintFileNum(printFileNum int64) {
	u.PrintFileNum = printFileNum
}

func (u *User) SetNoPrintFile(noPrintFile int) {
	u.NotPrintFile = noPrintFile
}
