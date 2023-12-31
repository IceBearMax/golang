package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Human achieve SayHi method
func (h Human) SayHi() {
	fmt.Printf("Hi ,I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s ,I work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"paul", 26, "111-222-xxx"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-xxx"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-xxx"}, "Thing Ltd", 5000}

	var i Men

	i = mike
	fmt.Println("This is tom ,a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is tom,an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
		
	}

}
