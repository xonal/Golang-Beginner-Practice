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
	var point = 189.7 //Membuat variabel "point" dengan nilai 189.7

	if percent := point / 100; percent > 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s bad\n", percent, "%")
	}
}

func secondTemporaryIf() { //3
	if point := 65; point >= 100 {
		fmt.Println("Perfect!!")
	} else if point >= 70 {
		fmt.Println("Good!!")
	} else {
		fmt.Println("Bad")
	}
}

func switchCase() { //#4
	var point = 10

	switch point {
	case 5:
		fmt.Println("Bad")

	case 7:
		fmt.Println("Good")

	case 10:
		fmt.Println("Perfect!!")

	}
}
