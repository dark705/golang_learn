//если все горутины окажутся в состоянии ожидания
// runtime выкенет выкидывает панику с сообщением deadlock
package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits uint64
}

func main() {
	playingField := make(chan *Ball)
	go player("Vasya", playingField)
	go player("Petya", playingField)

	playingField <- new(Ball) //если не впустим мятч то deadlock

	fmt.Scanln()   //конец игры
	<-playingField //забираем мятч
}

func player(name string, playingField chan *Ball) {
	for {
		ball := <-playingField //получаем мятч
		ball.hits++
		fmt.Println("Player:", name, "get ball on", ball.hits, "ball hits")
		time.Sleep(time.Millisecond * 200) //удерживаем мятч 200 милисекунд
		playingField <- ball               //возвращаем мятч
	}
}
