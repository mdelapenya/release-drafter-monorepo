package b_test

import "github.com/mdelapenya/release-drafter-monorepo/modules/b"

func ExampleDoSomething() {
	b.DoSomething("This is an info message")

	// Output:
	// [B] INFO: This is an info message
}
