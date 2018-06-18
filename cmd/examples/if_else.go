package main

import "fmt"

func main() {

	var a int = 25

	/* if/else/if else */
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	} else if a == 20 {
		fmt.Printf("a is 20\n")
	} else {
		fmt.Printf("a is more than 20\n")
	}

	fmt.Printf("value of a is : %d\n", a)

}
