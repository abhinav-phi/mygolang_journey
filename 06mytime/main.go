// run exe file ->> .\mytime.exe 

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// "01-02-2006 Monday" this is standard

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.September, 10,23,23,0,0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
