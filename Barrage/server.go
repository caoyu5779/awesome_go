package main

import (
	"net/http"
	"github.com/gorilla/websocket"
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
		conn *websocket.Conn
		err error
		data []byte
	)
	//完成协议转换
	// Upgrade : webSocket
	if conn, err = upgrader.Upgrade(w, r, nil) ; err != nil {
		return
	}

	//webSocket.Conn
	for {
		// text | binary
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil{
			goto ERR
		}
	}

	ERR:
		conn.Close()

}

func main()  {
	//写一个http服务端
	// http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("0.0.0.0:7777", nil)
}

 