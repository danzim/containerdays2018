package main

import "fmt"

func main() {
	/* for loop infinite
	for true {
		fmt.Printf("This loop will run forever.\n")
	} */

	/* for loop */
	var g int = 15
	var h int

	numbers := [6]int{1, 2, 3, 5}

	for h = 0; h < 10; h++ {
		fmt.Printf("value of h: %d\n", h)
	}
	for h < g {
		h++
		fmt.Printf("value of h: %d\n", h)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}
