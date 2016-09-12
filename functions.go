package main

import "fmt"

// Initialize main variables
var var1,var2 string = "Being", "End"
var no bool

// Adds two intergers returns sum
func add(x,y int) int {
	return x+y
}

// Swaps two integers and retuns the two swapped values
func swap(x,y string) (string, string){
	return y,x
}

// Main function where the "magic happens"
func main() {
	var1, var2 := swap(var1, var2)

	fmt.Println(add(10, 20))
	fmt.Println(var1,var2)
	fmt.Println(no)
}