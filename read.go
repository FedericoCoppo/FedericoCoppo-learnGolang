package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	//"strings"
)


type person struct {
	fname string
	lname  string
}

// main routine
func main() {
	sli := make([] person , 0, 3)
	var err error
	var inputFile string
	var tmpPerson person

	// check the file name
	fmt.Print("Insert the file name:\n")
	if runtime.GOOS == "windows" {
		_, err = fmt.Scanf("%s\n", &inputFile)       // FOR WINDOWS
	} else{
		_, err = fmt.Scanf("%s", &inputFile) 		 // OK IN UNIX
	}

	// Open the file
	readFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	// Create a Scanner for the file
	scanner := bufio.NewScanner(readFile)

	// Read and print each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// parse the line
		_, err = fmt.Sscan(line, &tmpPerson.fname, &tmpPerson.lname);

		// add the struct to the slice
		if tmpPerson.fname != "" && tmpPerson.lname != ""{
			sli = append(sli, tmpPerson)
		}
	}

	fmt.Print("Slice content:\n")
	fmt.Print(sli)

	fmt.Print("\nSlice iteration:\n")
	for i, val := range sli{
		fmt.Printf("first name %d: %s, last name %d: %s \n",i, val.fname, i, val.lname)
	}
}

/*
RUNNING EXAMPLE

Insert the file name:
persone.txt
Slice content:
[{Bob White} {John Red} {Ted Brown} {Mario Rossi}]
Slice iteration:
first name 0: Bob, last name 0: White
first name 1: John, last name 1: Red
first name 2: Ted, last name 2: Brown
first name 3: Mario, last name 3: Rossi

*/