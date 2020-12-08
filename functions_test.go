package wechat_bot_go

import (
	"testing"
)

var b *bot

func TestMain(m *testing.M) {
	prefix := "http://127.0.0.1:5555"
	b = NewBot(prefix)
	m.Run()
}

func TestBot_GetContactList(t *testing.T) {
	res, err := b.GetContactList()
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}
	for i := 0; i < len(res); i++ {
		t.Logf("name:%s, wxid: %s\n", res[i].Name, res[i].WxId)
	}
}

func TestBot_GetMemberId(t *testing.T) {
	res, err := b.GetMemberId()
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}
	t.Logf("res is %v\n", res)
}

func TestBot_GetMemberNick(t *testing.T) {
	res, err := b.GetMemberNick("7720909145@chatroom")
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}
	for i := 0; i < len(res); i++ {
		t.Logf("nick:%s, wxid:%s\n", res[i].NickName, res[i].WxId)
	}
}

func TestBot_SendTxtMsg(t *testing.T) {
	err := b.SendTxtMsg("/music/top", "gh_74f024ea6c9b")
	t.Logf("err: %v\n", err)
}

func TestBot_SendPic(t *testing.T) {
	err := b.SendPic("wxid_uspyqfp09fb621", "E:\\Documents\\wechat-bot-go\\pic.jpg")
	t.Logf("err: %v\n", err)
}

func TestBot_SendAttach(t *testing.T) {
	err := b.SendAttach("wxid_uspyqfp09fb621", "E:\\Documents\\wechat-bot-go\\msgtype.go")
	t.Logf("err: %v\n", err)
}

func TestBot_SendAtMsg(t *testing.T) {
	err := b.SendAtMsg(
		"生活总是这样",
		"7720909145@chatroom",
		"wxid_uspyqfp09fb621",
		"xxx", // @之后的名字
	)
	t.Logf("err: %v\n", err)
}

func TestBot_RefreshMemberList(t *testing.T) {
	err := b.RefreshMemberList()
	t.Errorf("err: %v\n", err)
}

func TestBot_SendDestroy(t *testing.T) {
	err := b.SendDestroy()
	t.Logf("err: %v\n", err)
}
