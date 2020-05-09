package main

import "fmt"

type MyInt int16

func (MyInt) sayHello() {
	println("Hello")
}

func main() {
	var someVar MyInt

	someVar = 23
	someVar.sayHello()

	someVar2 := MyInt(2)
	fmt.Println(someVar2)

}
