package main

func main() {
	println(Fact(6))
}

func Fact(fact int) (res int) {
	res = 1
	for i := 1; i <= fact; i++ {
		res *= i
	}

	return
}
