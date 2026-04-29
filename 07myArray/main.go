package main

import "fmt"

func main(){
	var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"

	fmt.Println("Array contains :", fruitList)
	fmt.Println("Array Length :", len(fruitList))

	var vegList = [3]string {"potato", "beans", "mushroom"}


	fmt.Println("Array contains :", vegList)
	fmt.Println("Array Length :", len(vegList))


}
