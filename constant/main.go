package main

import "fmt"

func main(){
	const firstName string = "john" //const adalah variable yang tidak bisa diubah(tetap)
	const lastName = "wick" //teknik interface pada konstanta
	fmt.Print("halo", firstName, lastName) //fmt.Print tidak memberikan baris baru seperti fmt.Println()
}