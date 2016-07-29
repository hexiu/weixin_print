package models

import (
	// "fmt"
	"log"
	// "time"
)

type User struct {
	Uid                int64 `xorm:"index" autoincr`
	Wid                string
	OpenId             string `xorm:"index"`
	Username           string
	Password           string
	Email              string
	Telnumber          string
	TotalConsumption   float64
	FileSavePath       string
	UploadFileNum      int64
	PrintFileNum       int64
	CreateTime         int64 `xorm:"index"`
	UpdateTime         int64 `xorm:"index"`
	NearUpdateFileTime int64 `xorm:"index"`
	NotPrintFile       int
	Flag               int
	Nickname           string
	Sex                int
	Language           string
	City               string
	Province           string
	Country            string
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
