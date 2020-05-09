package main

import "fmt"

func main() {
	fmt.Println("Show some text")
	//panic("Panic!!!")
	somePanic()

}

func somePanic() {
	fmt.Println("Show some text before")
	panic("Panic in function some Panic!!!")
	fmt.Println("Show some text after")
}
