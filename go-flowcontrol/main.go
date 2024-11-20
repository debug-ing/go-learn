package main

import (
	"fmt"
	"time"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	withStartAndEnd()
	withOutStartAndEnd()
	forWithOutStop()
	IfSample(1)
	IfSample(-1)
	if v := SampleData(); v == 1 {
		fmt.Println("Sample Data is 1")
	} else {
		fmt.Println("Sample Data is not 1")
	}
	fmt.Println(SampleSwitch(1))
	SwitchWithNoCondition()
	SampleDefer()
	//
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
func SampleDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
func SwitchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
func SampleSwitch(x int) string {
	switch x {
	case 1:
		return "One"
	case 2:
		return "Two"
	default:
		return "Unknown"
	}
}
func SampleData() int {
	return 1
}
func IfSample(x int) string {
	if x < 0 {
		return "Negative"
	}
	return "Positive"
}
func withStartAndEnd() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func withOutStartAndEnd() {
	sum := 1
	for sum < 1000 { // this is while in go :)
		sum += sum
	}
	fmt.Println(sum)
}
func forWithOutStop() {
	for { // this is while(true) in go :)
		fmt.Println("Hello")
	}
}
