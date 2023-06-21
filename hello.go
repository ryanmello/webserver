package main

import "fmt"

func main() {
	fmt.Println("Welcome to my game")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v", name)
	
	fmt.Println()

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)
	fmt.Printf("You are %v years old", age)
}