package main

import "fmt"

func Swap (sli [] int, index int){
	var temp int= sli[index]
	sli[index] = sli[index + 1]
	sli[index + 1] = temp
}

func BubbleSort(sli [] int) {
	n := len(sli)
	for i, _ := range sli {
		for j := 0; j < n-i-1; j++ {
			if (sli[j] > sli[j+1]) {
				Swap(sli, j)
			}
		}
	}
}

func main(){

	var val int
	// slice creation
	sli:= [] int {}

	for i:= 0;  i < 10; i++ {
		fmt.Printf("Insert value [%d]  ", i)
		fmt.Scanf("%d\n", &val)
		sli = append(sli,val)
	}

	fmt.Printf("\nSlice before ordering:")
	fmt.Print(sli)

	// slice ordering
	BubbleSort(sli)

	fmt.Printf("\nOrdered slice:")
	for i, _ := range sli {
	fmt.Printf("  %d ", sli[i])
	}
}