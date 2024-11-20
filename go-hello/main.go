package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")
	//info time package
	fmt.Println("The time is", time.Now())
	//info rand package
	fmt.Println("My favorite number is", rand.Intn(10))
	//math package
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
