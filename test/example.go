package test

import "fmt"

func SayHello() {
	fmt.Println("Hello world")
}

func SayGoodBye() {
	fmt.Println("Good")
	fmt.Println("Bye")
}

func PrintNames() {
	students := make(map[int]string, 4)
	students[0] = "bob"
	students[1] = "sari"
	students[2] = "emm"
	students[3] = "susan"
	for _, v := range students {
		fmt.Println(v)
	}
}
