package main

import "fmt"

// main routine
func main() {
	i := 1;
	defer fmt.Printf("value of i is %d, Bye!\n", i)
	i = 2
	fmt.Println("Hello word!")
}