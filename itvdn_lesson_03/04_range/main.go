package main

import "fmt"

func main() {
	var array1 = [...]string{"one", "two", "three", "four", "five"}
	var slice = append(array1[:])
	iterator(slice)

	someMap := map[string]string{}
	someMap["a"] = "apple"
	someMap["b"] = "book"
	someMap["с"] = "cream"
	someMap["d"] = "dream"
	iteratorOfMap(someMap)

}

func iterator(iterable []string) {
	for index, value := range iterable {
		fmt.Println(index, value)
	}
}

func iteratorOfMap(iterable map[string]string) {
	for key, value := range iterable { //порядок обхода не гарантируется!!!
		fmt.Println(key, value)
	}
}
