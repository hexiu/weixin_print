package NetSendPrintMsg

//服务器端
import (
	"fmt"
	"log"
	"net" //支持通讯的包
	"weixin_dayin/modules/initConf"
)

const (
	Server string = "wx.jaxiu.cn"
	Port   string = "7777"
)

var (
	server string = Server
	port   string = Port
)

type TxMsg struct {
	OpenId   string
	MsgType  string
	FileType string
	MsgInfo  string
	MsgURL   string
	Time     string
}

var TxChan = make(chan TxMsg, 5)

type ClientInfo map[string]string

var ClientList = make(map[string]string, 0)

func init() {
	conf, err := initConf.InitConf()
	if err != nil {
		log.Println("[Read Config File Error] : ", err)
	}
	if ok, err := conf.GetValue("Communication", "Server"); err == nil {
		server = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("Communication", "Port"); err == nil {
		port = ok
	} else {
		log.Println(err)
	}
}

func StartServer() {
	//连接主机、端口，采用ｔｃｐ方式通信，监听７７７７端口
	listener, err := net.Listen("tcp", server+":"+port)
	checkError(err)
	fmt.Println("建立成功!")
	for {
		//等待客户端接入
		conn, err := listener.Accept()
		checkError(err)
		//开一个goroutines处理客户端消息，这是golang的特色，实现并发就只go一下就好
		go doServerStuff(conn)
	}
}

//处理客户端消息
func doServerStuff(conn net.Conn) {
	nameInfo := make([]byte, 512) //生成一个缓存数组
	_, err := conn.Read(nameInfo)
	checkError(err)
	// fmt.Println(conn.RemoteAddr())
	ClientList[conn.RemoteAddr().String()] = "connect"
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf) //读取客户机发的消息
		flag := checkError(err)
		if flag == 0 {
			break
		}

		msginfo := fmt.Sprintln(string(buf))
		midinfo := msginfo[1:]
		midinfo = strings.Replace(midinfo, "}", " ", 1)
		info := strings.Fields(midinfo)
		msg := new(TxMsg)
		msg.OpenId = info[0]
		msg.MsgType = info[1]
		msg.FileType = info[2]
		msg.MsgInfo = info[3]
		msg.MsgURL = info[4]
		msg.PrintCode = info[5]
		msgtime := info[6]
		msg.Time, err = strconv.ParseInt(msgtime, 10, 64)
		if err != nil {
			log.Println(err)
		}
		if msg.MsgInfo == "quit" {
			return
		}

		fmt.Println(string(buf)) //打印出来
		MsgInfo := <-TxChan
		msgInfo := fmt.Sprintf("%v", MsgInfo)
		fmt.Println(msgInfo, ClientList)
		conn.Write([]byte(msgInfo))
	}
	return
}

//检查错误
func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			//fmt.Println("用户退出了")
			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}
