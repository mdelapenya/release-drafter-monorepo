package c_test

import "github.com/mdelapenya/release-drafter-monorepo/modules/c"

func ExampleDoSomething() {
	c.DoSomething("This is an info message")

	// Output:
	// [C] INFO: This is an info message
}
