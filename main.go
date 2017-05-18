package main

import (
	"fmt"
	"log"
	"net/http"
	"websocker-hub/util"

	"golang.org/x/net/websocket"

	"encoding/json"
	"io/ioutil"
	"sync"
)

var (
	newsChan   = make(chan string)
	wsConnPool WsConnectionPool
)

// WsConnectionPool websocket连接池
type WsConnectionPool struct {
	websocketConn map[string]*websocket.Conn
	sync.RWMutex
}

type result struct {
	Module     string          `json:"module"`
	NoticeType string          `json:"type"`
	Content    json.RawMessage `json:"content"`
}

// 建立websocket的连接池并保存
func connection(ws *websocket.Conn) {
	var err error
	key := util.GetGUID()
	wsConnPool.websocketConn[key] = ws

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
		newsChan <- string(b)
	}
}

// 向建立的链接的websocker发送消息
func noticeMessagetoClient() {
	var err error
	for {
		x := <-newsChan
		for key, ws := range wsConnPool.websocketConn {
			if err = websocket.Message.Send(ws, x); err != nil {
				fmt.Println("Can't send ", err)
				ws.Close()
				wsConnPool.Lock()
				delete(wsConnPool.websocketConn, key)
				wsConnPool.Unlock()
			}
		}
	}
}

func main() {
	fmt.Println("websocket proxy start ... ")
	wsConnPool.websocketConn = make(map[string]*websocket.Conn)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/v1/socket/ws", websocket.Handler(connection))
	http.HandleFunc("/v1/socket/notice", notice)

	go noticeMessagetoClient()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
