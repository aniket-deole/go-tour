package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "WW"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
