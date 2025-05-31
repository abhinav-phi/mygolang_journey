package main

import "fmt"

func main() {

	fmt.Println("weclcome ot loops class")

	days := []string{"sunday", "mon", "tues","wed","thrus","fri","sat"}

	fmt.Println(days)

	for d:=0; d<len(days); d++{
		fmt.Println(days[d])
	}

	for i:= range days{
		fmt.Println(days[i])
	}

	for index, day :=range days{
		fmt.Printf("index is %v and value is %v\n", index,day)
	}

	roughvalue:=1

	for roughvalue<10{

		if roughvalue ==2{
			goto lco
		}

		if roughvalue ==5{
			roughvalue++
			continue
			// break 
		}

		fmt.Println("valyue is:", roughvalue)
		roughvalue++
	}

	lco:
	fmt.Println("jumpin at abhinav.in")

}