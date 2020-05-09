package main

import (
	"fmt"
)

func main() {
	var someMap1 map[string]string                        //объявление
	someMap1 = map[string]string{}                        //инициализация
	var someMap11 map[string]string = map[string]string{} //в одну строку
	var someMap111 = map[string]string{}                  //можно и так

	someMap1["someKey"] = "SomeValue"
	someMap11["someKey"] = "SomeValue"
	someMap111["someKey"] = "SomeValue"
	fmt.Println(someMap1, someMap11, someMap111)

	someMap2 := map[string]string{} //полная инициализация
	someMap2["someOneKey"] = "someAnother value"
	fmt.Println(someMap2)

	var someMap3 = make(map[string]string) //короткая инициализация
	someMap3["some"] = "value"
	fmt.Println(someMap3)

	test, ok := someMap3["testkey"]
	fmt.Println(test, ok)

	someMap3["testAnotherKey"] = "some value"
	testKey(someMap3, "testAnotherKey")
	delete(someMap3, "testAnotherKey")
	testKey(someMap3, "testAnotherKey")

}

func testKey(someMap map[string]string, key string) {
	_, exist := someMap[key]

	if exist {
		println(key, "exist")
	} else {
		println(key, "not exist")
	}

}
