package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	b := new(bytes.Buffer)
	for i := 0; i <= 5; i++ {
		b.WriteString(fmt.Sprintf("Time: %s\n", time.Now().String()))
		time.Sleep(time.Millisecond * 200)
	}
	file, err := os.Create("time.txt")
	defer file.Close()
	failOnError(err)

	writeBytes, err := io.Copy(file, b)
	fmt.Printf("Write bytes: %d", writeBytes)

}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
