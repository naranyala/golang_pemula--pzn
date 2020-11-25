package main

import "fmt"

//operasi matematika dan logika
func main() {
	fmt.Print("Hasil dari 1 + 1 = ")
	fmt.Println(1 + 1)
	fmt.Print("Hasil dari 10 - 1 = ")
	fmt.Println(10 - 1)
	fmt.Print("Hasil dari 10 * 2 = ")
	fmt.Println(10 * 2)
	fmt.Print("Hasil dari 10 / 2 = ")
	fmt.Println(10 / 2)
	fmt.Print("Hasil dari 10.0 / 3.0 = ")
	fmt.Println(10.0 / 3.0)
	fmt.Print("Hasil dari 10 % 2 = ")
	fmt.Println(10 % 2)
	fmt.Print("Hasil dari 11 % 2 = ")
	fmt.Println(11 % 2)

	fmt.Println("---")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
