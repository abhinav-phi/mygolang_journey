package main

import "fmt"

func main() {
	fmt.Println("welcome to structures class")
	// no inheritance in golang; no super or parent or child

	abhinav := User{"abhinav", "abhinav@gmail.com", true, 16}
	fmt.Println(abhinav)
	fmt.Printf("abhinav's details are: %+v\n", abhinav)
	fmt.Printf("name is %v and email is %v", abhinav.Name, abhinav.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}