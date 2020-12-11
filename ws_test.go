package wechat_bot_go

//var wsBot *websocketBot
//var httpBot *bot
//
//func TestMain(m *testing.M) {
//	url := "ws://127.0.0.1:5555"
//	wsBot = NewWebsocketBot(url)
//	wsBot.Connect()
//	prefix := "http://127.0.0.1:5555"
//	httpBot = NewBot(prefix)
//	defer wsBot.Close()
//	m.Run()
//}
//
//func TestWebsocketBot_SendTxtMsg(t *testing.T) {
//	wxid := "wxid_uspyqfp09fb621"
//	err := wsBot.SendTxtMsg("今天是个好日子", wxid)
//	t.Logf("err: %v\n", err)
//}
