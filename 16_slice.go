package main

import "fmt"

func main() {
	names := [5]string{
		"Eko",
		"Budi",
		"Joko",
		"Rudi",
		"Brian",
	}

	fmt.Println(names)
	fmt.Println(names[1])

	//slice atau potongan
	slice1 := names[1:4]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	names[1] = "Nugraha"
	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice1[1] = "Moro"
	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//slice (ibarat pointer) dmn data ttp mengikuti array
	//hati2 mengubah slice dan array
}
