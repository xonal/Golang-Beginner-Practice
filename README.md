# Golang-Beginner-Practice (In Bahasa)

<img src="https://user-images.githubusercontent.com/33995016/103326949-6f738400-4a07-11eb-9dd2-d39075489763.png" width="600" height="400">

Ini adalah catatan singkat Golang dasar mulai dari pembuatan variabel, array, pembuatan fugsi, dan lain sebagainya.
Dibuatnya Repositori ini adalah untuk memudahkan dalam mengingat syntax dasar Golang. Catatan ini didapatkan dari referensi blog https://dasarpemrogramangolang.novalagung.com/

## Instalasi
Download golang pada https://golang.org/ dan pilih file instalasi sesuai dengan sistem operasi yang saat ini sedang digunakan

## Init project
Inisialisasi projek yang dibuat untuk dapat melakukan instalasi
> C:\Users\user>go mod init <nama_project> 

## Variabel

```
var nama string = "johan" //menggunakan tipe data

city := "jakarta //tanpa tipe data
    
food, drink = "nasi goreng", "jus jeruk" //multiple variabel

_ := "pria" //reserved variabel (variabel yang dapat di deklarasi tanpa perlu dipanggil)
```

## Tipe Data

Tanda        | Penjelasan
------------ | -------------
uint8        | 0 ↔ 255
uint16       | 0 ↔ 65535
uint32       | 0 ↔ 4294967295
uint64       | 0 ↔ 18446744073709551615
uint         | sama dengan uint32 atau uint64 (tergantung nilai)
byte         | sama dengan uint8
int8         | -128 ↔ 127
int16        | -32768 ↔ 32767
int32        | -2147483648 ↔ 2147483647
int64        | -9223372036854775808 ↔ 9223372036854775807
int          | sama dengan int32 atau int64 (tergantung nilai)
rune         | sama dengan int32

```
var num int16 = 12 //variabel dengan int16

```

## Fungsi *fmt.Print()* dan *fmt.Println()*

```

package main

import "fmt"

func main() {

   fmt.print("Hello World") //menampilkan tulisan Hello World pada console
   
   fmt.println("Hello") //menampilkan tulisan Hello dengan new line
   
   fmt.println("World") //menampilkan tulisan World dengan new line
}

```

## Operator

```
Tanda | Penjelasan
----- | -------------
*     | Perkalian
/     | Pembagian
+     | Penjumlahan
-     | Pengurangan



Tanda | Penjelasan
----- | -------------
==    | Membandingkan kedua nilai antara **kiri** dan **kanan**
<=    | Kurang dari sama dengan
>=    | Lebih dari sama dengan
!=    | Tidak sama dengan


    if( 2 == 2) { //2 sama dengan 2
    
        fmt.Print("true") //print "true"
   
   }else if(2 <= 1) { // 2 lebih kecil sama dengan 1
    
        fmt.Print("false") //print "false"
   
   }else if(2 >= 1) { // 2 lebih besar sama dengan 1
    
        fmt.Print("true")//print "true"
   
   }else if( 2 != 1) {// 2 tidak sama dengan 1
    
        fmt.Print("true")// print "true"
   
   }

```

## Perbandingan Dengan *if*, *else-if*, *else*

Perbandingan **if**

```
package main

import "fmt"

func main() {

   if( 2 == 2) { //2 sama dengan 2
    
    fmt.Print("true") //print "true"
   
   }
}

```

Perbandingan **else if**
```
package main

import "fmt"

func main() {

   if( 2 == 2) { //2 sama dengan 2
    
    fmt.Print("true") //print "true"
   
   } else if ( 2 != 1) {
   
    fmt.Print("true") //print "true"
   
   }
}

```


