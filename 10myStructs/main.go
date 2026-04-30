package main

import "fmt"

//Golang version of classes = Structs
//No inheritance, no super or parent

func main(){

	mayank := User{"Mayank", "mayank@go.dev", true, 20}
	fmt.Println(mayank)
	fmt.Printf("Type of struct is %+v\n", mayank)
	fmt.Printf("Name is %v and email is %v", mayank.Name, mayank.Email)

}

type User struct{
	Name string
	Email	string
	Status bool
	Age int
}

