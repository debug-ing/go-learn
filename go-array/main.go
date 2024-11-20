package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	//slice
	//between slice and array is that slices are dynamically-sized, whereas arrays are fixed size.
	var s []int = primes[1:4] // 3, 5, 7 this is slice
	fmt.Println(s)

	l := []int{2, 3, 5, 7, 11, 13}

	s = l[1:4] // 3, 5, 7
	fmt.Println(s)

	s = l[:2] // 2, 3
	fmt.Println(s)

	s = l[1:] // 3, 5, 7, 11, 13
	fmt.Println(s)
	//len slice
	fmt.Println(len(s))

	//

	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}

	//heep and stake
	ll := make([]int, 5) //make slice with 5 length and create this
	printSlice("a", ll)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sampleMap() {
	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	//
	var l = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(l)
}

type Vertex struct {
	Lat, Long float64
}
