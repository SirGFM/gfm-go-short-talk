package main

import "fmt"

type SomeError int

func (err *SomeError) Error() string {
	return fmt.Sprintf("Failed with error: %d", *err)
}

func main() {
	var err error

	// err is initially nil
	fmt.Printf("err: %+v (err == nil: %+v)\n", err, err == nil)

	var some_err *SomeError

	// now err is storing some_err, so it's not nil... but some_err IS nil
	err = some_err
	fmt.Printf("err as some_err: %+v (err == nil: **%+v**)\n", err, err == nil)

	// this panics
	fmt.Printf("\n------------\n\n")
	fmt.Printf("%+s\n", err.Error())
}
