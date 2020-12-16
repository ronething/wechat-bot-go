package server

import (
	"github.com/ronething/wechat-bot-go/trie"
	"strings"
	"time"
)

var wechatRouter *trie.Router

func InitWechatHandlerRouter() {
	wechatRouter = trie.NewRouter()
	wechatRouter.AddRoute("帮助说明", "/help", HelpUsage)
	wechatRouter.AddRoute("网易云音乐排行榜", "/music/top", MusicTop)
	wechatRouter.AddRoute("gocn 新闻", "/gocn", GocnNews)
}

func GocnNews(c *trie.Context) error {
	publishTime := time.Now()
	g := Gocn{}
	err, contents := g.GetNewsContent(publishTime)
	if err != nil {
		return err
	}
	content := strings.Join(contents, "")
	return c.Bot.SendTxtMsg(content, c.Param("wechat_wxid"))
}

func MusicTop(c *trie.Context) error {
	n := NetEaseRank{}
	s, err := n.GetTop10()
	if err != nil {
		return err
	}
	return c.Bot.SendTxtMsg(s, c.Param("wechat_wxid"))
}

func HelpUsage(c *trie.Context) error {
	text := wechatRouter.PrintRoutes()
	usage := "usage:\n" + text
	return c.Bot.SendTxtMsg(usage, c.Param("wechat_wxid"))
}
