package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"io/ioutil"

	"code.google.com/p/go.net/websocket"
)

var (
	newsChan      = make(chan string)
	websocketConn []*websocket.Conn
)

// Echo echo
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

func send(res http.ResponseWriter, req *http.Request) {
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
			for index, ws := range websocketConn {
				if err = websocket.Message.Send(ws, x); err != nil {
					fmt.Println("Can't send ", err)
					websocketConn = append(websocketConn[:index], websocketConn[index+1:]...)
					ws.Close()
					break
				}
			}
		}
	}
}

func main() {
	fmt.Println("begin")

	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/socket", websocket.Handler(connection))
	http.HandleFunc("/v1/send", send)

	// go generate()
	go sendMsg()

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("end")
}

func generate() {
	for {
		time.Sleep(1 * time.Second)
		times := time.Now().Unix()
		newsChan <- fmt.Sprintf("%d", times)
	}
}
