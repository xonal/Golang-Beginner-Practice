package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func main() {
	// var names = []string{"abi", "adi"} //array dinamis berisi string
	// printMessages("halo ", names)      //panggil fungsi "printMessages"
	// closure()

	// var minMax = singleReturn(50, 30)
	// fmt.Println(minMax)

	stopProcesswithReturn(2, 0)
	
}

func printMessages(message string, arr []string) { //parameter terdapat string dan array string
	var nameString = strings.Join(arr, " ") // menghilangkan [] dan menambahkan spasi setiap beda kata
	fmt.Print(message, nameString)          //print parameter satu dan dua

	//hasil "halo abi adi"
}

func singleReturn(min, max int) int{
	var value = rand.Int() % (max - min + 1) + min 
	return value
}

func closure(){ //Closure function adalah fungsi didalam fungsi
	var  getMinMax = func(n []int) (int, int) { //Contoh closure function yaitu dengan membuat variable yang menampung sebuah function
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers) //Buat dua variable dengan nama min dan max karena fungsi mengembalikan nilai dua "return"
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}

func stopProcesswithReturn(m, n int){
	if n == 0{
		fmt.Printf("Invalid divider. %d cannot divide by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

