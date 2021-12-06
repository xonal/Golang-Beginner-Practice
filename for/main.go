package main

import "fmt"

func main() {
	// loopingFor() //1

	// loopingForWithArgument() //2

	// loopingForNoArgument() //3

	// loopingForWithContinueBreak() //4

	// loopingNested() //5
}

func loopingFor() { //#1
	for i := 0; i < 5; i++ { //Basic for pada golang
		fmt.Printf("Angka %d \n", i) //Menampilkan angka dari 0 sampai 4
	}
}

func loopingForWithArgument() { //#2
	var i = 0 //Membuat variabel i dengan nilai 0

	for i < 5 { //Membuat perulangan dari 0 sampai 4
		fmt.Println("Angka ", i) //Menampilkan tulisan "Angka" dari 0 sampai 4
		i++                      //Menambahkan nilai i sampai 5
	}
}

func loopingForNoArgument() { //#3
	var i = 0 //Membuat variabel i

	for { //Perulangan tanpa ada argumen diperlukan if-statement
		fmt.Println("Angka", i) //Menampilkan tulisan "Angka" dari 0 hingga 9
		i++                     //Menambahkan nilai variabel i hingga 10
		if i == 10 {            //If-argument berfungsi agar tidak terjadi infinite-loop
			break //Membuat "break" agar looping-statement dapat dihentikan
		}
	}
}

func loopingForWithContinueBreak() { //#4
	//Perulangan dengan continue atau melewatkan value yang ingin ditampilkan

	for i := 0; i <= 10; i++ { //Membuat statement-looping
		if i%2 == 0 { //Membuat if-statement untuk nilai yang habis dibagi dengan 2
			continue //Continue digunakan untuk melewati bilangan yang habis dibagi dengan 2
			//Sehingga akan muncul bilangan ganjil
		}

		if i > 8 { //Membuat if-statement jika nilai i lebih besar dari 8
			break //Maka looping-statement dihentikan
		}

		fmt.Println("Angka ", i) //Menampilkan tulisan "Angka" dengan bilangan ganjil 1,3,5, dan 7
	}
}

func loopingNested() { //#5
	//Perulangan didalam perulangan
	for i := 0; i < 5; i++ { //Membuat statement-looping parent
		for j := 1; j < 5; j++ { //Membuat statement-looping child
			fmt.Println(j, "") //Menampilkan nilai dari hasil looping
		}
	}

	fmt.Println()
}
