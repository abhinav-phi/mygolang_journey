package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome to our pizza app")
	fmt.Println("pls rate our pizza ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thnks for rating ", input)

	// adding 1 to the user rating
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("added 1 to ur rating:", numRating+1)
	}
}
