package routers

import (
	"container/list"
	"log"

	"../base"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketRouter struct {
	base.Router
}

func (this *WebSocketRouter) WebSocket() {
	this.Data["url"] = "ws://" + this.Ctx.Request.Host + "/wsecho"
	this.Show("websocket.tpl")
}

var conns = list.New()

func (this *WebSocketRouter) WSEcho() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("connected.")
	conns.PushBack(ws)
	//defer c.Close()
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		for e := conns.Front(); e != nil; e = e.Next() {
			err = e.Value.(*websocket.Conn).WriteMessage(messageType, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func init() {
	beego.Router("/ws", &WebSocketRouter{}, "*:WebSocket")
	beego.Router("/wsecho", &WebSocketRouter{}, "*:WSEcho")
}
