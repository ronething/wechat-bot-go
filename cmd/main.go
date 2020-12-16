package main

import (
	"flag"
	"fmt"
	wechat_bot_go "github.com/ronething/wechat-bot-go"
	"github.com/ronething/wechat-bot-go/config"
	"github.com/ronething/wechat-bot-go/scheduler"
	"github.com/ronething/wechat-bot-go/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	filePath string // 配置文件路径
	help     bool   // 帮助
)

func usage() {
	fmt.Fprintf(os.Stdout, `wechat-robot-demo - wechat push service demo
Usage: wechat-robot-demo [-h help] [-c ./config.yaml]
Options:
`)
	flag.PrintDefaults()
}

func main() {

	flag.StringVar(&filePath, "c", "./config.yaml", "配置文件所在")
	flag.BoolVar(&help, "h", false, "帮助")
	flag.Usage = usage
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}

	// 设置配置文件和静态变量
	config.SetConfig(filePath)

	// init wechat Router
	server.InitWechatHandlerRouter()
	// 初始化微信机器人
	//httpBot := wechat_bot_go.NewBot(config.Config.GetString("wechat.http"))
	//TODO: websocket 可作为后续被动回复使用
	wsBot := wechat_bot_go.NewWebsocketBot(config.Config.GetString("wechat.ws"))
	wxReply := server.NewWxReply(wsBot)
	wsBot.BindOnTextMessage(wxReply.BindFunc)
	wsBot.Connect() // 记得连接 不然会 nil error
	defer wsBot.Close()
	sched := scheduler.NewScheduler()
	//sched.InitJob(wsBot) TODO: 注册定时任务 后续可以作为 系统添加任务，我感觉脚本可以使用 python/shell 然后返回规定的状态码
	sched.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// DONE: 优雅关停
	for {
		s := <-c
		log.Printf("[main] 捕获信号 %s", s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			// 停止调度器 并等待正在 running 的任务执行结束 TODO: 有没有必要设置一个 timeout 假设一直不停止怎么办
			ctx := sched.Stop()
			timer := time.NewTimer(1 * time.Second)
			for {
				select {
				case s = <-c: // 再次接收到中断信号 则直接退出
					if s == syscall.SIGINT {
						log.Printf("[main] 再次接收到退出信号 %s", s.String())
						goto End
					}
				case <-ctx.Done():
					log.Printf("[main] 调度器所有任务执行完成")
					goto End
				case <-timer.C:
					log.Printf("[main] 调度器有任务正在执行中")
					timer.Reset(1 * time.Second)
				}
			}
		End:
			timer.Stop()
			return // 很重要 不然程序无法退出
		case syscall.SIGHUP:
			log.Printf("[main] 终端断开信号，忽略")
		default:
			log.Printf("[main] other signal")
		}
	}

}
