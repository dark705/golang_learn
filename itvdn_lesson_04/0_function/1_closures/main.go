package main

import (
	"fmt"
	"time"
)

func main() {
	showExecuteTime := showExecuteTime() //можно было defer showExecuteTime()()
	defer showExecuteTime()

	// простое замыкание
	var inc = inc()
	inc()
	inc()
	inc()
	inc()
	//если бы 	inc()(), то всегда возвращалось бы 1

	//анонимная функция
	anonFunc := func() {
		fmt.Println("Execute anon func")
	}
	anonFunc()

	//самовызывная анонимная функция
	func() {
		fmt.Println("Execute right now")
	}()

	time.Sleep(time.Millisecond)
}

func showExecuteTime() func() {
	startTime := time.Now()
	return func() {
		fmt.Println(time.Since(startTime))
	}
}

func inc() func() {
	var count int64
	return func() {
		count++
		fmt.Println(count)
	}
}
