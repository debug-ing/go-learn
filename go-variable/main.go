package main

import (
	"fmt"
)

var c, python, java bool
var i, j int = 1, 2

func main() {
	var i int //0
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	//Short variable declarations
	k := 3
	test, test1, test2 := true, false, "no!"
	fmt.Println(k, test, test1, test2)

	Test()
	//Type inference
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

// about basic type
var (
	ToBe    bool    = false
	Uint    uint64  = 100   //uint uint8 uint16 uint32 uint64
	Int     int64   = -1000 //int int8 int16 int32 int64
	Flout   float64 = 3.14  //float32 float64
	Message string  = "Hello :)"
	//complex128 later add this
)

func Test() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

const Pi = 3.14

func Constants() {
	const World = "Mahdi" //can add type
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true //can add type
	fmt.Println("Go rules?", Truth)
}
