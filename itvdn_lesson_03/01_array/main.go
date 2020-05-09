package main

import "fmt"

func main() {
	ar := [...]string{"one", "two", "three", "four"}

	for i, value := range ar {
		fmt.Println(i, "=", value)
	}
}
