package main

import "fmt"

func main() {
	fmt.Println("Belajar Golang")
	fmt.Println("Belajar" + " " + "Go")

	//menghitung jumlah karakter
	fmt.Println(len("Belajar"))

	fmt.Println("Belajar"[0])
	fmt.Println("Belajar"[0:4])
	//fmt.Println("Belajar"[1:4])
	fmt.Println("Belajar"[4:])
	fmt.Println("Belajar"[:4])
}
