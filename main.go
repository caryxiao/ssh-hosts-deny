package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type CmdConfig struct {
	File            string // secure file path
	SshLoginFailCnt int    // default ssh login failed count
}

func main() {

	var Config CmdConfig

	flag.StringVar(&Config.File, "f", "", "Please specify a file you need to monitor")
	flag.IntVar(&Config.SshLoginFailCnt, "cnt", 5, "ssh login failed count")
	flag.Parse()

	fmt.Println(Config)
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
