package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill, os.Interrupt)

	go func() {
		pid := os.Getpid()
		fmt.Printf("My pid is %d\n", pid)
		time.Sleep(time.Second * 5)
		process, _ := os.FindProcess(pid)
		_ = process.Kill()
	}()
	<-c
}
