package main

import "fmt"

// main is the entry point
func main(){
	fmt.Println("Hi from main")

	greeter1()
	greeter2()

	result := adder(3,5)
	fmt.Println(result)

	result1 := proadder(3,5,3,5,3,5)
	fmt.Println(result1)
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proadder(values ...int) int{
	total := 0
	for _, value := range(values){
		total = total + value
	}
	return total
}


func greeter1(){
	fmt.Println("He from greeter1")
}

func greeter2(){
	fmt.Println("He from greeter2")
}