package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	//sample anonymous function function
	func(x, y int) int {
		return x + y
	}(1, 2)

	adder() //function closures

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func add(x int, y int) int { // can write with func add(x, y int) int
	return x + y
}

func sub(x int, y int) int { // can write with func sub(x, y int) int
	return x - y
}

func swap(x, y string) (string, string) { // Multiple results
	return y, x
}

func split(sum int) (x, y int) { // Named return values
	x = sum * 4 / 9
	y = sum - x
	return
}

// about private and public functions
func PublicFunction() {
	fmt.Println("Public Function")
}
func privateFunction() {
	fmt.Println("private Function")
}
