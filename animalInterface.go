package main

import (
	"bufio"
	"fmt"
	"os"
)

// interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// types
type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

// Type who fit interface

// Eat
func (t Cow) Eat()   { fmt.Println("grass") }
func (t Bird) Eat()  { fmt.Println("worms") }
func (t Snake) Eat() { fmt.Println("mice") }

// Move
func (t Cow) Move()   { fmt.Println("walk") }
func (t Bird) Move()  { fmt.Println("fly") }
func (t Snake) Move() { fmt.Println("slither") }

// Speak
func (t Cow) Speak()   { fmt.Println("moo") }
func (t Bird) Speak()  { fmt.Println("peep") }
func (t Snake) Speak() { fmt.Println("hsss") }

// Search slice function
func searchCow(sli []Cow, name string) (int, bool) {
	var found = false
	var i int
	var val Cow

	for i, val = range sli {
		if val.name == name {
			found = true
			break
		}
	}

	return i, found
}

func searchBird(sli []Bird, name string) (int, bool) {
	var found = false
	var i int
	var val Bird

	for i, val = range sli {
		if val.name == name {
			found = true
			break
		}
	}

	return i, found
}

func searchSnake(sli []Snake, name string) (int, bool) {
	var found = false
	var i int
	var val Snake

	for i, val = range sli {
		if val.name == name {
			found = true
			break
		}
	}

	return i, found
}

// main routine
func main() {
	var input, animalName, typeTmp, command string
	var a Animal
	// animals slice
	sliceCow := make([]Cow, 0, 3)
	sliceBird := make([]Bird, 0, 3)
	sliceSnake := make([]Snake, 0, 3)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nInsert 3 string command: newanimal/query name animal_type/animal_info, -1 to exit\n")
		bytes, _, _ := reader.ReadLine()
		input = string(bytes)
		fmt.Sscanf(input, "%s %s %s", &command, &animalName, &typeTmp)

		if input == "-1" {
			break
		}

		if command == "newanimal" {

			// check if already present
			_, foundCow := searchCow(sliceCow, animalName)
			_, foundBird := searchBird(sliceBird, animalName)
			_, foundSnake := searchSnake(sliceSnake, animalName)

			if foundCow || foundBird || foundSnake {
				fmt.Println("Animal already present!")
				continue
			}

			// new animal creation
			if typeTmp == "cow" {
				cow := Cow{animalName}
				sliceCow = append(sliceCow, cow)
				fmt.Println("Created it!")
			} else if typeTmp == "bird" {
				bird := Bird{animalName}
				sliceBird = append(sliceBird, bird)
				fmt.Println("Created it!")
			} else if typeTmp == "snake" {
				snake := Snake{animalName}
				sliceSnake = append(sliceSnake, snake)
				fmt.Println("Created it!")
			} else {
				fmt.Println("Invalid animal type!")
			}
		} else if command == "query" {

			// check if present
			iCow, foundCow := searchCow(sliceCow, animalName)
			iBird, foundBird := searchBird(sliceBird, animalName)
			iSnake, foundSnake := searchSnake(sliceSnake, animalName)

			// assign using interface
			if foundCow {
				a = sliceCow[iCow]
			} else if foundBird {
				a = sliceBird[iBird]
			} else if foundSnake {
				a = sliceSnake[iSnake]
			} else {
				fmt.Println("Animal not found!")
				continue
			}
			switch typeTmp {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Println("Invalid info request!")
			}
		} else {
			fmt.Println("Invalid command!")
		}
	}
}
