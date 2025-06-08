/*

Quick Update!!
after GO 1.20 we do not need rand.seed() is deprecated we can use directly rand.Intn()

func main() {
	fmt.Println("Switch and case in golang")
	diceNumber := rand.Intn(6) +1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and now you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
                fallthrough

	case 4:
		fmt.Println("You can move 4 spot")
                fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}
}

*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome to switchcase class")
 
	dicenumber := rand.Intn(6) + 1

	fmt.Println("value of dice is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move to 2 spot")
	case 3:
		fmt.Println("you can move to 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move to 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot and roll the dice again")
	default:
		fmt.Println("what was that!")

	}

}
