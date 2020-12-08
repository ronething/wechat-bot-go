package wechat_bot_go

import (
	"fmt"
	"github.com/imroc/req"
	"log"
	"time"
)

func getId() string {
	return time.Now().Format("20060102150405")
}

type GetContactList struct {
	Content []*Friend `json:"content"`
	Type    int64     `json:"type"`
}

type GetMemberList struct {
	Content []*Member `json:"content"`
	Type    int64     `json:"type"`
}

type Friend struct {
	Name string `json:"name"`
	WxId string `json:"wxid"`
}

type Member struct {
	NickName string `json:"nickname"`
	RoomId   string `json:"roomid"`
	WxId     string `json:"wxid"`
}

type Bot interface {
	//群聊发送 @member 消息
	SendAtMsg(content, roomId, atWxId, nick string) error
	//发送文本消息
	SendTxtMsg(content, wxId string) error
	//发送图片
	SendPic(wxId, filePath string) error
	//发送附件
	SendAttach(wxId, filePath string) error
	//获取群成员昵称
	GetMemberNick(roomId string) (res []*Member, err error)
	//获取群成员 id
	GetMemberId() (resp interface{}, err error)
	//获取微信通讯录好友和群列表
	GetContactList() (res []*Friend, err error)
	//刷新好友列表
	RefreshMemberList() error
	//动态卸载 ddl
	SendDestroy() error
}

type bot struct {
	client *req.Req
	// TODO: config
	prefix string
}

func NewBot(prefix string) *bot {
	return &bot{client: req.New(), prefix: prefix}
}

func (b *bot) SendAtMsg(content, roomId, atWxId, nick string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendatmsg", req.QueryParam{
		"id":      getId(),
		"type":    AtMsg,
		"roomid":  roomId,
		"wxid":    atWxId,
		"content": content,
		"nick":    nick,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *bot) SendTxtMsg(content, wxId string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendtxtmsg", req.QueryParam{
		"id":      getId(),
		"type":    TxtMsg,
		"wxid":    wxId,
		"content": content,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v, status code is %d\n", resp.String(), resp.Response().StatusCode)
	return nil
}

func (b *bot) SendPic(wxId, filePath string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendpic", req.QueryParam{
		"id":   getId(),
		"type": PicMsg,
		"wxid": wxId,
		"path": filePath,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *bot) SendAttach(wxId, filePath string) error {
	resp, err := b.client.Get(b.prefix+"/api/sendattatch", req.QueryParam{
		"id":       getId(),
		"wxid":     wxId,
		"filepath": filePath,
	})
	if err != nil {
		return err
	}
	log.Printf("resp is %v\n", resp.String())
	return nil
}

func (b *bot) GetMemberNick(roomId string) (res []*Member, err error) {
	resp, err := b.client.Get(b.prefix+"/api/getmembernick", req.QueryParam{
		"id":     getId(),
		"type":   ChatroomMemberNick,
		"roomid": roomId,
	})
	if err != nil {
		return nil, err
	}
	var getMemberList GetMemberList
	if err = resp.ToJSON(&getMemberList); err != nil {
		return nil, err
	}
	//fmt.Printf("resp is %v\n", resp)
	return getMemberList.Content, nil
}

func (b *bot) GetMemberId() (res interface{}, err error) {
	resp, err := b.client.Get(b.prefix + "/api/getmemberid")
	if err != nil {
		return nil, err
	}
	fmt.Printf("resp is %v\n", resp)
	return resp, nil
}

func (b *bot) GetContactList() (res []*Friend, err error) {
	resp, err := b.client.Get(b.prefix + "/api/getcontactlist")
	if err != nil {
		return nil, err
	}
	//fmt.Printf("resp is %v\n", resp.String())
	var getContactList GetContactList
	if err = resp.ToJSON(&getContactList); err != nil {
		return nil, err
	}
	return getContactList.Content, nil
}

func (b *bot) RefreshMemberList() (err error) {
	resp, err := b.client.Get(b.prefix+"/api/refresh_chatroom", req.QueryParam{
		"id": getId(),
	})
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp)
	return nil
}

func (b *bot) SendDestroy() error {
	resp, err := b.client.Get(b.prefix+"/api/destroy", req.QueryParam{
		"id": getId(),
	})
	if err != nil {
		return err
	}
	fmt.Printf("resp is %v\n", resp)
	return nil
}
