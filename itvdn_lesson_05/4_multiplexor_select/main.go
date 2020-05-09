package main

import (
	"fmt"
	"time"
)

func main() {
	doubleChannel := doubleChannelMultiplexer(someSay("mes1"), someSay("mes2"))
	for {
		fmt.Println(<-doubleChannel)
	}
	fmt.Scanln()
}

func doubleChannelMultiplexer(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case c := <-ch1: //читаем из канала ch 1 в переменную c
				ch <- c //пишем в канал сh из переменной с
			case c := <-ch2:
				ch <- c
			}
		}
	}()
	return ch
}

func someSay(msg string) <-chan string {
	channel := make(chan string)
	i := 0
	go func() {
		for {
			i++
			channel <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * 300)
		}
	}()
	return channel
}
