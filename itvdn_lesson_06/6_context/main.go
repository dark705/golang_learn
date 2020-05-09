package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, chancel := context.WithTimeout(context.Background(), time.Second*5)
	c := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go writer(ctx, c, wg)
	go reader(ctx, c, wg)

	//chancel by os signal
	go func(chancel func()) {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-osSignals
		chancel()
	}(chancel)

	wg.Wait()
	close(c)
}

func writer(ctx context.Context, c chan string, wg *sync.WaitGroup) {
	ticker := time.NewTicker(time.Millisecond * 600)
	for {
		select {
		case <-ticker.C:
			c <- "some text"
		case <-ctx.Done():
			ticker.Stop()
			wg.Done()
			return
		}
	}
}

func reader(ctx context.Context, c chan string, wg *sync.WaitGroup) {
	for {
		select {
		case text := <-c:
			fmt.Printf("%s income: %s\n", time.Now().String(), text)
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
