package main

import (
	"fmt"
	"time"
)

func main() {
	someSimpleChannel := make(chan string) //создали обычный каннал канал типпа string, это синхронный канал
	go someFunc(someSimpleChannel)

	for i := 0; i <= 2; i++ {
		fmt.Println(<-someSimpleChannel, ":", <-someSimpleChannel) //читакм из канала <-someSimpleChannel, если бы писали то было бы: someSimpleChannel<-
	}

	//асинхронно
	someBuffChannel := make(chan string, 10) //создали буферезированный канал типпа string, длиной 10, это АСИНХРОННЫЙ канал
	for i := 0; i <= 5; i++ {
		someBuffChannel <- "someString" // пишем в канал 5 раз , если бы записали больше чем 10, то заблокировали бы канал (dead lock)
	}
	close(someBuffChannel)                     //закрытик канала, больше в него писать нельзя
	fmt.Println(myConcatFunc(someBuffChannel)) //выполняем функцию

	someGeneratedChannel := channelGenerator(false)
	for i := 0; i <= 4; i++ {
		fmt.Println(<-someGeneratedChannel)
	}
}

func someFunc(c chan<- string) { //<- не озяательно, бубем только писать, передавать в канал
	for { //бесконечный цикл
		c <- "Ping" //передаём, пишем в канал
		c <- "Pong"
	}
}

func myConcatFunc(c chan string) string {
	r := ""
	for part := range c {
		r += part + "_"
	}
	return r
}

//функция генератор возвращает числа, либо чётные либо не чётные
func channelGenerator(even bool) <-chan int {
	channel := make(chan int)
	go func() {
		i := 0
		for {
			if i%2 == 0 {
				if even {
					channel <- i
				}
			} else {
				if !even {
					channel <- i
				}
			}
			time.Sleep(time.Millisecond * 100)
			i++
		}
	}()

	return channel
}
