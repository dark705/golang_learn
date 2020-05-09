package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	Name   string `json:"SomeName"`
	Age    uint8  `xml:"ageXml" json:"ageJson"`
	Gender string
}

func main() {
	user1 := User{
		"Vasya",
		20,
		"M",
	}

	b1, error1 := json.Marshal(&user1)
	b2, error2 := xml.Marshal(&user1)

	if error1 != nil && error2 != nil {
		panic("Some wrong")
	}

	fmt.Println(user1)
	fmt.Println(string(b1))
	fmt.Println(string(b2))
}
