package main

import (
	"fmt"
	"sort"
)

func main(){
	//Slices are basically arrays but more powerful 
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	//using memory management - make()
	highScores := make([]int, 4)

	// highScores[0] = 100
	highScores[1] = 200
	highScores[2] = 300
	highScores[3] = 400

	//Error -
	// highScores[4] = 500

	//This is allowed 
	highScores = append(highScores, 999, 888, 777)


	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)


}
