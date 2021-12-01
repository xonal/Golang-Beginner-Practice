package main

import (
	"fmt"
)

func main() {
	fixedArray()
	dynamicArray()
	rangeArray()
	multidimensionalArray()
	defaultArray()
	makeArray()

}

func defaultArray() {
	var names [4]string
	names[0] = "rafathar"
	names[1] = "reza"

	fmt.Println(names[0])
	fmt.Println(names[1])
}

func fixedArray() {
	var fruits [4]string

	fruits = [4]string{
		"apple",
		"banana",
		"mango",
		"orange",
	}

	last := len(fruits) - 1 //len adalah panjang array

	fmt.Println("Jumlah buah ada ", len(fruits))
	fmt.Println("Buah pertama adalah", fruits[last])

}

func dynamicArray() {
	var numbers = [...]int{1, 2, 4, 6} //dynamic array ditulis menggunakan [...]

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Angka", numbers[i])
	}
}

func rangeArray() {
	var names = [...]string{"Abdul", "Bisma", "Celine", "Desi"}
	for i, name := range names { //range adalah solusi perulangan untuk array
		// i(underscore) adalah index array (0, 1, 2, 3) dan name adalah index hasil dari array names ("Abdul", "Bisma", "Celine", "Desi")
		fmt.Println("Nama ke ", i, "Adalah", name)
	}
}

func multidimensionalArray() {
	numbers := [2][4]int{{3, 2, 3, 4}, {3, 4, 5, 6}}
	fmt.Println("Numbers1:", numbers)
}

func makeArray() {
	fruits := make([]string, 2) //fungsi make digunakan untuk membuat dynamic array
	fruits[0] = "apple"
	fruits[1] = "mango"

	fmt.Println(fruits)
}
