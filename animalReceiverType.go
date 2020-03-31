package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	food string
	locomotion  string
	noise  string
}

// Init the obj
func (a* Animal) InitMe (food, locomotion, noise string){
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

// Eat
func (a* Animal) Eat() {
	fmt.Println(a.food)
}

// Move
func (a* Animal)  Move() {
	fmt.Println(a.locomotion)
}

// Speak
func (a* Animal)  Speak() {
	fmt.Println(a.noise)
}

// main routine
func main() {
	var input, animalTmp, typeTmp string
	var cow Animal
	var bird Animal
	var snake Animal
	var tmpAnimal * Animal
	var good bool = false

	// init the 3 struct (hardcoded)
	cow.InitMe("grass","walk","moo")
	bird.InitMe("worms","fly","peep")
	snake.InitMe("mice","slither","hsss")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nInsert animal (cow, bird or snake) a space and then info type (eat, move or speak), -1 to exit\n")
		bytes, _, _ := reader.ReadLine()
		input =  string(bytes)
		fmt.Sscanf(input, "%s %s", &animalTmp, &typeTmp);

		if (input == "-1") {
			break
		}

		good = false

		if animalTmp == "cow" || animalTmp == "bird" || animalTmp == "snake" {

			if typeTmp == "eat" || typeTmp == "move" || typeTmp == "speak" {
				good = true
			}
		}

		if (good == false){
			fmt.Println("Invalid input!")
		} else {
				switch animalTmp {
				case "cow":
					tmpAnimal = &cow
				case "bird":
					tmpAnimal = &bird
				case "snake":
					tmpAnimal = &snake
				default:
					tmpAnimal = &cow
				}

				switch typeTmp {
				case "eat":
					tmpAnimal.Eat()
				case "move":
					tmpAnimal.Move()
				case "speak":
					tmpAnimal.Speak()
				default:
					tmpAnimal.Eat()
				}
			}
		}
}
