package main

import "fmt"

func main() {
	var name string = "John"
	var score float64 = 5555.5
	fmt.Println("Hello", name)
	var x int = 50

	fmt.Println(x, name, score, &x, &name, &score)
}
