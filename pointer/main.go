package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA //Membuat variable dengan tipe data pointer yang menampung nilai pointer dari numberA

	fmt.Println(&numberA) //Mengambil nilai pointer pada variable biasa
	fmt.Println(numberB)  //Menampilkan variableB yang menampung nilai pointer numberA
}
