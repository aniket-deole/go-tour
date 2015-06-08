package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//	c <- 3 // buffer overfill.. causes deadlock
	fmt.Println(<-c)
	fmt.Println(<-c)
}
