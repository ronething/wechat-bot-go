package scheduler

import (
	"context"
	wechat_bot_go "github.com/ronething/wechat-bot-go"
	"github.com/ronething/wechat-bot-go/config"
	"github.com/ronething/wechat-bot-go/server"
	"log"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	C *cron.Cron
}

//NewScheduler 创建调度器
func NewScheduler() *Scheduler {
	optLogs := cron.WithLogger(
		cron.VerbosePrintfLogger(
			log.New(os.Stdout, "[Cron]: ", log.LstdFlags)))

	c := cron.New(optLogs)
	return &Scheduler{C: c}

}

func (s *Scheduler) Run() {
	s.C.Start()
}

//TODO: 重构 -> 可配置化，这样增加任务才不需要改代码
func (s *Scheduler) InitJob(bot wechat_bot_go.Bot) {
	var err error
	n := server.NetEaseRank{ // pre init
		Pre: config.Config.GetString("netease.pre"),
	}
	if _, err = s.C.AddFunc(config.Config.GetString("netease.spec"), func() {
		now := time.Now().Format("2006-01-02")
		if n.Pre == now {
			log.Printf("当天已经推送过了")
			return
		}
		res, err := n.GetTop10()
		if err != nil {
			log.Printf("获取排行榜失败, err: %v\n", err)
			return
		}
		// push wechat
		users := config.Config.GetStringSlice("netease.user")
		log.Printf("netease send user is %v\n", users)
		success := 0 // 计数
		failed := 0
		for i := 0; i < len(users); i++ {
			if err = bot.SendTxtMsg(res, users[i]); err != nil {
				log.Printf("推送发生 err: %v\n", err)
				failed += 1
			} else {
				//time.Sleep(1 * time.Second) // 休眠一秒才发
				success += 1
			}
		}
		log.Printf("推送成功: %v, 推送失败: %v\n", success, failed)
		n.Pre = now // 设置标志
	}); err != nil {
		log.Printf("添加任务失败, err: %v\n", err)
		return
	}

	g := server.Gocn{
		Pre: config.Config.GetString("gocn.pre"),
	}
	if _, err = s.C.AddFunc(config.Config.GetString("gocn.spec"), func() {
		log.Printf("执行任务将 gocn 新闻推送到微信")
		now := time.Now().Format("2006-01-02")
		log.Printf("dingtalk pre is %v, now is %v\n", g.Pre, now)
		if g.Pre != now { // 抓取
			err, contents := g.GetNewsContent(time.Now())
			if err != nil {
				log.Printf("获取新闻发生错误, err: %v\n", err)
				return
			}
			content := strings.Join(contents, "")
			// push wechat
			users := config.Config.GetStringSlice("gocn.user")
			log.Printf("netease send user is %v\n", users)
			success := 0 // 计数
			failed := 0
			for i := 0; i < len(users); i++ {
				if err = bot.SendTxtMsg(content, users[i]); err != nil {
					log.Printf("推送发生 err: %v\n", err)
					failed += 1
				} else {
					//time.Sleep(1 * time.Second) // 休眠一秒才发
					success += 1
				}
			}
			log.Printf("推送成功: %v, 推送失败: %v\n", success, failed)
			g.Pre = now
		}
		return
	}); err != nil {
		log.Printf("添加任务失败, err: %v\n", err)
		return
	}
}

func (s *Scheduler) Stop() context.Context {
	return s.C.Stop()
}
