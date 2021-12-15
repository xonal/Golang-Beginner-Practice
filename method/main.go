package main

import (
	"fmt"
	"strings"
)

func main() {
	/*Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya).
	Method bisa diakses lewat variabel objek.*/
	//Perbedaan method get dan set adalah dari awal penulisan method
	//Method set tidak mengembalikan nilai, oleh karena itu tidak dituliskan tipe data di akhir nama fungsi
	//Method get mengembalikan nilai, oleh karena itu dituliskan tipe data diakhir
	//Contoh method set ada di #set-1
	//Contoh method get ada di #get-1 dan #get-2

	runMethod() //1
}

func runMethod() { //#1
	var s1 = student{"John wick", 21} //Membuat variabel s1 dengan nilai struct  "student" dengan mengisi value "john wick" dan "21"
	s1.sayHello()                     //Memanggil method sayHello()

	var name = s1.getNameAt(2)   //Membuat variabel name dengan nilai method get-getNameAt(2). "2" adalah parameter dari fungsi getter
	fmt.Println("huruf :", name) //Menampilkan huruf dari nama "John wick"

	var fullName = s1.getName() //Membuat variabel fullName dengan nilai method getName()
	fmt.Println(fullName)       //Menampilkan variabel fullName
}

type student struct { //Membuat struct dengan nama student
	name string //Membuat variabel name dengan tipe data string
	age  int    //Membuat variabel age dengan tipe data int
}

func (s student) sayHello() { //#set-1 //Contoh penulisan method diawali dengan inisialisasi struct dan kemudian nama method
	fmt.Println("halo", s.name) //Menampilkan tulisan "halo" ditambah nama dari variabel struct
}

func (s student) getNameAt(i int) string { //#get-1 //Contoh penulisan method untuk mengembalikan nama dari struct student
	return strings.Split(s.name, "")[i-1] //Mengembalikan satu karakter dari nama "student" berupa string
}

func (s student) getName() string { //#get-2
	return s.name //Mengembalikan nama "student" berupa string
}
