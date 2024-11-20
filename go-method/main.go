package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

type IVertex interface {
	Get() Vertex
}

var _ IVertex = &Vertex{}

func (v Vertex) Get() Vertex {
	return v
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Get())
	//
	i := MyInt(100)
	fmt.Println(i.Check())
}

type MyInt int

func (i MyInt) Check() bool {
	if i == 1 {
		return true
	}
	return false
}
