package main

func main() {
	var in int = -34
	var in64 int64 = -3345
	var autoInt = 333
	var biggerZero uint64 = 222
	println(in, in64, autoInt, biggerZero)

	var fl32 float32 = -23.6345
	println(fl32)

	var someBoolTrue bool = true
	var someBoolFlase bool = false
	println(someBoolTrue, someBoolFlase)

	var someString string = "Some integer"
	println(someString)

	//Краткое объявление
	someVariable := "Some value"
	println(someVariable)

	//Приведение типов
	println(int(fl32))

	//Множественное присваивание
	var var1, var2 string = "Some1", "Some2"
	println(var1, var2)

	var (
		sm1 int     = 12122
		sm2         = "stringggggg"
		sm3 float32 = 1.23444
	)

	println(sm1, sm2, sm3)
	println(string(sm1) + sm2)
}
