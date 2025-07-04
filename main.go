package main

import "fmt"

func main() {
	fmt.Println("### Loops in Golang ###")

	fruits := []string{"apple", "banana", "mango", "orange", "watermelon"}
	fmt.Println("Fruits:", fruits)

	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

}
