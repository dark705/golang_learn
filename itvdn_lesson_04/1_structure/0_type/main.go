package main

import "fmt"

type Distance int64
type Library []string

func main() {
	var toMoon Distance = 2
	ShowDistance(toMoon)

	//так нельзя
	//var toSun int64 = 23
	//ShowDistance(toSun)

	var lib1 Library = Library{"book1", "book2", "book3"}
	var lib2 = Library{"book11", "book12", "book13"}
	ShowBooksInLibrary(lib1)
	ShowBooksInLibrary(lib2)
	ShowBooksInLibrary(Library{"windows", "linux", "someAnotherBook"})
}

func ShowDistance(d Distance) {
	fmt.Println("Distance to object is", d)
}

func ShowBooksInLibrary(l Library) {
	for index, book := range l {
		fmt.Println(index, book)
	}
}
