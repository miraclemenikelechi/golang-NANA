package main

import "fmt"

func main() {
	fmt.Println("### Arrays and Slices in Golang ###")

	var arrayOfNumbers = []int{1, 2, 3, 4, 5}
	fmt.Println(arrayOfNumbers)

	arrayOfStrings := [2]string{"a", "b", "c", "d", "e"}
	fmt.Println(arrayOfStrings)

	// array is fixed ie `arrayOfStrings` is an actual array in golang
	// slices are dynamic ie `arrayOfNumbers` is a slice

	// the code will not execute until `arrayOfStrings` holds the actual amount of data that has been assigned to it. currently it is more than.
}
