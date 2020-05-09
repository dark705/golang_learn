package main

import "fmt"

type Plant struct {
	name  string
	color string
}

type Flower struct {
	petals int
	name   string
}

type Rouse struct {
	Plant
	Flower
	age int8
}

func (f Flower) ShowInfoFlower() {
	fmt.Printf("%+v\n", f)
}

func (r Rouse) ShowInfoRouse() {
	fmt.Printf("%+v\n", r)
}

func main() {
	rouse1 := Rouse{
		Plant:  Plant{"plantRouse", "red"},
		Flower: Flower{8, "flowerRouse"},
		age:    0,
	}

	fmt.Println(rouse1.Plant.name) //надо однозначно указывать, т.к. своййство name есть и в Plant и в Flower
	fmt.Println(rouse1.Flower.name)
	fmt.Println(rouse1.petals) // будет унаследованно от Flower, неоднозначности не будет, т.к. в Plant такого свойства нет
	rouse1.ShowInfoFlower()    //будут показаны свойства определённые только в Flower, не смотря на то что метод вызван на rouse
	rouse1.ShowInfoRouse()
}
