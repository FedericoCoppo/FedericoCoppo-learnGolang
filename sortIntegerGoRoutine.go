package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var sharedVariable int = 0

func sortSli (sli [] int) {
	sort.Ints(sli)
	fmt.Println(sli)
}

func task1 (wg * sync.WaitGroup, sli [] int)  {

	fmt.Println("Task1 series")
	sortSli(sli)
	wg.Done()
}

func task2 (wg * sync.WaitGroup, sli [] int)  {

	fmt.Println("Task2 series")
	sortSli(sli)
	wg.Done()
}

func task3 (wg * sync.WaitGroup, sli [] int)  {

	fmt.Println("Task3 series")
	sortSli(sli)
	wg.Done()
}

func task4 (wg * sync.WaitGroup, sli [] int)  {

	fmt.Println("Task4 series")
	sortSli(sli)
	wg.Done()
}

// main routine
func main() {
	sli := make([] int , 0, 3)

	fmt.Println("Enter a list of integers separated by space")

	in := bufio.NewReader(os.Stdin)
	data, _ := in.ReadString('\n')
	for _, el := range strings.Split(data, " ") {
		i, _ := strconv.Atoi(strings.TrimSpace(el))
		sli = append(sli, i)
	}

	fmt.Println("Original series:")
	fmt.Println(sli)

	lenslice := len(sli)
	chunk := lenslice/4

	fmt.Println("Sub-series:")

	var wg sync.WaitGroup
	wg.Add(4)
	go task1(&wg, sli[0:chunk])
	go task2(&wg, sli[chunk:chunk*2])
	go task3(&wg, sli[chunk*2:chunk*3])
	go task4(&wg, sli[chunk*3:lenslice])

	wg.Wait()

	// main sorting
	fmt.Println("Actual slice after task1,2,3,4 ordering:")
	fmt.Println(sli)
	sort.Ints(sli)
	fmt.Println("Final slice after main() ordering:")
	fmt.Println(sli)
}