package main

import "fmt"

func main(){

	var languages = make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")

	fmt.Println(languages)

}
