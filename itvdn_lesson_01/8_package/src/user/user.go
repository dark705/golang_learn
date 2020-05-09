package user

import "fmt"

var SomePublic = "Some public prop"
var somePrivate = "Some private prop"

const SomePublicConst = "const public"
const somePrivateConst = "const private"

func SomePublicFunc() {
	fmt.Println("Some public method")
	somePrivateFunc()
	fmt.Println(somePrivateConst)
}

func somePrivateFunc() {
	fmt.Println("Some private method")
}
