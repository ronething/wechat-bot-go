package main

import (
	"encoding/json"
	"fmt"
	wechat_bot_go "github.com/ronething/wechat-bot-go"
	"github.com/sacOO7/gowebsocket"
	"log"
	"os"
	"os/signal"
	"time"
)

type SendMsg struct {
	Id string `json:"id"`
	Type int64 `json:"type"`
	Content string `json:"content"`
	WxId string `json:"wxid"`
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("ws://127.0.0.1:5555")

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Received connect error - ", err)
	}

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		//TODO: 这里做逻辑处理
		log.Println("Received message - " + message)
	}

	//socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
	//	log.Println("Received ping - " + data)
	//}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received pong - " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()
	defer socket.Close()

	s := SendMsg{
		Id:      time.Now().Format("20060102150405"),
		Type:    wechat_bot_go.UserList,
		Content: "user list",
		WxId:    "null",
	}
	b, err := json.Marshal(&s)
	if err !=nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//socket.SendText()
	socket.SendBinary(b)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt, now ready to exit..,")
			return
		}
	}
}
