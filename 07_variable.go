package main

import "fmt"

func main() {
	//var x string
	var hello string = "Hello World"

	fmt.Println(hello)
	fmt.Println(hello)
	fmt.Println(hello)
	//fmt.Println("Hello World")
	//fmt.Println("Hello World")
	//fmt.Println("Hello World")

	hello = "Hello Golang"
	fmt.Println(hello)
	fmt.Println(hello)
	fmt.Println(hello)

	fmt.Println("---")
	var nama string
	//mencetak variabel kosong
	fmt.Println(nama)
	nama = "Fudzer"
	fmt.Println(nama)

	namaLengkap := "Fudzer Mifthakul Huda"
	umur := 21
	fmt.Println(namaLengkap)
	fmt.Print("Umur = ")
	fmt.Println(umur)

	var nilai int32 = 75
	nilai2 := 80
	fmt.Println(nilai)
	fmt.Println(nilai2)

	fmt.Println("---")
	const perusahaan string = "Naranyala Dev"
	fmt.Println(perusahaan)

	//perusahaan = "Tulis Naranyala"
}
