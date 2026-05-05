package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("Response is of type %v\n\n\n", response)

	//it is the callers responsibility to close the response
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil{
		fmt.Println(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
