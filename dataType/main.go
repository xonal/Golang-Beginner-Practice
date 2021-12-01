package main

import "fmt"

func main() {
	dataType() //1
}

func dataType() { //#1
	var positiveNumber uint8 = 80                       //deklarasi variabel dengan tipe data uint8
	var negativeNumber = -332123                        //deklarasi variabel tanpa tipe data namun compiler secara cerdas langsung menentukan tipe data int32
	addNumber := positiveNumber + uint8(negativeNumber) //deklarasi variabel tanpa tipe data dengan nilai positiveNumber+NegativeNumber

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)  //%d digunakan untuk memformat data numerik non-desimal.
	fmt.Printf("Bilangan negative: %d\n", negativeNumber) //%d digunakan untuk memformat data numerik non-desimal.
	fmt.Println("Tambah bilangan: ", addNumber)           //Mencetak variabel "addNumber"
}
