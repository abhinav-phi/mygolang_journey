package main

import "fmt"

func main() {
	fmt.Println("welcome to array class")

	var fruitlist [4]string

	fruitlist [0] = "apple"
	fruitlist [1] = "banana"
	fruitlist [2] = "mango"
	fruitlist [3] = "grapes"

	// this can also be the format :
	// var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	// even if the input is apple mango banana, only 3 then also it will return length as 4 as it is defined above

	fmt.Println("fruit list is:", fruitlist)
	fmt.Println("fruit list is:", len(fruitlist))


}