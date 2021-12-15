package main

import "fmt"

func main() {
	//Perbedaan slice dan array
	//Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.
	//Slice dibentuk dari array yang sudah didefinisikan

	/*Perbedaan Array dan Slice*/
	// var fruitsA	= []string{"apple", "grape"}      // slice
	// var fruitsB = [2]string{"banana", "melon"}    // array
	// var fruitsC = [...]string{"papaya", "grape"}  // array

	// initializeSlice() //1

	// lenSclice() //2

	// sliceCap() //3

	// lenCap() //4
}

func initializeSlice() { //#1
	var fruits = []string{"apple", "mango", "orange", "pineapple"} //Inisialisasi slice
	fmt.Println(fruits[0])                                         //Mencetak tulisan "apple"
}

func sliceLen() { //#2
	var fruits = []string{"apple", "mango", "orange", "pineapple"} //Inisialisasi slice
	fmt.Println(len(fruits))                                       //len() digunakan untuk menghitung jumlah elemen slice yang ada
}

func sliceCap() { //#3
	var fruits = []string{"apple", "mango", "orange", "pineapple"} //Inisialisasi slice
	var fruit = fruits[1:4]                                        //Mengakses elemen dari slice dari "mango" hingga "pineapple"
	fmt.Println(fruit)                                             //cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice

}

func sliceAppend() { //#4
	var fruits = []string{"apple", "banana", "orange", "pineapple"} //Inisialisasi slice
	addFruits := append(fruits, "papaya")                           //append() digunakan untuk menambah elemen pada slice dibagian akhir

	fmt.Println(addFruits) //Mencetak tulisan "papaya"
}
