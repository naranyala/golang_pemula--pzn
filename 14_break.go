package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		if i%2 == 1 {

			continue
		}

		fmt.Print(i)

		if i == 50 || i == 51 {
			break
		}

		//if i%2 == 1 {
		if i%2 == 0 {
			fmt.Println(" Genap")
		}
		//}
		//? continue berfungsi untuk melewati (skip) kondisi yg di deklarasikan

		//? break berfungsi untuk menghentikan paksa proses yg berjalan sesuai kondisi yg diinginkan
	}
}
