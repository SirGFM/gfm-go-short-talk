package main

import "fmt"

type NonNullError int

func (err NonNullError) Error() string {
	return fmt.Sprintf("Failed with error: %d", err)
}

func main() {
	var err error

	// err is initially nil
	fmt.Printf("err: %+v (err == nil: %+v)\n", err, err == nil)

	var some_err NonNullError

	// now err is storing some_err, so it's not nil, and some_err cannot be nil
	err = some_err
	fmt.Printf("err as some_err: %+v (err == nil: **%+v**)\n", err, err == nil)

	// this works
	fmt.Printf("%+s\n", err.Error())
}
