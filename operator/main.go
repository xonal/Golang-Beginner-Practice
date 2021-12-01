package main

import "fmt"

func main() {
	operatorBasic() //1
}

func operatorBasic() { //#1
	var value int = (((2+6)%3)*4 - 2) / 3 //Membuat variabel value dengan tipe data int
	var isEqual = (value == 2)            //jika value sama dengan 2

	fmt.Printf("nilai %d (%t) \n", value, isEqual) //Menampilkan nilai value dan isEqual
}
