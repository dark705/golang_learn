package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "13"
	i1, err := strconv.Atoi(str1) //strconv.ParseInt(str1, 10, 64)
	if err != nil {
		fmt.Println("error on parse")
	}
	fmt.Println(str1, i1+1)

	str2 := "4f"
	i2, err := strconv.ParseInt(str2, 16, 64)
	if err != nil {
		fmt.Println("error on parse")
	}
	fmt.Println(str2, i2)

	str3 := "1001111"
	i3, err := strconv.ParseInt(str3, 2, 64)
	if err != nil {
		fmt.Println("error on parse")
	}
	fmt.Println(str3, i3)

	str4 := "23.6"
	i4, err := strconv.ParseFloat(str4, 64)
	if err != nil {
		fmt.Println("error on parse")
	}
	fmt.Println(str4, i4+1.2)

	i5 := 34
	str5 := strconv.Itoa(i5)
	fmt.Println(str5 + "some")
}
