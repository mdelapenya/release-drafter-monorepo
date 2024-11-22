package a_test

import (
	"github.com/mdelapenya/release-drafter-monorepo/modules/a"
)

func ExampleDoSomething() {
	a.DoSomething("This is an info message")

	// Output:
	// [A] INFO: This is an info message
}
