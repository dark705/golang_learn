package main

import "fmt"

func main() {
	var someSlice = []int{1, 2, 3, 5, 3, 1, 6, 7, 2, 1, 1}

	fmt.Println(someSlice)
	statistic(someSlice)

}

func statistic(slice []int) {
	statistic := map[int]int{} //карта а не слайс, иначе нам надо будет делать слай длиной в максимальное значение слайса
	for _, value := range slice {
		statistic[value]++
	}
	fmt.Println(statistic)
}
