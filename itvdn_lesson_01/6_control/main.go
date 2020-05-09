package main

func main() {
	a := true
	var b bool
	str := "some"
	sl := []string{"one", "two", "three", "four"}

	if a {
		println("a is true")
	}

	if b == false {
		println("b is false")
	}

	if len(str) > 4 {
		println("str is bigger then 4")
	} else if len(str) == 4 {
		println("str is equals 4")
	} else {
		println("str is less then 4")
	}

	//infinity loop
	for {
		println("infinity loop")
		break
	}

	for i := 0; i < len(sl); i++ {
		if i == 2 {
			continue
		}
		println(sl[i])
	}

	//итерация по циклу
	for index, value := range sl {
		println(index, ":", value)
	}

	someSwitchVar := 2
	switch someSwitchVar {
	case 1:
		println("one") //break не нужен
	case 2:
		println("two")
		fallthrough //если хотим провалиться ещё ниже
	default:
		println("default case")
	}

	var arr [2][2]string
	arr[0][0] = "a11"
	arr[0][1] = "a12"
	arr[1][0] = "a21found"
	arr[1][1] = "a22"

SomeMark:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			println(arr[i][j])
			if arr[i][j] == "a21found" {
				println("Found!!!")
				break SomeMark
			}
		}
	}

SomeMark2:
	for indexI, valueI := range arr {
		for indexJ, _ := range valueI {
			println(arr[indexI][indexJ])
			if arr[indexI][indexJ] == "a21found" {
				println("Found!!!")
				break SomeMark2
			}
		}
	}
}
