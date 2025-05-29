package main
/* this will give a error as ":=" is used outside the function

package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}
  
*/

/* how to use cariables outside the function

package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
  
*/
import "fmt"

const A int = 1
const abcd string = "abhinav" //public

func main() {
	fmt.Println("Variables")

	var username string = "abhinav"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallVal1 uint16 = 256
	fmt.Println(smallVal1)
	fmt.Printf("variable is of type : %T \n", smallVal1)

	var smallFloat float32 = 255.6789
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	var smallFloat1 float64 = 255.678967896789
	fmt.Println(smallFloat1)
	fmt.Printf("variable is of type : %T \n", smallFloat1)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable id of type: %T \n", anothervariable)

	// implicit type
	var website = "www.abhinav.com"
	fmt.Println(website)
	fmt.Printf("variable id of type: %T \n", website)

	// no var style
	numberofuser :=300
	fmt.Println(numberofuser)

	// constants
	fmt.Println(abcd)
	fmt.Println(A)
	fmt.Printf("variable id of type: %T \n", abcd)
	fmt.Printf("variable id of type: %T \n", A)


}


