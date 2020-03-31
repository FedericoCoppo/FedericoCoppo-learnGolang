package main

import (
	"fmt"
	"runtime"
)

// routine for float truncation
func floatTruncate(x float64) int64 {
	return int64(x)
}

// main routine
func main() {
	var input float64 = 0.0
	var windowsOS bool = false
	var err error

	// check the OS
	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		windowsOS = true
	}

	fmt.Print("Insert the floating point number: ")
	if windowsOS == true {
		_, err = fmt.Scanf("%f\n", &input)       // FOR WINDOWS
	} else {
		_, err = fmt.Scanf("%f", &input) 		 // OK IN UNIX
	}

	if err != nil {
		fmt.Println(err)
	} else
	{
		fmt.Printf("The truncation value of %f is %d \n", input, floatTruncate(input))
		fmt.Print("Insert a second floating point number: ")
		if windowsOS == true {
			_, err = fmt.Scanf("%f\n", &input)       // FOR WINDOWS
		} else {
			_, err = fmt.Scanf("%f", &input) 		 // OK IN UNIX
		}

		if err != nil {
			fmt.Println(err)

		} else
		{
			fmt.Printf("The truncation value of %f is %d \n", input, floatTruncate(input))
		}
	}
}


/*

TEST EXAMPLE

C:\Users\coppo\go\src\Week1Project>go run main.go
Hello from Windows
Insert the floating point number: 3.5
The truncation value of 3.500000 is 3
Insert a second floating point number: 2.9
The truncation value of 2.900000 is 2

 */