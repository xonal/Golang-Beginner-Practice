package main

import (
	"fmt"
)

func main() {

	// defaultArray() //1

	// verticalHorizontalArray() //2

	// dynamicArray() //3

	// rangeArray() //4

	// multidimensionalArray() //5

	// makeArray() //6

	postTest()
}

func defaultArray() { //#1
	var names [4]string   //Membuat variabel names array dengan tipe data string
	names[0] = "rafathar" //names[0] berisi nama "rafathar"
	names[1] = "reza"     //names[1] berisi nama "reza"

	fmt.Println(names[0]) //Tampilkan names[0] berisi "rafathar"
	fmt.Println(names[1]) //Tampilkan names[1] berisi "reza"
}

func verticalHorizontalArray() { //#2
	var fruits [4]string //Membuat variabel fruits dengan array string dengan maksimum nilai 4

	fruits = [4]string{"apple", "banana"} //Penulisan array horizontal

	fruits = [4]string{ //Penulisan array vertical
		"apple",
		"banana",
		"mango",
		"orange",
	}

	last := len(fruits) - 1 //len adalah panjang array

	fmt.Println("Jumlah buah ada ", len(fruits))     //Menampilkan jumlah buah dengan variabel "fruits"
	fmt.Println("Buah pertama adalah", fruits[last]) //Menampilkan nama buah pertama

}

func dynamicArray() { //#3
	//Dynamic Array adalah array yang nilainya tidak dibatasi jumlah maksimum nilainya
	var numbers = [...]int{1, 2, 4, 6} //dynamic array ditulis menggunakan [...]

	for i := 0; i < len(numbers); i++ { //Melakukan perulangan
		fmt.Println("Angka", numbers[i]) //Menampilkan angka secara berurut
	}
}

func rangeArray() { //#4
	var names = [...]string{"Abdul", "Bisma", "Celine", "Desi"}
	for i, name := range names { //range adalah solusi perulangan untuk array
		// i(underscore) adalah index array (0, 1, 2, 3) dan name adalah index hasil dari array names ("Abdul", "Bisma", "Celine", "Desi")
		fmt.Println("Nama ke ", i, "Adalah", name)
	}
}

func multidimensionalArray() { //#5
	//MultidimensionalArray adalah array 2 dimensi
	numbers := [2][4]int{{3, 2, 3, 4}, {3, 4, 5, 6}}

	for i := 0; i < len(numbers); i++ { //Melakukan pengulangan untuk menampilkan masing-masing dari baris (looping baris)
		fmt.Println("Numbers: ", numbers[i]) //Menampilkan hasil perulangan baris
	}

	for i := 0; i < len(numbers); i++ { //Melakukan pengulangan baris
		for j := 0; j < len(numbers[i]); j++ { //Melakukan pengulangan kolom
			fmt.Println("Numbers: ", numbers[i][j]) //Menampilkan hasil perulangan nilai satuan dari masing-masing baris
		}

	}
	// fmt.Println("Numbers1:", numbers)
}

func makeArray() { //#6
	fruits := make([]string, 2) //fungsi make digunakan untuk membuat dynamic array
	fruits[0] = "apple"
	fruits[1] = "mango"

	fmt.Println(fruits)
}

func postTest() {
	const pers = 4

	var person [pers]string
	person[0] = "Adi"

	for i := len(person) - 1; i > 0; i-- {
		fmt.Println(person[i])
	}

}
