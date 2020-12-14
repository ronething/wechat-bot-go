package server

import (
	"encoding/json"
	"fmt"
	wechat_bot_go "github.com/ronething/wechat-bot-go"
	"github.com/ronething/wechat-bot-go/config"
	"github.com/sacOO7/gowebsocket"
	"log"
	"strings"
)

type GeneralMessage struct {
	Id       string `json:"id"`
	Content  string `json:"content"` //3 pic msg 是一个字典
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Time     string `json:"time"`
	MsgType  int64  `json:"type"`
}

type WxReply struct {
	bot wechat_bot_go.Bot
}

func NewWxReply(bot wechat_bot_go.Bot) *WxReply {
	return &WxReply{bot: bot}
}

//业务逻辑
func (w *WxReply) BindFunc(message string, socket gowebsocket.Socket) {
	// TODO: 个人账号昵称 这里其实可能有点问题 如果@的昵称一样 则会有问题 不过我们还有群聊列表限制
	self := config.Config.GetString("self")
	log.Printf("bind received message: %s\n", message)
	var msg GeneralMessage
	var key string
	var element string
	var router string
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Printf("发生错误, err: %v\n", err)
		return
	}
	log.Printf("msg is %v\n", msg)
	msg.Content = strings.TrimSpace(msg.Content) //去除空格
	if msg.Receiver == "self" {
		log.Printf("个人接收\n")
		key = "admin"
		element = msg.Sender
		router = msg.Content
	} else if strings.HasSuffix(msg.Receiver, "@chatroom") &&
		strings.HasPrefix(msg.Content, self) { // TODO: 群聊 可能不严谨
		//群聊需要符合 @onething
		router = strings.Split(msg.Content, self)[1]
		log.Printf("router is %v\n", router)
		key = "group"
		element = msg.Receiver
	} else {
		log.Printf("不匹配,返回\n")
		return
	}

	if isInArray(key, element) { //判断是否在其中
		switch msg.MsgType {
		case wechat_bot_go.RecvTxtMsg:
			log.Printf("接收到文本消息\n")
			if strings.TrimSpace(router) != ""{
				err := w.bot.SendTxtMsg(router, element) // TODO: 路由树匹配 业务处理
				if err != nil {
					log.Printf("发送文本消息发生错误, err: %v\n", err)
					return
				}
			}else{
				err := w.bot.SendTxtMsg("未匹配到路由,请正确输入~", element)
				if err != nil {
					log.Printf("发送文本消息发生错误, err: %v\n", err)
					return
				}
			}
		case wechat_bot_go.RecvPicMsg:
			log.Printf("接收到图片消息\n")
		case wechat_bot_go.HeartBeat: // 这里其实走不到 因为上面有过滤条件
			log.Printf("接收到心跳消息\n")
		default: // 这里理论上也走不到
			log.Printf("其他消息\n")
		}
	} else {
		log.Printf("没有匹配到群聊/个人\n")
		return
	}

}

func isInArray(key, element string) bool {
	arr := config.Config.GetStringSlice(fmt.Sprintf("wx_reply.%s", key))
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			return true
		}
	}
	return false
}
