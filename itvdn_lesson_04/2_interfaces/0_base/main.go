package main

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

type Animal interface {
	SayYourName()
	SaySound()
}

func (c Cat) SayYourName() {
	println(c.Name)
}

func (c Cat) SaySound() {
	println("Myau")
}

func (d Dog) SayYourName() {
	println(d.Name)
}

func (d Dog) SaySound() {
	println("Gav")
}

func main() {
	cat1 := Cat{"Murka"}
	dog1 := Dog{"Reks"}
	CheckAnimal(cat1)
	CheckAnimal(dog1)
}

func CheckAnimal(a Animal) {
	a.SaySound()
	a.SayYourName()
}
