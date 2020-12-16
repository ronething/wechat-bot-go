package trie

import wechat_bot_go "github.com/ronething/wechat-bot-go"

type HandlerFunc func(*Context) error

type Context struct {
	Path   string
	Bot    wechat_bot_go.Bot
	Params map[string]string
}

func NewContext(path string, bot wechat_bot_go.Bot) *Context {
	return &Context{
		Path: path,
		Bot:  bot,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
