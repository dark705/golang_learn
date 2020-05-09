package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Worker struct {
	name  string
	chDie chan bool
	*sync.WaitGroup
}

func (w *Worker) DoSomeWork() <-chan string {
	c := make(chan string)
	go func() {
		//defer w.Done()       //метод за счёт встраивания
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond): //с некоторой периодичностью делаем раобту
				c <- fmt.Sprintf("%s done some work# %d", w.name, rand.Intn(20)) //как только сделали пишем в канал.
			case <-w.chDie:
				fmt.Printf("%s, get die signal, must die\n", w.name)
				time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) //эмулирую некоторую задержку остановки
				w.Done()                                                      //метод за счёт встраивания, можно вынести в defer
				return
			}
		}
	}()

	return c
}

func (w *Worker) Die() {
	fmt.Printf("send to worker %s die signal\n", w.name)
	w.chDie <- true
	close(w.chDie) //закрываем канал
}

func NewWorker(name string, wg *sync.WaitGroup) *Worker {
	wg.Add(1)
	chDie := make(chan bool, 1)
	return &Worker{name, chDie, wg}
}

func main() {
	var wg sync.WaitGroup

	w1 := NewWorker("Vasya", &wg)
	w2 := NewWorker("Petya", &wg)

	cW1 := w1.DoSomeWork()
	cW2 := w2.DoSomeWork()

	for i := 0; i < 6; i++ {
		select {
		case r := <-cW1:
			fmt.Println(r)
		case r := <-cW2:
			fmt.Println(r)
		}
	}
	w1.Die()
	fmt.Println("send die 1", runtime.NumGoroutine())
	w2.Die()
	fmt.Println("send die 2", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("done", runtime.NumGoroutine())
}
