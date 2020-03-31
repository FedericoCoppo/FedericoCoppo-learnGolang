package main

import (
	"bufio"
	"fmt"
	"os"
)

// main routine
func main() {
	var name string
	var address string
	var idMap = make(map[string] string)
	reader := bufio.NewReader(os.Stdin)
    var i int = 1
	for {
		fmt.Printf("\nInsert the name %d (to exit print -1) ", i)
		bytes, _, _ := reader.ReadLine()
		name =  string(bytes)
		if name == "-1" {
			break
		}

		fmt.Printf("\nInsert the address %d (to exit print -1) ", i)
		bytes, _, _ = reader.ReadLine()
		address =  string(bytes)

		if address == "-1" {
			break
		}

		// fill the map
		idMap[name] = address
		i++
	}

	fmt.Print("\nMap is:\n")
	fmt.Print(idMap)
	
	// add marshalling
}
