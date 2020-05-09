package main

import "fmt"

func main() {
	var someVar interface{} //пустой интерфейс, неопределённый тип

	someVar = 123              //с ним нельзя использовать преобразование типов,
	in, isInt := someVar.(int) // можно только использовать утверждения  типа(type assertion)
	if isInt {
		fmt.Println("integer!", in)
	}
	Print(someVar)

	someVar = "asdf"
	str, isString := someVar.(string) //утверждение что  someVar теперь string, можно использовать и свои типпы someVar.(User)
	if isString {
		fmt.Println("string!", str)
	}
	Print(someVar)

	switch v := someVar.(type) {
	case int:
		fmt.Println("integer!", v)
	case string:
		fmt.Println("string!", v)
	}

	switch someVar.(type) {
	case int:
		fmt.Println("integer!", someVar)
	case string:
		fmt.Println("string!", someVar)
	}

}

func Print(variable interface{}) { //корректность работы с пустым интерфейсом ложиться на разработчика!
	fmt.Println(variable)
}
