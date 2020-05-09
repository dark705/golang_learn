package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type CustomError1 struct {
	msg string
}

func (e *CustomError1) Error() string {
	return e.msg
}

type CustomError2 struct {
	msg string
}

func (e *CustomError2) Error() string {
	return fmt.Sprintf("Some additional text %s", e.msg)
}

func main() {
	c := make(chan error)

	//error generator
	go func(c chan error) {
		rand.Seed(time.Now().UnixNano())
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			if rand.Intn(10) < 5 {
				c <- &CustomError1{msg: "Custom error1"}
			} else {
				c <- &CustomError2{msg: "Custom error2"}
			}
		}
	}(c)

	go func(c chan error) {
		for {
			err := <-c
			switch err.(type) {
			case *CustomError1:
				fmt.Printf("Custom error1: %s\n", err)
			case *CustomError2:
				fmt.Printf("Custom error2: %s\n", err)
			default:
				fmt.Printf("Default or unknown error2: %s", err)
			}
		}

	}(c)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Printf("Got OS signal: %v. Exit...\n", <-osSignals)
}
