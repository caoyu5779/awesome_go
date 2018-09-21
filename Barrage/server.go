package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"selfLearning/Barrage/impl"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//握手过程中 涉及到跨域问题，允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		data []byte
		conn *impl.Connection
	)
	//完成协议转换
	// Upgrade : webSocket
	if wsConn, err = upgrader.Upgrade(w, r, nil) ; err != nil {
		return
	}

	if conn,err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heart beat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}

	}()

	for {
		 if data, err = conn.ReadMessage(); err != nil {
		 	goto ERR
		 }

		 if err = conn.WriteMessage(data); err != nil {
		 	goto ERR
		 }
	}

	ERR:
		//TODO: 关闭链接的操作
		conn.Close()

}

func main()  {
	//写一个http服务端
	// http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("0.0.0.0:7777", nil)
}

 