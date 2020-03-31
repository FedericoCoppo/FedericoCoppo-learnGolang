package main

import "fmt"

// Gen displace routine
func GenDisplaceFn(acc,v0,s0 float64) func (float64) float64 {

	fn:= func (t float64) float64 {
		return 0.5*acc*t*t +v0*t +s0
	}

	return fn
}

// Function as a return value example
func main(){

	var acc float64 = 0;
	var v0 float64 = 0;
	var s0 float64 = 0;
	var time float64 = 0;

	fmt.Printf("\nInsert acceleration:\n")
	fmt.Scanf("%f\n", &acc)
	fmt.Printf("\nInsert starting speed v0:\n")
	fmt.Scanf("%f\n", &v0)
	fmt.Printf("\nInsert starting spaces0:\n")
	fmt.Scanf("%f\n", &s0)

	fmt.Printf("\nSpecial purpose function linked...\n")
	dist := GenDisplaceFn(acc,v0,s0)

	fmt.Printf("\nInsert  time :\n")
	fmt.Scanf("%f\n", &time)

	// dist(time) allow you to calculate the time

	fmt.Printf("The space calculation (with acc = %f, v0 = %f, s0 = %f, time = %f) is %f", acc, v0, s0, time, dist(time))
}