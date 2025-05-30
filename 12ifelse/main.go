package main

import "fmt"

func main() {
	fmt.Println("welcome to ifelse class")

	logincount := 2
	var result string

	if logincount<10{
		result = "regular user"
	} else if logincount >10{
		result = "watch out"
	} else {
		result = "exactly 10"
	}

	fmt.Println(result)

	if 9%2 ==0{
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	num := 3

	if num<10{
		fmt.Println("<10")
	} else {
		fmt.Println(">10")
	}

	// if err!= nil{

	// }

	
}