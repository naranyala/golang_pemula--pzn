package main

import "fmt"

func main() {
	//mahasiswa := map[string]string{
	//	"1001": "Eko",
	//	"1002": "Kurniawan",
	//	"1003": "Khannedy",
	//}

	mahasiswa := map[string]map[string]string{
		"1001": map[string]string{
			"name":    "Eko",
			"address": "Indonesia",
			"gender":  "Male",
		},
		"1002": map[string]string{
			"name":    "Kurniawan",
			"address": "Indonesia",
			"gender":  "Male",
		},
		"1003": map[string]string{
			"name":    "Khannedy",
			"address": "Indonesia",
			"gender":  "Male",
		},
	}

	fmt.Println(mahasiswa)
	fmt.Println("---")
	fmt.Println(mahasiswa["1001"])
	fmt.Println("---")
	fmt.Println(mahasiswa["1001"]["name"])

	delete(mahasiswa, "1003")
	fmt.Println(mahasiswa)
}
