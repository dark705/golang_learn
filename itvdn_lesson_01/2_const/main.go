package main

func main() {
	const someConst int64 = 55

	const (
		some1 int32  = 23
		some2 string = "Some value"
		some3        = 223
	)
	println(someConst, some1, some2, some3)
}
