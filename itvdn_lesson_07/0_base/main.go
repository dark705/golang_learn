package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//sample 1
	file, err := os.Open("some.json")
	failOnError(err)
	stat, err := file.Stat()
	failOnError(err)
	buf := make([]byte, stat.Size())
	_, err = file.Read(buf)
	failOnError(err)
	err = file.Close()
	failOnError(err)

	fmt.Println(string(buf))

	var users []User
	err = json.Unmarshal(buf, &users)
	failOnError(err)
	fmt.Println(users)

	//sample2
	file2, err := os.Open("some.json")
	failOnError(err)
	var users2 []User
	err = json.NewDecoder(file2).Decode(&users2)
	failOnError(err)
	err = file2.Close()
	failOnError(err)
	fmt.Println(users2)

}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
