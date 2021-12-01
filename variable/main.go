package main

import "fmt"

func main() {

	// basicVariable() //1

	// fullName() //2

	// multipleVariable() //3

	// underscoreVariable() //4
}

func basicVariable() { //#1
	var nama string = "Rizal" //Variabel nama dengan tipe data "String"
	umur := 20                //variable umur tanpa tipe data ditulis menggunakan tanda ":="

	fmt.Println("Nama: ", nama) //Mencetak nama
	fmt.Println("Umur: ", umur) //Mencetak umur
}

func fullName() { //#2
	var firstName string = "john" //Variabel firstname dengan tipe data "String"

	var lastName string //Variabel lastName dengan tipe data "String" tanpa nilai

	lastName = "wick"                             //Mengisi nilai pada variabel lastName
	fmt.Printf("Halo %s %s", firstName, lastName) //fmt.Printf() digunakan untuk mengubah karakter tertentu seperti %s

}

func multipleVariable() { //#3
	var first, second, third string              //Pembuatan banyak variabel dengan tipe data "String"
	first, second, third = "satu", "dua", "tiga" //Pengisian nilai pada variabel "first", "second", dan "third"

	fmt.Println(first, second, third) //Mencetak nilai "first", "second", dan "third"
}

func underscoreVariable() { //#4
	name, _ := "niken", "gurerro" // _ adalah variable yang tidak wajib dipanggil
	// sehingga bisa tetap ditulis walaupun tidak di eksekusi
	fmt.Println(name) //Cotoh variabel _ tidak dipanggil namun fungsi tetap bisa di eksekusi
}
