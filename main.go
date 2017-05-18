package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"

	"io/ioutil"
)

var (
	newsChan      = make(chan string)
	websocketConn []*websocket.Conn
)

// 建立websocket的连接池并保存
func connection(ws *websocket.Conn) {
	var err error
	websocketConn = append(websocketConn, ws)
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive ")
			ws.Close()
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

// 通过api请求接受通知内容
func notice(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
		newsChan <- string(b)
	}
}

// 向建立的链接的websocker发送消息
func sendMsg() {
	var err error
	for {
		for x := range newsChan {
			fmt.Println("xx ", x)
			for index, ws := range websocketConn {
				if err = websocket.Message.Send(ws, x); err != nil {
					fmt.Println("Can't send ", err)
					websocketConn = append(websocketConn[:index], websocketConn[index+1:]...)
					ws.Close()
				}
			}
		}
	}
}

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/v1/socket/ws", websocket.Handler(connection))
	http.HandleFunc("/v1/socket/notice", notice)

	go sendMsg()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("end")
}
