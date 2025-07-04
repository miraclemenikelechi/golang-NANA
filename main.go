package main

import "fmt"

func main() {
	fmt.Println("### Functions in Golang ###")

	greet("world")

}

func greet(name string) {
	fmt.Printf("hello %s", name)
}
