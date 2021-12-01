package main

import "fmt"

func main() {
	var nama string = "Rizal" //variable dengan tipe data
	umur := 18                //variable tanpa tipe data (pakai :=)

	fmt.Println("Nama: ", nama) //print nama
	fmt.Println("Umur: ", umur) //print umur
	fullName()
	multipleVariable()
	underscoreVariable()
}

func fullName(){
	var firstName string = "john"

	var lastName string

	lastName = "wick"
	fmt.Printf("Halo %s %s", firstName, lastName) //fmt.Printf() mengubah karakter tertentu seperti %s
	fmt.Print("\n")
}

func multipleVariable(){
	var first, second, third string //multiple variable
	first, second, third = "satu", "dua", "tiga"

	fmt.Println(first, second, third)
}

func underscoreVariable(){
	name, _ := "niken", "gurerro" // _ variable adalah variable yang tidak wajib dipanggil
								  // sehingga bisa tetap ditulis walaupun tidak di eksekusi
	fmt.Println(name)
}
