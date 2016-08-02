package models

import (
	// "fmt"
	"log"
	// "time"
)

type FileInfo struct {
	Id             int64
	Uid            int64
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
	Fee            float64
	FileUrl        string
	Flag           int
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

func GetNotPrintFileInfo(openid, wid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("wid = ? and open_id = ? and file_print_time = ? and flag = ? ", wid, openid, 0, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil
}

func GetAllFileInfo(openid, wid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("wid = ? and open_id = ? and flag = ? ", wid, openid, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil

}

func GetPrintFileInfo(openid, wid string) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("wid = ? and open_id = ? and file_print_time > ? and flag = ? ", wid, openid, 1, 0).Find(&fileinfolist)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	return fileinfolist, nil
}

func GetPayNotPrintFileInfo(openid, wid string, payinfo bool) (fileinfolist []*FileInfo, err error) {
	connectDB()
	fileinfolist = make([]*FileInfo, 0)
	err = engine.Where("wid = ? and open_id = ?  and file_print_time = ? and file_pay_info = ? and flag = ? ", wid, openid, 0, payinfo, 0).Find(&fileinfolist)
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

func UpdateFileInfo(fileinfo *FileInfo) (err error) {
	connectDB()
	defer engine.Close()
	_, err = engine.Update(fileinfo)
	if err != nil {
		return err
	}
	return nil
}
