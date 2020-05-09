package main

import "fmt"

func main() {
	Function1()
	fmt.Println(Function2(1, 2))
	fmt.Println(Function3(1, 2))
	fmt.Println(Function3(1, 2))
	fmt.Println(Function4(1, 2))
	fmt.Println(Function5([]int{1, 2, 3, 4}...))
	fmt.Println(Function6([]int{1, 2, 3, 4}))

	fmt.Println(Function7(2, 4))
	fmt.Println(Function8(2, 4, 5))
}

func Function1() {
	fmt.Println("Function 1")
}

func Function2(a int, b int) int {
	fmt.Println("Function 2")
	return a + b
}

func Function3(a, b int) int { //тип а одну строку
	fmt.Println("Function 3")
	return a + b
}

func Function4(a, b int) (res int) { //указываем какая переменная будет возвращена
	fmt.Println("Function 4")
	res = a + b
	return
}

func Function5(slice ...int) (sum int) {
	for _, value := range slice {
		sum += value
	}
	return
}

func Function6(slice []int) (sum int) {
	for index := range slice {
		sum += slice[index]
	}
	return
}

func Function7(a, b int) (int, int) {
	return a + b, a * b
}

func Function8(ints ...int) (sum int, mul int) {
	mul = 1
	for _, value := range ints {
		sum += value
		mul *= value
	}
	return
}
