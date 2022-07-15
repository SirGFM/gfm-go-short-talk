package main

import (
	"fmt"
	"os"
)

func main() {
	f1, err := os.Open("some-file")
	if err != nil {
		fmt.Printf("1. %p: %+v\n", err, err)

		// even though 'err' exists elsewhere, it's shadowed here
		f2, err := os.Open("some-other-files")
		fmt.Printf("2. %p: %+v\n", err, err)

		_ = f2
	}

	// err is back to the original value
	fmt.Printf("3. %p: %+v\n", err, err)

	_ = f1
}
