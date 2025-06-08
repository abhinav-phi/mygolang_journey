package main

import "fmt"

func main() {
	greater()
	g()

	result := adder(3,4)
	fmt.Println("result :", result)

	fmt.Println("welcome to functions class")

	proresult, _ := proadder(2,4,5,6,7,8,9)
	fmt.Println("proresult is :", proresult)

	// "_" means no-care tool.... like here i dont need to declare a string name abcd to print "this is proresult negga"

}

func adder(one int, two int) int{
	return one + two
}

// func proadder(values ...int) int{
func proadder(values ...int) (int, string){
	total := 0

	for _, value := range values{
		total += value
	}
	 return total, "this is proresult negga"
} 

func greater(){
	fmt.Println("hello")
}


func g(){
		fmt.Println("okk")
}

