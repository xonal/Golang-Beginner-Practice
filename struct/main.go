package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct { //Mendefinisikan model struct dengan nama "student"
	name  string //Membuat properti "name" dengan tipe data string
	age   int    //Membuat properti "age" dengan tipe data int
	grade string //Membuat properti "grade" dengan tipe data string
	//Embedded struct yaitu menempelkan sebuah struct sebagai properti pada struct lain
	// person //Contoh embedded struct

}

func main() {
	/*Struct adalah kumpulan definisi variabel (atau property)
	dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru
	dengan nama tertentu. */
	// basicStruct() //1
	// objectStruct() //2
	// anonymouseStruct() //3
}

func basicStruct() { //#1
	var std student          //Membuat variabel std dengan tujuan memanggil struct "student"
	std.name = "John"        //Memberikan nilai properti name pada struct student dengan "John"
	std.age = 22             //Memberikan nilai properti age pada struct student dengan "22"
	std.grade = "University" //Memberikan nilai properti grade pada struct student dengan "University"

	fmt.Println("Name :", std.name)   //Menampilkan properti "Name"
	fmt.Println("Age :", std.age)     //Menampilkan properti "Age"
	fmt.Println("Grade :", std.grade) //Menampilkan properti "Grade"

}

func objectStruct() { //#2
	var s1 = student{}      //Membuat object struct kosong
	s1.name = "Abdul"       //Memberikan nilai properti name pada struct student dengan "Abdul"
	s1.age = 23             //Memberikan nilai properti age pada struct student dengan "23"
	s1.grade = "University" //Memberikan nilai properti grade pada struct student dengan "University"

	var s2 = student{"Johan", 21, "University"} //Membuat object student dengan nama variabel s2
	var s3 = student{name: "Michael"}           //Membuat object student dengan nama variabel s3

	fmt.Println("Student 1 : ", s1.name) //Menampilkan nama student 1
	fmt.Println("Student 2 : ", s2.name) //Menampilkan nama student 2
	fmt.Println("Student 3 : ", s3.name) //Menampilkan nama student 3
}

func anonymouseStruct() { //#3
	/*Anonymouse struct adalah struct yang tidak dideklarasikan di
	awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek. */
	//Anonymouse struct dapat dilakukan untuk pembuatan variabel objek yang hanya dipakai sekali
	var person = struct { //Membuat anonymouse struct dengan nama person
		name string //membuat properti name
		age  int    //Membuat properti age
	}{}

	person.name = "Agung" //Mengisi nilai properti name dengan "Agung"
	person.age = 25       //Mengisi nilai properti age dengan 25

	fmt.Println("name : ", person.name) //Menampilkan properti name
	fmt.Println("Age : ", person.age)   //Menampilkan properti age
}

func sliceAndStruct() {
	var allStudents = []person{ //Membuat variabel allStudents berbentuk struct dan slice
		{name: "Abdul", age: 25},
		{name: "Desi", age: 23},
	}

	for _, student := range allStudents { //Melakukan perulangan pada allStudents
		//_ merupakan simbol key pada perulangan allStudents
		//student merupakan simbol value pada perulangan allStudents
		fmt.Println(student.name, "age is", student.age) //Menampilkan nama dan umur
	}

}
