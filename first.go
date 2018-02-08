package main

import (
	"fmt"
)

type Human struct {
	name, lname string
	age         int
}

type I interface {
	printInfo()
	changeInfo()
}

func (h Human) printInfo() {
	fmt.Println(h.name, h.lname, h.age)
}

func (h *Human) changeInfo() {
	h.age = h.age + 20
}

func add(x, y int) int {
	return x + y
}

func mean(fr func(int, int) int, x, y int) float64 { //trochę jak wskaźniki w C++, dobre skojarzenie
	return float64(fr(x, y)) / 2
}

func main() {
	var i I
	slice := []Human{
		{name: "Barbara", lname: "Liwia", age: 42},
		{name: "Jan", lname: "Nowak", age: 50},
		{name: "Bonifacy", lname: "Kuchowski", age: 34}}

	i = &slice[1]
	i.printInfo()
	i.changeInfo()
	i.printInfo()

	human_map := make(map[int]Human)

	for i, y := range slice { //pierwszy zwróci numerek, kolejny wartość, jeśli coś musi być zwrócone ale jest namn niepotrzebne używamu "_"
		fmt.Println(y)
		human_map[i] = slice[i]
	}
	fmt.Println(human_map[1])

	meann := mean(add, 4, 99)
	fmt.Println(meann)
}
