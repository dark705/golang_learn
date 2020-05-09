package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run main.go -in=1000 -str="somestring" -in2=1234 -str2="string"

	//first usage
	some1 := flag.Int("in", 10, "")
	some2 := *flag.String("str", "some", "")

	//second usage
	var in2 int
	var str2 string
	flag.IntVar(&in2, "in2", 55, "")
	flag.StringVar(&str2, "str2", "some2", "")

	flag.Parse()
	fmt.Println(*some1, some2, in2, str2)

}
