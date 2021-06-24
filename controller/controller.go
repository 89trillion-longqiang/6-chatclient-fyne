package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"chatClient/module"
	"chatClient/module/protobuf"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	HChan module.HandleChan
	Username string
	Server string
	c *websocket.Conn
)

func SetName(content string){
	Username = content
	return
}
func SetServer(content string){
	Server = content
	return
}
func HandleSetupHttp() string{
	if len(Server) == 0 || len(Username)== 0{
		log.Println("len(Server) || len(Username)  == 0")
		return "len(Server) || len(Username)  == 0"
	}
	c = InitRouter(Server, Username)
	if c == nil {
		log.Println("c *websocket.Conn is nil")
		return "c *websocket.Conn is nil"
	}
	HChan.UserChatChan = make(chan int)
	HChan.UserListChan = make(chan int)
	go HandReadMsg()
	return ""
}
func HandleDisCon() string {

	if c == nil {
		log.Println("c *websocket.Conn is nil")
		return  "c *websocket.Conn is nil"
	}

	err := c.Close()
	if err != nil {
		return "Close error"
	}
	return ""
}
func HandSendMsg(msg string) string{

	if c == nil {
		log.Println("c *websocket.Conn is nil")
		return "c *websocket.Conn is nil"
	}
	sendMsg ,_:= proto.Marshal(&protobuf.Communication{Class: "Talk",Msg: msg})
	err := c.WriteMessage(1,sendMsg)
	if err != nil{
		return err.Error()
	}
	return ""
}
func HandReadMsg(){
	if c == nil {
		log.Println("c *websocket.Conn is nil")
		return
	}

	c.SetReadLimit(module.MaxMessageSize)
	err := c.SetReadDeadline(time.Now().Add(module.PongWait))
	if err != nil {
		return
	}
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(module.PongWait)); return nil })

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var recCom protobuf.Communication
		errPro := proto.Unmarshal(message,&recCom)
		if errPro != nil {
			log.Printf("error: %v", err)
		}

		if recCom.Msg != ""{
			switch recCom.Class {
			case "Talk":
				HChan.UserChatMsg = recCom.Msg
				fmt.Println(recCom.Msg)
				HChan.UserChatChan <- 1
			case "userlist":
				HChan.UserListMsg = recCom.Msg
				fmt.Println(recCom.Msg)
				HChan.UserListChan <- 1
			}
		}
	}
}

func InitRouter(Server string,Username string)*websocket.Conn{///连接服务器

	if len(Server) == 0 && len(Username)== 0{
		return nil
	}
	log.Printf("connecting to %s", Server)
	head := http.Header{"username": {Username}}
	c, _, err := websocket.DefaultDialer.Dial(Server,head)
	if err != nil {

		log.Println("dial:", err)
		return nil
	}
	return c
}
