package main

import "fmt"

func main() {
	fmt.Println("### Data Types in Golang ###")
	fmt.Println("This is a string data type => ", "Hello World")
	fmt.Println("This is a boolean data type => ", true)
	fmt.Println("This is a float data type => ", 3.14)
	fmt.Println("This is a integer data type => ", 100)
	fmt.Println("This is a complex data type => ", 2i)
	fmt.Println()
	fmt.Println("but as you see, they are all printed in strings. now let's see their types properly.")
	fmt.Println()
	fmt.Println("This is a string data type => ", fmt.Sprintf("%T", "Hello World"))
	fmt.Println("This is a boolean data type => ", fmt.Sprintf("%T", true))
	fmt.Println("This is a float data type => ", fmt.Sprintf("%T", 3.14))
	fmt.Println("This is a integer data type => ", fmt.Sprintf("%T", 100))
	fmt.Println("This is a complex data type => ", fmt.Sprintf("%T", 2i))
}
