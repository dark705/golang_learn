package main

import (
	"fmt"
	"time"
)

func main() {
	someChannel1 := make(chan string)
	someChannel2 := make(chan string)

	go func() {
		for {
			someChannel1 <- "Ping to channel 1"
			time.Sleep(time.Millisecond * 700)
		}
	}()

	go func() {
		for {
			someChannel2 <- "Ping to channel 2"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	//в горутине читаем чере оператор select
	go func() {
		for {
			select {
			case msg1 := <-someChannel1:
				fmt.Println("Case msg1", msg1)
			case msg2 := <-someChannel2:
				fmt.Println("Case msg2", msg2)
			case <-time.After(time.Millisecond * 200): //каждые 200 милисекунд секунду true, это канал
				fmt.Println("TimeOut")
			}
		}
	}()

	fmt.Scanln()
}
