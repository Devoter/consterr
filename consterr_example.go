package consterr

import "fmt"

func ExampleError() {
	const ErrTest = Error("Test error")

	fmt.Println(ErrTest)

	// Output:
	// Test error
}
