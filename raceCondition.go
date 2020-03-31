package main

import (
	"fmt"
	"time"
)

var sharedVariable int = 0

func task1 ()  {

	// critical section
	sharedVariable = - 10

	for i := 0; i < 3000; i++ {
		// do stuff
	}

	fmt.Println("Task1: shared counter value is %d", sharedVariable)
	// end of critical section
}

func task2 ()  {

	// critical section
	sharedVariable = 10

	time.Sleep(1)

	fmt.Println("Task2: shared counter value is %d", sharedVariable)
	// end of critical section
}

// main routine
func main() {

	// loop the two go routine
	for i := 0; i < 5000; i++ {
		go task1()
		go task2()
	}
}


/*
Race Condition: program  output is not deterministic
because the shared variable is not semaphored.

for example if task2 change the variable value and
the Go scheduler give control task1, once the control
is again up to task2, it is printed the wrong value

example:
Task2: shared counter value is %d 10
Task1: shared counter value is %d -10
Task1: shared counter value is %d -10
Task1: shared counter value is %d -10
Task1: shared counter value is %d -10
Task1: shared counter value is %d -10
Task1: shared counter value is %d -10
Task2: shared counter value is %d -10

*/