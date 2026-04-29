package main

import "fmt"


func main(){
	var username string = "Mayank"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.523456
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//Default 

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Variable is of type: %T \n", defaultInt)

	var defaultStr string
	fmt.Println(defaultStr)
	fmt.Printf("Variable is of type: %T \n", defaultStr)

	//Implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numerOfUser := 500
	fmt.Println(numerOfUser)

	
}