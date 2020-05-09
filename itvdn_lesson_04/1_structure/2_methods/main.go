package main

import "fmt"

type User struct {
	name string
	age  int16
}

//Обычная функция
func ShowInfo(u User) {
	fmt.Printf("%+v\n", u)
}

//Метод типпа User
func (u User) MethodShowInfo() {
	fmt.Printf("%+v\n", u)
}

//Ещё один метод
func (u User) GetAge() int16 {
	return u.age
}

func main() {
	user1 := User{"SomeName", 23}
	ShowInfo(user1)
	user1.MethodShowInfo()
	fmt.Println("Age", user1.GetAge())
}
