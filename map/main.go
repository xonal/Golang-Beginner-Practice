package main

import "fmt"

func main() {
	// initializeMap()
	// intMap()
	// loopingMap()
	// deleteMap()
	foundItemKey()
	// sliceAndMap()
}

func initializeMap() {
	chicken := map[int]string{}

	chicken[10] = "Ando"
	chicken[20] = "Siska"

	fmt.Println("Januari:", chicken[10])
}

func intMap() {
	monthly := map[string]int{}

	monthly["januari"] = 20
	monthly["februari"] = 40
	monthly["maret"] = 100

	fmt.Println("Januari: ", monthly["januari"])
	fmt.Println("Februari: ", monthly["februari"])
	fmt.Println("Maret: ", monthly["maret"])

}

func loopingMap() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 200,
		"maret":    100,
		"april":    150,
	}

	for key, val := range chicken {
		fmt.Println(key, ": ", val)
	}
}

func deleteMap() {
	var chicken = map[string]int{
		"januari":  20,
		"februari": 50,
	}

	// delete(chicken, "januari")

	fmt.Println(chicken["januari"])
}

func foundItemKey() {
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

	if isExs {
		fmt.Println(val)
	} else {
		fmt.Println("item is not exists") //"item is not exists" akan ditampilkan karena bernilai false
	}
}

func sliceAndMap() {
	var students = [...]map[string]string{ //Slice and Map adalah kombinasi untuk data array berisi lebih dari satu baris data
		map[string]string{"name": "abdul", "gender": "male"},
		map[string]string{"name": "bila", "gender": "female"},
		map[string]string{"name": "citra", "gender": "female"},
		map[string]string{"name": "desi", "gender": "female"},
	}

	for _, val := range students {
		if val["gender"] == "male" {
			// fmt.Println(val["name"])
			continue
		} else {
			fmt.Println(val["name"])
		}
	}
}
