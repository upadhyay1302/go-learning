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

	//Deleting a index from a slice
	var courses = []string{"Reactjs", "Javascript", "Swift", "Python", "Ruby"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)


	fmt.Println(courses)


}
