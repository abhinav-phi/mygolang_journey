package main

import "fmt"

func main() {
	fmt.Println("welcome to defer class")

	
	defer fmt.Println("2")
	defer fmt.Println("3")
	
	defer fmt.Println("5")
	fmt.Println("1")
	defer fmt.Println("4")
	mydefer() 

	// 1, 43210, 4,5,3,2
	

	// defer -> (works as last in first out,,,, LIFO) 
	// output ki last line, last curly brakcet
	// ke just pehle print krega vo
}

func mydefer (){
	for i := 0; i<5; i++{
		defer fmt.Print(i)
	}
}
