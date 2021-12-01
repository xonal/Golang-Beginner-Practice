package main

import "fmt"

func main() {
	constFunction() //1
}

func constFunction() { //#1
	const firstName string = "john"        //const adalah variable yang tidak bisa diubah(tetap)
	const lastName = "wick"                //teknik interface pada konstanta
	fmt.Print("halo", firstName, lastName) //fmt.Print tidak memberikan baris baru seperti fmt.Println()
}
