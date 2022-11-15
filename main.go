package main

var y int

var exname string = "peck"

func main() {
	cal(sub)
}

func cal(f func(int, int) int) {
	sum := f(50, 10)
	println(sum, exname)
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
