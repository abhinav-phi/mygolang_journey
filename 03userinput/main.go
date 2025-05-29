package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user login"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of pizza:")

	// comma ok syntax OR error ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks", input)
	fmt.Printf("type of this rating is %T", input)
}
