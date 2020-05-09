package main

import (
	"fmt"
	"os"
)

func main() {
	someSlice := []string{"one", "two"}

	ShowSliceInfo(someSlice)
	someSlice = append(someSlice, "three") // было 2, не влезло, увеличиваем в 2, стало 4
	ShowSliceInfo(someSlice)
	someSlice = append(someSlice, "four") //было 4, влезло, не увеличиваем
	ShowSliceInfo(someSlice)
	someSlice = append(someSlice, "five") //было 4, не вдезло, увеличиваем в 2, стало 8
	ShowSliceInfo(someSlice)

	someAnotherSlice := []string{"on", "the"}
	someSlice = append(someSlice, someAnotherSlice...) //синтаксический сахар ...
	ShowSliceInfo(someSlice)

	slMore := make([]string, 6) // длина и капасити 6
	ShowSliceInfo(slMore)

	slMoreCap := make([]string, 6, 7) // длина 6 и капасити 7
	ShowSliceInfo(slMoreCap)

	someSlice2 := someSlice //передаётся по ссылке
	someSlice2[0] = "new"
	ShowSliceInfo(someSlice)
	ShowSliceInfo(someSlice2)

	someSliceCopy := make([]string, len(someSlice), cap(someSlice)) //создаём аналогичный слайс
	copy(someSliceCopy, someSlice)                                  //копируем
	someSliceCopy[0] = "NEW"
	ShowSliceInfo(someSlice)
	ShowSliceInfo(someSliceCopy)

	//часть слайся
	fmt.Println(someSliceCopy[2:4])
	fmt.Println(someSliceCopy[:1])
	fmt.Println(someSliceCopy[5:])

	var someArray = [2]string{"one_array", "two_array"}
	fmt.Println(someArray)
	sliceFromArray := someArray[:]
	fmt.Println(sliceFromArray)

	os.Exit(0)

}

func ShowSliceInfo(slice []string) {
	fmt.Println(slice, "len:", len(slice), "cap:", cap(slice))
}
