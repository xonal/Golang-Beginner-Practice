package main

import "fmt"

func main() {
	var positiveNumber uint8 = 80                            //deklarasi variabel dengan tipe data
	var negativeNumber = -332123                             //deklarasi variabel dengan tipe data
	tambahBilangan := positiveNumber + uint8(negativeNumber) //deklarasi variabel tanpa tipe data

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)  //%d digunakan untuk memformat data numerik non-desimal.
	fmt.Printf("Bilangan negative: %d\n", negativeNumber) //%d digunakan untuk memformat data numerik non-desimal.
	fmt.Println("Tambah bilangan: ", tambahBilangan)
}
