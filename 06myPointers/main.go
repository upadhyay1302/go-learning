package main

import "fmt"

func main(){
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)


}
