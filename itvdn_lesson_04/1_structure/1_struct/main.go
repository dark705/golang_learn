package main

import "fmt"

type User struct {
	Id       int64
	Name     string
	LastName string // this only works with strings
	Age      int16
}

func main() {
	var user1 User
	user1.Id = 1
	user1.Name = "Ivan"
	user1.LastName = "Ivanov"
	user1.Age = 23

	user2 := User{Id: 2, Name: "Vasya", LastName: "Pupkin", Age: 45} //декларация литералов
	user3 := User{3, "Peater", "J", 33}                              //без дакларации, последовательность определенна структурой
	user4 := User{}                                                  //получим значения по умолчанию

	//var user5 *User
	var user5 = &User{
		Id:       3,
		Name:     "as",
		LastName: "fd",
		Age:      5,
	}

	ShowUserInfoWithFields(user1)
	ShowUserInfoWithFields(user2)
	ShowUserInfoWithFields(user3)
	ShowUserInfoWithFields(user4)
	fmt.Println(user3.Name)
	fmt.Println(*user5)

}

func ShowUserInfoWithFields(u User) {
	fmt.Printf("%+v\n", u)
}
