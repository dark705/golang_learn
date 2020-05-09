package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	bytes, err := ioutil.ReadFile("somefile.txt")
	failOnError(err)
	fmt.Println(string(bytes))

	timeAsString := []byte(time.Now().Local().String())
	err = ioutil.WriteFile("time.txt", timeAsString, 0666)
	failOnError(err)

}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
