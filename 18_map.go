package main

import "fmt"

func main() {
	//key dalam map bisa selain int, misal str
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Eko"
	mahasiswa["1002"] = "Kurniawan"
	mahasiswa["1003"] = "Khannedy"

	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1002"])
	fmt.Println(mahasiswa["1003"])
	//fmt.Println(mahasiswa["1004"])

	for nim, name := range mahasiswa {
		fmt.Println("Nim", nim, "bernama", name)
	}
}
