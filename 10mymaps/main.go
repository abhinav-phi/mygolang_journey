package main

import "fmt"

func main() {
	fmt.Println("welcome to maps class")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["CPP"] = "c++"

	fmt.Println(languages)
	fmt.Println("js shorts for:" ,languages["JS"])

	delete(languages, "CPP")
	fmt.Println(languages)


	// loops are interesting in golang

	// for key, value:= range languages{
	// 	fmt.Printf("for key %v, value is %v\n", key, value)
	// }
	
	// the above and lower syntax, both are same,

	for _, value:= range languages{
		fmt.Printf("for key v, value is %v\n", value)
	}


}