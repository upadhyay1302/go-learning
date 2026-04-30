package main

import "fmt"

// Defer statements - delays the statement and puts it in the bottom  
// When we have mutiple defer statements we follow LIFO fo the defer statements, the non defer statements are normal execution
func main(){
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i:=0; i <5;i++{
		defer fmt.Println(i)
	}
}
