package models

import (
	// "fmt"
	"log"
	// "time"
)

type FileInfo struct {
	Id int64
	// Uid            int64
	Wid            string
	OpenId         string `xorm:"index"`
	FileWherePath  string //标识文件存储位置：在互联网还是微信端。
	FileName       string
	FileReName     string
	FileUploadDate string
	FilePrintTime  int64
	FilePayInfo    bool
	FileType       string
	MediaId        string
	MsgId          int64
	PrintNum       int
	Fee            int64
	FileUrl        string
	Flag           int
	OutTradeNo     string
	FileUploadTime int64 `xorm:"index"`
}

func AddFileInfo(fileinfo *FileInfo) (err error) {
	connectDB()
	_, err = engine.Insert(fileinfo)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func GetNotPrintFileInfo(openid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("open_id = ? and file_print_time = ? and flag = ? ", openid, 0, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil
}

func GetAllFileInfo(openid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("open_id = ? and flag = ? ", openid, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil

}

func GetPrintFileInfo(openid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("open_id = ? and file_print_time > ? and flag = ? ", openid, 1, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil
}

func GetPayNotPrintFileInfo(openid string, payinfo bool) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("open_id = ?  and file_print_time = ? and file_pay_info = ? and flag = ? ", openid, 0, payinfo, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil

}

func (f *FileInfo) SetFlag(flag int) {
	f.Flag = flag
	connectDB()
	_, err := engine.Update(f)
	if err != nil {
		log.Println("[Database Insert Error (FileInfo Updated): ]", err)
	}
	defer engine.Close()
}

func (f *FileInfo) SetFilePayInfo(payinfo bool) {
	f.FilePayInfo = payinfo
	connectDB()
	_, err := engine.Update(f)
	if err != nil {
		log.Println("[Database Insert Error (FileInfo Updated): ]", err)
	}
	defer engine.Close()
}

func GetFileInfo(id int64) (fileinfo *FileInfo, err error) {
	connectDB()
	defer engine.Close()
	fileinfo = &FileInfo{
		Id: id,
	}
	_, err = engine.Get(fileinfo)
	if err != nil {
		return nil, err
	}
	return fileinfo, nil
}

func UpdatePayFileInfo(oid string) (err error) {
	connectDB()
	defer engine.Close()
	sql := "update `file_info` set file_pay_info = ? where out_trade_no = ?"
	_, err = engine.Exec(sql, 1, oid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateFileInfo(fileinfo *FileInfo) (err error) {
	connectDB()
	defer engine.Close()
	log.Println("models Fileinfo:", fileinfo)
	_, err = engine.Id(fileinfo.Id).Update(fileinfo)
	if err != nil {
		return err
	}
	return nil
}
