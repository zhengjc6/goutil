package signalx

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SinganlxCancelFunc func()

func CatchSignal(f SinganlxCancelFunc) {
	//创建监听退出chan
	c := make(chan os.Signal, 1)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Printf("PID = %d,Program Exit...Catch Singal %v\n", os.Getpid(), s)
				fmt.Println("call CancelFunc")
				f()
				for i := 3; i >= 1; i-- {
					fmt.Printf("quit after %d sec\n", i)
					time.Sleep(time.Second)
				}
				os.Exit(0)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()
}
