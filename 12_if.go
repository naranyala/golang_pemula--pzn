package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		//fmt.Println(i)
		if i%2 == 1 {
			fmt.Println(i, "Ganjil")
		} else if i%2 == 1 {
			fmt.Println("Ganjil lagi")
		} else if i%2 == 1 {
			fmt.Println("Ganjil lagi lagi")
		} else {
			//if i%2 == 0 {
			fmt.Println(i, "--Genap")
		}
	}
}
