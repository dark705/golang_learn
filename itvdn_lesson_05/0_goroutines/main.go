package main

import (
	"fmt"
)

type User struct {
	id int
}

func (u User) SomeMethod() {
	fmt.Println("Method of user", u.id)
}

func main() { //основоной поток
	fmt.Println("Start")
	// последовательность не гарантиркуется
	for i := 0; i <= 4; i++ {
		user := User{i}
		go someGoRoutineFunct(i) // это вызов горутины, легковестного потока
		go user.SomeMethod()
	}

	go func() {
		fmt.Println("Anon function, go routine")
	}()

	fmt.Println("End")

	fmt.Scanln() //должны тем или иным способом дождаться выполнения
	//time.Sleep(time.Second)
}

func someGoRoutineFunct(variable int) {
	fmt.Println(variable)
}
