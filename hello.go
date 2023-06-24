package main

import "fmt"

func main() {
	fmt.Println("Welcome to my game")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)
	fmt.Printf("You are %v years old \n", age)

	fmt.Println(age >= 10)

	if age >= 10 {
		fmt.Println("You can play the game")
	} else {
		fmt.Println("You can't play the game")
		return
	}

	fmt.Println("Lets start the game")

	fmt.Println("Today is my birthday")
}