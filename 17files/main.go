package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	content := "Testing writing to a file"

	file, err := os.Create("./myfile.txt")

	if err != nil{
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil{
		panic(err)
	}

	fmt.Println("Length of the content written is :", length)
	defer file.Close()

	ReadFile("./myfile.txt")

}

func ReadFile(filename string){
	databyte, err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(databyte))

}