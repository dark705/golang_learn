package main

import (
	levvv "01/8_package/src/some/level"
	"01/8_package/src/user"
	"fmt"
)

/*
GOPATH путь к проекту
GOROOT путь к месту где установлен go

Структура
	src - код и другие пакеты
	bin - бинарники
	pkg - объектные файлы создающиеся при компиляции
*/

func main() {
	fmt.Println(user.SomePublic)
	user.SomePublicFunc()
	fmt.Println(user.SomePublicConst)
	println(levvv.SomeLevelVar)
}
