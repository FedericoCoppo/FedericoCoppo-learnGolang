package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sort"
)

// main routine
func main() {
	var windowsOS bool = false
	var err error
	var input string
	var n int
	sli := make([] int , 0, 3)

	// check the OS
	if runtime.GOOS == "windows" {
		windowsOS = true
	}

	for input != "X" {
		fmt.Print("\nInsert the integer value to be added to slice, X to exit ")
		if windowsOS == true {
			_, err = fmt.Scanf("%s\n", &input) // FOR WINDOWS
		} else {
			_, err = fmt.Scanf("%s", &input) // FOR WINDOWSX
		}

		if err != nil {
				fmt.Println(err)
		} else {
			n, err = strconv.Atoi(input)
			if err != nil {
				if input != "X" {
					fmt.Println(err)
				}
			} else {
				// add
				sli = append(sli, n)

				// sort
				sort.Ints(sli)

				// print
				fmt.Print(sli)
			}
		}
	}

	fmt.Print("Exit!")
}
