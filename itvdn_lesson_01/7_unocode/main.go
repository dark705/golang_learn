package main

import "fmt"

func main() {

	var str1 string = "Zz"
	var str2 string = "Яя"
	var symbol rune = 'Я' // всегда одиночные кавычки, и всегда один символ
	fmt.Println(len(str1), str1, len(str2), str2, symbol)

	byt := []byte(str2)
	fmt.Println(byt)

	for index, value := range str1 {
		println(index, string(value))
	}

	for index, value := range str2 {
		println(index, string(value))
	}

}
