package main

import "fmt"

func main() {
	someChannel := doubleChannelMultiplexer(someFunc("Ping"), someFunc("Pong"))

	for i := 0; i <= 9; i++ {
		fmt.Println(<-someChannel)
	}

}

//возвращает канал, в который объединяются 2 других канала
func doubleChannelMultiplexer(channel1, channel2 <-chan string) <-chan string {
	combinedChannel := make(chan string)
	go func() {
		for {
			combinedChannel <- <-channel1
		}
	}()
	go func() {
		for {
			combinedChannel <- <-channel2
		}
	}()

	return combinedChannel
}

//функция которая создаёт и возвращает канал, и пишет в него сообщения
func someFunc(message string) <-chan string {
	channel := make(chan string)
	count := 0
	go func() {
		for {
			count++
			channel <- fmt.Sprintf("%s %d", message, count)
		}
	}()
	return channel
}
