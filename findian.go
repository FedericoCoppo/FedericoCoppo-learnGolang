package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var found bool = false;
	fmt.Println("Insert the string")
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	s =  string(bytes)

	if s[0] == 'i' || s[0] == 'I' {
			if s[len(s)-1] == 'n' || s[len(s)-1]  == 'N'{
				for i := 0; i < len(s); i++ {
					if s[i] == 'a' || s[i] == 'A' {
						found = true
						break
					}
				}

			}
		}

		if found {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
}
