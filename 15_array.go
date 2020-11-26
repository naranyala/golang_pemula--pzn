package main

import "fmt"

func main() {

	//array atau larik
	var names [5]string

	names[0] = "Fudzer"
	names[1] = "Huda"
	names[2] = "Rizal"
	fmt.Println(names)
	fmt.Println(names[0] + " " + names[1])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println("---")

	//for i := 0; i < 5; i++ {
	for i := 0; i < len(names); i++ {
		fmt.Print(i, " ")
		fmt.Println(names[i])
	}
	fmt.Println("---")
	//for index, value := range names {
	for _, value := range names {
		//fmt.Println("index", index, "=", value)
		fmt.Println(value)
	}
}
