package main

import (
	"fmt"
	"io"
)

type SomeErr int

func (se SomeErr) Error() string {
	return fmt.Sprintf("Failed with err: %d", se)
}

type SomeObject int

func (so SomeObject) Close() error {
	fmt.Printf("Closing...\n")

	// do stuff

	// pretend it failed
	return SomeErr(10)
}

func DoSomething() error {
	so := SomeObject(5)
	defer so.Close()

	// do stuff

	// success(?)
	return nil
}

func main() {
	err := DoSomething()
	fmt.Printf("DoSomething: %+v\n", err)

	fmt.Printf("\n-----------\n\n")

	err = DoSomethingProperly()
	fmt.Printf("DoSomethingProperly: %+v\n", err)
}

func DoSomethingProperly() (err error) {
	so := SomeObject(5)
	defer func(c io.Closer) {
		// Note how here, instead of simply deferring Close(),
		// a wrapper is deferred, which then overwrites the returned error
		// ('err', captured by closure and modified here).

		fmt.Printf("Running wrapped defer...\n")
		_err := c.Close()
		if err == nil {
			// Only overwrites the outer error if the main function succeeded.
			// In a real scenario, if the defer failed,
			// then the error must at least be logged.
			err = _err
		}
	}(so)

	// do stuff

	// success(?)
	return nil
}
