package main

import "fmt"

func main() {
	var array1 [5]int
	array2 := [...]string{"some1", "some2"}
	array1[1] = 23
	fmt.Println(array1, array2)
}
