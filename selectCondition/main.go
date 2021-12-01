package main

import "fmt"

func main(){
	// var point = 10

	// if point == 10{
	// 	fmt.Println("you got best score!!")
	// }else if point == 9{
	// 	fmt.Println("you got high score!!")
	// }else if point == 8{
	// 	fmt.Println("you got good score!!")
	// }

	// temporaryIf()
	// secondTemporaryIf()
	switchCase()
}

func temporaryIf(){
	var point = 189.7

	if percent := point / 100; percent > 100{
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	}else if percent >= 70{
		fmt.Printf("%.1f%s good\n", percent, "%")
	}else{
		fmt.Printf("%.1f%s bad\n", percent, "%")
	}
}

func secondTemporaryIf(){
	if point := 65; point >= 100{
		fmt.Println("Perfect!!")
	}else if point >= 70{
		fmt.Println("Good!!")
	}else{
		fmt.Println("Bad")
	}
}

func switchCase(){
	var point = 10

	switch point{
		case 5:
			fmt.Println("Bad")
		
		case 7:
			fmt.Println("Good")
			
		case 10:
			fmt.Println("Perfect!!")

	}
}