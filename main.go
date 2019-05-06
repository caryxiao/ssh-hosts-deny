package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("error: ", err)
	}

	defer watcher.Close()

	err = watcher.Add("/tmp/ghost-shadowsocks.log")

	if err != nil {
		fmt.Println("error: ", err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				if ev.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("写入文件")
					fmt.Printf("data: %v \n", ev.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error: ", err)
			}
		}
	}()

	waitSignal()
}

func waitSignal() {
	var signalChan = make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT)
	for sig := range signalChan {
		if sig == syscall.SIGHUP {
			fmt.Println("SIGHUP")
		} else if sig == syscall.SIGINT {
			fmt.Println(fmt.Sprintf("捕捉到信号signal: %v, ctrl+c", sig))
			os.Exit(0)
		} else {
			fmt.Println(fmt.Sprintf("捕捉到信号signal: %v", sig))
			os.Exit(0)
		}
	}
}
