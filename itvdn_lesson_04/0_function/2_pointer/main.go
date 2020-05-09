package main

import "fmt"

func main() {
	someInt := 1
	address := &someInt //указатель на ячейку памяти
	fmt.Println(someInt, &someInt, address, someInt)

	var someInt64 int8 = 1
	address64 := &someInt64
	fmt.Println(someInt64, &someInt64, address64, someInt64)

	someNum := 2
	setNum(someNum, 10)
	fmt.Println(someNum)

	//передаём ссылку на переменную
	setNumPtr(&someNum, 10)
	fmt.Println(someNum)

}

func setNum(num int, value int) {
	num = value
}

//первый аргумент указатель на тип int, второй аргумент просто значение int
func setNumPtr(num *int, value int) {
	*num = value
}
