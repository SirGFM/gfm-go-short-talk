package main

import (
	"fmt"
)

func main() {
	source := make([]byte, 20)
	fmt.Printf("Source initial state: %+v\n", source)

	// Use part of the slice, but add more data at the end
	slice := source[:5]
	for i := 0; i < len(source); i++ {
		slice = append(slice, byte(i))
	}

	// slice writes to source until it's expanded
	fmt.Printf("Source after writing to slice: %+v\n", source)
}
