package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
