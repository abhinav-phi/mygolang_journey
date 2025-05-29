package main

import "fmt"

func main() {
	fmt.Println("welcome to class of pointers")

	var ptr *int 
	fmt.Println("value of pointer is ", ptr)

	// for refrencing we use "&"

	mynumber1 := 23

	var ptr1 = &mynumber1 
	fmt.Println("value of my pointer is", ptr1)
	fmt.Println("value of my pointer is", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("new value of ptr1 is", mynumber1)


}