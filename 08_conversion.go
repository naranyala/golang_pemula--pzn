package main

import (
	"fmt"
	"strconv"
)

func main() {
	//konversi
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(x)
	fmt.Println(y)

	var z float64 = float64(y)
	fmt.Println(z)

	var b float64 = 3.9

	fmt.Println(b)
	var a int32 = int32(b)
	fmt.Println(a)
	fmt.Println("---")

	var nilai string = "100"
	//var nilaiInt int = strconv.Atoi(nilai)

	//return ke 2 variabel
	//nilaiInt, error := strconv.Atoi(nilai)

	//variable kedua akan dilewati
	nilaiInt, _ := strconv.Atoi(nilai)
	fmt.Println(nilaiInt)

	nilaiString := strconv.Itoa(-100)
	fmt.Println(nilaiString)
}
