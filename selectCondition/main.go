package main

import "fmt"

func main() {
	// basicSelectCondition() //1

	// temporaryIf() //2

	// secondTemporaryIf() //3

	// switchCase() //4
}

func basicSelectCondition() { //#1
	var point = 10 //Membuat variabel "point" dengan nilai 10

	if point == 10 { //Jika nilai point sama dengan 10
		fmt.Println("you got best score!!") //Maka tampilkan "you got best score!!"
	} else if point == 9 { //Jika nilai point sama dengan 9
		fmt.Println("you got high score!!") //Maka tampilkan "you got hight score!!"
	} else if point == 8 { //Jika nilai point sama dengan 8
		fmt.Println("you got good score!!") //Maka tampilkan "you got good score!!"
	}
}

func temporaryIf() { //#2
	//Kondisi temporary digunakan untuk dapat langsung mengeksekusi kondisi tertentu pada line yang sama di "if-else"
	var point = 189.7 //Membuat variabel "point" dengan nilai 189.7

	if percent := point / 100; percent > 100 { //Dapat langsung membuat variabel "percent" dengan nilai point/100 dan dibandingkan apakah percent lebih besar dari 100
		fmt.Printf("%.1f%s perfect\n", percent, "%") //Tampilkan nilai percent
	} else if percent >= 70 { //Jika percent lebih besar sama dengan dari 70
		fmt.Printf("%.1f%s good\n", percent, "%") //Tampilkan nilai percent
	} else { //Jika tidak
		fmt.Printf("%.1f%s bad\n", percent, "%") //Tampilkan nilai percent
	}
}

func secondTemporaryIf() { //#3
	//Kondisi temporary digunakan untuk dapat langsung mengeksekusi kondisi tertentu pada "if-else"
	if point := 65; point >= 100 { //Dapat langsung membuat variabel "point" dengan nilai 65 dan langsung dibandingkan pada line yang sama
		fmt.Println("Perfect!!") //Tampilkan "Perfect"
	} else if point >= 70 { //Jika nilai "point" lebih besar sama dengan 70
		fmt.Println("Good!!") //Tampilkan "Good!!"
	} else { //Jika tidak
		fmt.Println("Bad") //Tampilkan "Bad"
	}
}

func switchCase() { //#4
	var point = 10 //Membuat variabel point dengan nilai 10

	//Switch case adalah suatu bentuk kondisi percabangan dimana jika suatu kondisi benar

	switch point { //membuat switch
	case 5: //Jika kondisi "point" bernilai 5
		fmt.Println("Bad") //Tampilkan "Bad"
	case 7: //Jika kondisi "point" bernilai 7
		fmt.Println("Good") //Tampilkan "Good"
	case 10: //Jika kondisi "point" bernilai 10
		fmt.Println("Perfect!!") //Tampilkan "Perfect!!"
	}
}
