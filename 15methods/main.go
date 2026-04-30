package main

import "fmt"

//Golang version of classes = Structs
//No inheritance, no super or parent

func main(){

	mayank := User{"Mayank", "mayank@go.dev", true, 20}
	fmt.Println(mayank)
	fmt.Printf("Type of struct is %+v\n", mayank)
	fmt.Printf("Name is %v and email is %v", mayank.Name, mayank.Email)

	mayank.GetStatus()
	mayank.SetNewMail()

	//the actual email is still mayank@go.dev we did not manupilate the pointers
	fmt.Printf("Name is %v and email is %v", mayank.Name, mayank.Email)

}

type User struct{
	Name string
	Email	string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("User status is", u.Status)
}

func (u User) SetNewMail(){
	u.Email = "test@go.dev"
	fmt.Println("User Email is", u.Email)
}