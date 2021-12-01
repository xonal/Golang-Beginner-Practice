package main

import "fmt"

func main() {
	// initializeSlice()
	// lenSclice()
	// lenCap()
}

func initializeSlice() {
	var fruits = []string{"apple", "mango", "orange", "pineapple"}

	fmt.Println(fruits[0])
}

func sliceLen() {
	var fruits = []string{"apple", "mango", "orange", "pineapple"}
	fmt.Println(len(fruits)) //len() digunakan untuk menghitung jumlah elemen slice yang ada
}

func sliceCap() {
	var fruits = []string{"apple", "mango", "orange", "pineapple"}
	var fruit = fruits[1:4]
	fmt.Println(cap(fruit)) //cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice

}

func sliceAppend() {
	var fruits = []string{"apple", "banana", "orange", "pineapple"}

	addFruits := append(fruits, "papaya") //append() digunakan untuk menambah elemen pada slice dibagian akhir

	fmt.Println(addFruits)
}
