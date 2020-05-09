package main

import (
	"fmt"
	"github.com/bhansconnect/factorial"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Factorial", i, "=", factorial.Factorial(int64(i)))
	}
}
