package main

import "fmt"

func main(){
	// loopingFor()
	// loopingForWithArgument()
	// loopingForNoArgument()
	// loopingForWithContinueBreak()
	loopingNested()
}

func loopingFor(){
	for i := 0; i < 5; i++{
		fmt.Printf("Angka %d \n", i)
	}
}

func loopingForWithArgument(){
	var i = 0

	for i < 5{
		fmt.Println("Angka ", i)
		i++
	}
}


func loopingForNoArgument(){
	var i = 0

	for{
		fmt.Println("Angka", i)
		i++
		if i == 10{
			break
		}
	}
}

func loopingForWithContinueBreak(){
	for i := 0; i <= 10; i++{
		if i % 2 == 0{ //Saat i bernilai 2 maka tidak di Print
			continue
		}

		if i > 8{ //Saat i bernilai 8 maka looping di hentikan
			break
		}

		fmt.Println("Angka ", i)
	}
}

func loopingNested(){
	for i := 0; i < 5; i++{
		for j := 1; j < 5; j++{
			fmt.Println(j, "")
		}
	}

	fmt.Println()
}