package main

import "fmt"

func main() {
	//map merupakan suatu fungsi key-value
	//Perbedaan map dan array adalah pada key yang mana key pada array tidak dapat bertipe "string"

	// initializeMap() //1

	// intMap() //2

	// loopingMap() //3

	// deleteMap() //4

	// foundItemKey() //5

	// sliceAndMap() //6
}

func initializeMap() { //#1
	chicken := map[int]string{} //membuat fungsi map dengan variabel bernama chicken

	chicken[10] = "Ando"  //key-chicken[10] berisi nilai "Ando"
	chicken[20] = "Siska" //key-chicken[20] berisi nilai "Siska"

	fmt.Println("Januari:", chicken[10]) //Menampilkan data dari key-chicken[10]
}

func intMap() { //#2
	monthly := map[string]int{} //Membuat map dengan variabel bernama monthly

	monthly["januari"] = 20  //key-monthly[januari] berisi nilai 20
	monthly["februari"] = 40 //key-monthly[februari] berisi nilai 40
	monthly["maret"] = 100   //key-monthly[maret] berisi nilai 100

	fmt.Println("Januari: ", monthly["januari"])   //Menampilkan data dari key-monthly[januari]
	fmt.Println("Februari: ", monthly["februari"]) //Menampilkan data dari key-monthly[februari]
	fmt.Println("Maret: ", monthly["maret"])       //Menampilkan data dari key-monthly[maret]

}

func loopingMap() { //#3
	var chicken = map[string]int{ //membuat map dengan variabel bernama chicken
		"januari":  50,  //key-chicken[januari] berisi nilai 50
		"februari": 200, //key-chicken[februari] berisi nilai 200
		"maret":    100, //key-chicken[maret] berisi nilai 100
		"april":    150, //key-chicken[april] berisi nilai 150
	}

	for key, val := range chicken { //Melakukan perulangan key-value pada variabel chicken
		fmt.Println(key, ": ", val) //Menampilkan hasil perulangan pada variabel chicken
	}
}

func deleteMap() { //#4
	var chicken = map[string]int{ //Membuat variabel chicken dengan tipe data map
		"januari":  20,
		"februari": 50,
	}

	// delete(chicken, "januari") //Menghapus nilai "januari" pada map-chicken

	fmt.Println(chicken["januari"]) //Menampilkan nilai chicken["januari"]
}

func foundItemKey() { //#5
	var month = map[int]string{1: "januari", 2: "februari"}
	// var value, isExist = month[1] //Terdapat dua variable yaitu value dan isExists
	//value adalah nilai dari month yaitu "januari"
	//isExist adalah nilai boolean "true"

	var val, isExs = month[3] //isExs adalah nilai boolean "false" karena tidak ada map di bulan ke 3

	// if isExist{
	// 	fmt.Println(value)
	// }else{
	// 	fmt.Println("item is not exists")
	// }

	if isExs { //Mengecek apakah isExs bernilai "true" atau "false"
		fmt.Println(val) //Jika bernilai true maka menampilkan nilai pada variabel "val"
	} else { //Jika "false"
		fmt.Println("item is not exists") //"item is not exists" akan ditampilkan karena bernilai false
	}
}

func sliceAndMap() { //#6
	var students = [...]map[string]string{ //Slice and Map adalah kombinasi untuk data array berisi lebih dari satu baris data
		map[string]string{"name": "abdul", "gender": "male"},
		map[string]string{"name": "bila", "gender": "female"},
		map[string]string{"name": "citra", "gender": "female"},
		map[string]string{"name": "desi", "gender": "female"},
	}

	for _, val := range students { //Melakukan perulangan key-value pada variabel students
		if val["gender"] == "male" { //Mengecek apakah perulangan val["gender"] berada di posisi "male"
			// fmt.Println(val["name"])
			continue //Jika "true" maka gender dengan data "male" akan di longkap atau skip
		} else {
			fmt.Println(val["name"]) //Jika "false" maka akan menampilkan nilai gender yang lain
		}
	}
}
