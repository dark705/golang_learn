package main

import (
	"encoding/json"
	"fmt"
)

type Encode struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

type NotEncode struct {
	Id       int    `json:"-"`
	PassHash string `json:"-"`
}

type User struct {
	Encode
	NotEncode
}

func main() {
	u1 := &User{
		Encode:    Encode{Name: "Vasya", LastName: "Pupkin"},
		NotEncode: NotEncode{1, "asdf"},
	}

	users := []User{*u1, *u1}
	j, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error on encode")
	}
	fmt.Println(users)
	fmt.Println(string(j))

	//decode
	jString := `[{"name":"Vasya","lastName":"Pupkin"},{"name":"Vasya","lastName":"Pupkin"}]`
	j = []byte(jString)
	var usersRestored []User
	err = json.Unmarshal(j, &usersRestored)
	if err != nil {
		fmt.Println("error on decode")
	}
	fmt.Println(usersRestored)

	var usersRestored2 []map[string]string
	err = json.Unmarshal(j, &usersRestored2)
	if err != nil {
		fmt.Println("error on decode")
	}
	fmt.Println(usersRestored2)

	//additional
	j = []byte(`[{"nameAdditional":"Petya", "name":"Vasya","lastName":"Petrov"}]`)
	var usersRestoredAdditional []struct {
		Name           string
		LastName       string
		NameAdditional string `json:"nameAdditional"`
	}
	err = json.Unmarshal(j, &usersRestoredAdditional)
	if err != nil {
		fmt.Println("error on decode")
	}
	fmt.Println(usersRestoredAdditional)
}
