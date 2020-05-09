package main

import "fmt"

var someGlobalMap = map[string]string{}

func init() {
	someGlobalMap["someGlobalKey"] = "someGlobalValue"
}

func main() {
	//объявление
	var map1 map[string]string
	//инициализация
	map1 = map[string]string{}
	//присвоение значений
	map1["someKey"] = "someValue"
	fmt.Println(map1)

	//полная инициализация
	map2 := map[string]string{}
	map2["key1"] = "Value1"
	map2["key2"] = "Value2"
	fmt.Println(map2)

	//короткая инициализация
	var map3 = make(map[string]string)
	map3["some"] = "value"
	fmt.Println(map3)

	//ещё один вариант
	var map4 map[string]string = map[string]string{}
	var map5 = map[string]string{}
	fmt.Println(map4, map5)

	//т.е. init выполняется до main
	fmt.Println(someGlobalMap)
}
