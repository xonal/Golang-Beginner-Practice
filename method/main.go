package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 = student{"John wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)
}

type student struct {
	name string
	age  int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, "")[i-1]
}
