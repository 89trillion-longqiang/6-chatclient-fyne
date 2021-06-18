package router

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

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