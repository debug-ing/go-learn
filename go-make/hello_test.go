package main

import "testing"

func TestSayHello(t *testing.T) {
	name := "World"
	expected := "Hello, World!"
	result := SayHello(name)

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
