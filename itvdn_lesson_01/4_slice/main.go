package main

import "fmt"

func main() {
	var sl1 []string
	fmt.Println("empty slice", sl1)

	sl1 = append(sl1, "Элемент1", "element2")
	fmt.Println(sl1, len(sl1))
}
