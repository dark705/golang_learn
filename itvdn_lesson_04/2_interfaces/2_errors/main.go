package main

import (
	"errors"
	"fmt"
)

type UserError struct {
	Message string
	Code    int
}

func (e UserError) Error() string { // IDE подсвечивает, что мы фактически реализовали интерфейс Error
	return fmt.Sprintf("%s %d", e.Message, e.Code)
}

func GeneratorError() error { // фактически мы возвращаем интерфейс типпа error
	return UserError{"GeneratedErrorMessage", 555}
}

func main() {
	someError1 := GeneratorError()
	someError2 := errors.New("Some Error message")
	ShowErrorType(someError1)
	ShowErrorType(someError2)
}

func ShowErrorType(e error) {
	switch e.(type) {
	case nil:
		fmt.Println("No Error")
	case UserError:
		fmt.Println("UserError")
		fmt.Println(e)
	case error:
		fmt.Println("PackError")
		fmt.Println(e)
	}
}
