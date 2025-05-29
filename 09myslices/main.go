package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices class")

	var fruitlist = []string{"a", "b", "c"}
	fmt.Printf("type of fruitlist is %T\n", fruitlist)
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "d", "e")
	fmt.Println(fruitlist)

	// fruitlist= append(fruitlist [2:])
	// fmt.Println(fruitlist)

	fruitlist = (fruitlist[2:4])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 123
	highscore[2] = 987
	highscore[3] = 789
	fmt.Println(highscore)

	highscore = append(highscore, 456, 444, 555)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	// how to remove a value from slices based on index

	var abc = []string{"a", "b", "c"}
	fmt.Println(abc)

	var index int = 2
	abc = append(abc[:index], abc[index+1:]...)
	fmt.Println(abc)

}
