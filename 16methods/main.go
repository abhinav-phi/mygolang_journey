package main

import "fmt"

func main() {
	fmt.Println("welcome to structures class")
	// no inheritance in golang; no super or parent or child

	abhinav := User{"abhinav", "abhinav@gmail.com", true, 16}
	fmt.Println(abhinav)
	fmt.Printf("abhinav's details are: %+v\n", abhinav)
	fmt.Printf("name is %v and email is %v \n", abhinav.Name, abhinav.Email)
	abhinav.GetStatus()
	abhinav.NewMail()
	fmt.Printf("name is %v and email is %v \n", abhinav.Name, abhinav.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "abhi@gmail.com"
	fmt.Println("email of this user is" ,u.Email)
}
