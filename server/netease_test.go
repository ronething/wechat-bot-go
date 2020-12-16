package server

import (
	"testing"

	"github.com/ronething/wechat-bot-go/config"
)

func TestNetEaseRank_GetTop10(t *testing.T) {
	config.SetConfig("E:\\Documents\\wechat-bot-go\\config\\config.yaml")
	n := NetEaseRank{}
	s, err := n.GetTop10()
	if err != nil {
		t.Errorf("err: %v\n", err)
		return
	}
	t.Logf("res is %v\n", s)
}
