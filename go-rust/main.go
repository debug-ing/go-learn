package main

/*
#cgo LDFLAGS: -L./ -llibcrate
#include <stdint.h>

int32_t add_numbers(int32_t a, int32_t b);
*/
import "C"
import "fmt"

func main() {
	a, b := int32(10), int32(20)
	result := C.add_numbers(C.int32_t(a), C.int32_t(b))
	fmt.Printf("Result from Rust: %d\n", result)
}
