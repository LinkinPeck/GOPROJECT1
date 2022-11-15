package main

var y int

func main() {
	cal(sub)
}

func cal(f func(int, int) int) {
	sum := f(50, 10)
	println(sum)
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
